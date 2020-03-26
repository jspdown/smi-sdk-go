/*
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

	splitv1alpha1 "github.com/servicemeshinterface/smi-sdk-go/pkg/apis/split/v1alpha1"
	versioned "github.com/servicemeshinterface/smi-sdk-go/pkg/gen/client/split/clientset/versioned"
	internalinterfaces "github.com/servicemeshinterface/smi-sdk-go/pkg/gen/client/split/informers/externalversions/internalinterfaces"
	v1alpha1 "github.com/servicemeshinterface/smi-sdk-go/pkg/gen/client/split/listers/split/v1alpha1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
	watch "k8s.io/apimachinery/pkg/watch"
	cache "k8s.io/client-go/tools/cache"
)

// TrafficSplitInformer provides access to a shared informer and lister for
// TrafficSplits.
type TrafficSplitInformer interface {
	Informer() cache.SharedIndexInformer
	Lister() v1alpha1.TrafficSplitLister
}

type trafficSplitInformer struct {
	factory          internalinterfaces.SharedInformerFactory
	tweakListOptions internalinterfaces.TweakListOptionsFunc
	namespace        string
}

// NewTrafficSplitInformer constructs a new informer for TrafficSplit type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewTrafficSplitInformer(client versioned.Interface, namespace string, resyncPeriod time.Duration, indexers cache.Indexers) cache.SharedIndexInformer {
	return NewFilteredTrafficSplitInformer(client, namespace, resyncPeriod, indexers, nil)
}

// NewFilteredTrafficSplitInformer constructs a new informer for TrafficSplit type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewFilteredTrafficSplitInformer(client versioned.Interface, namespace string, resyncPeriod time.Duration, indexers cache.Indexers, tweakListOptions internalinterfaces.TweakListOptionsFunc) cache.SharedIndexInformer {
	return cache.NewSharedIndexInformer(
		&cache.ListWatch{
			ListFunc: func(options v1.ListOptions) (runtime.Object, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.SplitV1alpha1().TrafficSplits(namespace).List(context.TODO(), options)
			},
			WatchFunc: func(options v1.ListOptions) (watch.Interface, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.SplitV1alpha1().TrafficSplits(namespace).Watch(context.TODO(), options)
			},
		},
		&splitv1alpha1.TrafficSplit{},
		resyncPeriod,
		indexers,
	)
}

func (f *trafficSplitInformer) defaultInformer(client versioned.Interface, resyncPeriod time.Duration) cache.SharedIndexInformer {
	return NewFilteredTrafficSplitInformer(client, f.namespace, resyncPeriod, cache.Indexers{cache.NamespaceIndex: cache.MetaNamespaceIndexFunc}, f.tweakListOptions)
}

func (f *trafficSplitInformer) Informer() cache.SharedIndexInformer {
	return f.factory.InformerFor(&splitv1alpha1.TrafficSplit{}, f.defaultInformer)
}

func (f *trafficSplitInformer) Lister() v1alpha1.TrafficSplitLister {
	return v1alpha1.NewTrafficSplitLister(f.Informer().GetIndexer())
}
