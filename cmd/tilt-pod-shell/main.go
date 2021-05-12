package main

import (
	"context"
	"fmt"
	clientset "github.com/tilt-dev/tilt-pod-shell/pkg/clientset/versioned"
	"github.com/tilt-dev/tilt-pod-shell/pkg/config"
	"github.com/tilt-dev/tilt/pkg/apis/core/v1alpha1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"os"
	"os/exec"
	"strings"
)

func watchDiscoveries() (chan *v1alpha1.KubernetesDiscovery, context.CancelFunc, error) {
	tiltAPIConfig, err := config.NewConfig()
	if err != nil {
		return nil, nil, err
	}

	cli := clientset.NewForConfigOrDie(tiltAPIConfig)

	ctx, cancel := context.WithCancel(context.Background())

	w, err := cli.TiltV1alpha1().KubernetesDiscoveries().Watch(ctx, v1.ListOptions{})
	if err != nil {
		return nil, cancel, err
	}

	ch := make(chan *v1alpha1.KubernetesDiscovery)
	go func() {
		for {
			select {
			case <-ctx.Done():
				w.Stop()
			case e, ok := <-w.ResultChan():
				if !ok {
					close(ch)
					return
				}
				ch <- e.Object.(*v1alpha1.KubernetesDiscovery)
			}
		}
	}()

	return ch, cancel, nil
}

func podIDsForResource(resource string, kds chan *v1alpha1.KubernetesDiscovery) chan string {
	ret := make(chan string)
	go func() {
		for kd := range kds {
			if kd.Annotations["tilt.dev/resource"] != resource {
				continue
			}
			var newestPod v1alpha1.Pod
			// if there are multiple pods, get the newest
			for _, pod := range kd.Status.Pods {
				if pod.Status == "Running" && pod.CreatedAt.After(newestPod.CreatedAt.Time) {
					newestPod = pod
				}
			}
			if newestPod.Name != "" {
				ret <- newestPod.Name
			}
		}
		close(ret)
	}()
	return ret
}

func parseArgs() string {
	if len(os.Args) != 2 {
		_, _ = fmt.Fprintln(os.Stderr, "Usage: tilt-pod-shell <resource name>")
		os.Exit(1)
	}
	return os.Args[1]
}

func run() error {
	kds, cancel, err := watchDiscoveries()
	if err != nil {
		return err
	}
	defer cancel()

	resourceName := parseArgs()
	var curCmd *exec.Cmd = nil
	curPod := ""
	for podID := range podIDsForResource(resourceName, kds) {
		if podID == curPod {
			continue
		}
		curPod = podID
		fmt.Printf("opening shell in pod %s\n", podID)
		if curCmd != nil {
			err := curCmd.Process.Kill()
			if err != nil && !strings.Contains(err.Error(), "process already finished") {
				fmt.Printf("error killing old ssh: %v\n", err)
			}
		}
		curCmd = exec.Command("kubectl", "exec", "-it", podID, "--", "/bin/sh")
		curCmd.Stdout = os.Stdout
		curCmd.Stderr = os.Stderr
		curCmd.Stdin = os.Stdin
		err := curCmd.Start()
		if err != nil {
			fmt.Printf("error launching a shell in the pod: %v\n", err)
		}
		go func() {
			err := curCmd.Wait()
			if err == nil {
				// user exited ssh - let's take that to mean they're done done
				cancel()
				return
			} else if strings.Contains(err.Error(), "exit status 137") {
				fmt.Printf("waiting for new pod id\n")
			} else {
				fmt.Printf("kubectl exited with error: %v\n", err)
			}
		}()
	}
	fmt.Println()
	return nil
}

func main() {
	err := run()
	if err != nil {
		_, _ = fmt.Fprintf(os.Stderr, "%v\n", err)
		os.Exit(1)
	}
}
