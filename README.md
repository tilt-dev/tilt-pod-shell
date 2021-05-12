# tilt-pod-shell

Opens a shell in a pod created by a Tilt resource and switches to new pods for that resource as they come up.

```
$ go run ./cmd/tilt-pod-shell/ vigoda
opening shell in pod matt-vigoda-564ddf9dd9-p7bkz
/go # command terminated with exit code 137
waiting for new pod id
opening shell in pod matt-vigoda-564ddf9dd9-hcc9q
/go #
```
