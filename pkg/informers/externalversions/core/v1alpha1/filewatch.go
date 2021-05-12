/*
Copyright 2020 The Tilt Dev Authors

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

// Code generated by informer-gen. DO NOT EDIT.

package v1alpha1

import (
	"context"
	time "time"

	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
	watch "k8s.io/apimachinery/pkg/watch"
	cache "k8s.io/client-go/tools/cache"

	corev1alpha1 "github.com/tilt-dev/tilt/pkg/apis/core/v1alpha1"
	versioned "github.com/tilt-dev/tilt-pod-shell/pkg/clientset/versioned"
	internalinterfaces "github.com/tilt-dev/tilt-pod-shell/pkg/informers/externalversions/internalinterfaces"
	v1alpha1 "github.com/tilt-dev/tilt-pod-shell/pkg/listers/core/v1alpha1"
)

// FileWatchInformer provides access to a shared informer and lister for
// FileWatches.
type FileWatchInformer interface {
	Informer() cache.SharedIndexInformer
	Lister() v1alpha1.FileWatchLister
}

type fileWatchInformer struct {
	factory          internalinterfaces.SharedInformerFactory
	tweakListOptions internalinterfaces.TweakListOptionsFunc
}

// NewFileWatchInformer constructs a new informer for FileWatch type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewFileWatchInformer(client versioned.Interface, resyncPeriod time.Duration, indexers cache.Indexers) cache.SharedIndexInformer {
	return NewFilteredFileWatchInformer(client, resyncPeriod, indexers, nil)
}

// NewFilteredFileWatchInformer constructs a new informer for FileWatch type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewFilteredFileWatchInformer(client versioned.Interface, resyncPeriod time.Duration, indexers cache.Indexers, tweakListOptions internalinterfaces.TweakListOptionsFunc) cache.SharedIndexInformer {
	return cache.NewSharedIndexInformer(
		&cache.ListWatch{
			ListFunc: func(options v1.ListOptions) (runtime.Object, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.TiltV1alpha1().FileWatches().List(context.TODO(), options)
			},
			WatchFunc: func(options v1.ListOptions) (watch.Interface, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.TiltV1alpha1().FileWatches().Watch(context.TODO(), options)
			},
		},
		&corev1alpha1.FileWatch{},
		resyncPeriod,
		indexers,
	)
}

func (f *fileWatchInformer) defaultInformer(client versioned.Interface, resyncPeriod time.Duration) cache.SharedIndexInformer {
	return NewFilteredFileWatchInformer(client, resyncPeriod, cache.Indexers{cache.NamespaceIndex: cache.MetaNamespaceIndexFunc}, f.tweakListOptions)
}

func (f *fileWatchInformer) Informer() cache.SharedIndexInformer {
	return f.factory.InformerFor(&corev1alpha1.FileWatch{}, f.defaultInformer)
}

func (f *fileWatchInformer) Lister() v1alpha1.FileWatchLister {
	return v1alpha1.NewFileWatchLister(f.Informer().GetIndexer())
}
