/*
Copyright 2022 The OpenFunction Authors.

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

	eventsv1alpha1 "github.com/openfunction/apis/events/v1alpha1"
	versioned "github.com/openfunction/pkg/client/clientset/versioned"
	internalinterfaces "github.com/openfunction/pkg/client/informers/externalversions/internalinterfaces"
	v1alpha1 "github.com/openfunction/pkg/client/listers/events/v1alpha1"
)

// ClusterEventBusInformer provides access to a shared informer and lister for
// ClusterEventBuses.
type ClusterEventBusInformer interface {
	Informer() cache.SharedIndexInformer
	Lister() v1alpha1.ClusterEventBusLister
}

type clusterEventBusInformer struct {
	factory          internalinterfaces.SharedInformerFactory
	tweakListOptions internalinterfaces.TweakListOptionsFunc
	namespace        string
}

// NewClusterEventBusInformer constructs a new informer for ClusterEventBus type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewClusterEventBusInformer(client versioned.Interface, namespace string, resyncPeriod time.Duration, indexers cache.Indexers) cache.SharedIndexInformer {
	return NewFilteredClusterEventBusInformer(client, namespace, resyncPeriod, indexers, nil)
}

// NewFilteredClusterEventBusInformer constructs a new informer for ClusterEventBus type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewFilteredClusterEventBusInformer(client versioned.Interface, namespace string, resyncPeriod time.Duration, indexers cache.Indexers, tweakListOptions internalinterfaces.TweakListOptionsFunc) cache.SharedIndexInformer {
	return cache.NewSharedIndexInformer(
		&cache.ListWatch{
			ListFunc: func(options v1.ListOptions) (runtime.Object, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.EventsV1alpha1().ClusterEventBuses(namespace).List(context.TODO(), options)
			},
			WatchFunc: func(options v1.ListOptions) (watch.Interface, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.EventsV1alpha1().ClusterEventBuses(namespace).Watch(context.TODO(), options)
			},
		},
		&eventsv1alpha1.ClusterEventBus{},
		resyncPeriod,
		indexers,
	)
}

func (f *clusterEventBusInformer) defaultInformer(client versioned.Interface, resyncPeriod time.Duration) cache.SharedIndexInformer {
	return NewFilteredClusterEventBusInformer(client, f.namespace, resyncPeriod, cache.Indexers{cache.NamespaceIndex: cache.MetaNamespaceIndexFunc}, f.tweakListOptions)
}

func (f *clusterEventBusInformer) Informer() cache.SharedIndexInformer {
	return f.factory.InformerFor(&eventsv1alpha1.ClusterEventBus{}, f.defaultInformer)
}

func (f *clusterEventBusInformer) Lister() v1alpha1.ClusterEventBusLister {
	return v1alpha1.NewClusterEventBusLister(f.Informer().GetIndexer())
}
