// Code generated by k8s code-generator DO NOT EDIT.

/*
Copyright 2018 Google LLC

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    https://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

// Code generated by informer-gen. DO NOT EDIT.

package v1beta1

import (
	time "time"

	sparkoperatork8siov1beta1 "github.com/kubeflow/arena/pkg/operators/spark-operator/apis/sparkoperator.k8s.io/v1beta1"
	versioned "github.com/kubeflow/arena/pkg/operators/spark-operator/client/clientset/versioned"
	internalinterfaces "github.com/kubeflow/arena/pkg/operators/spark-operator/client/informers/externalversions/internalinterfaces"
	v1beta1 "github.com/kubeflow/arena/pkg/operators/spark-operator/client/listers/sparkoperator.k8s.io/v1beta1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
	watch "k8s.io/apimachinery/pkg/watch"
	cache "k8s.io/client-go/tools/cache"
)

// ScheduledSparkApplicationInformer provides access to a shared informer and lister for
// ScheduledSparkApplications.
type ScheduledSparkApplicationInformer interface {
	Informer() cache.SharedIndexInformer
	Lister() v1beta1.ScheduledSparkApplicationLister
}

type scheduledSparkApplicationInformer struct {
	factory          internalinterfaces.SharedInformerFactory
	tweakListOptions internalinterfaces.TweakListOptionsFunc
	namespace        string
}

// NewScheduledSparkApplicationInformer constructs a new informer for ScheduledSparkApplication type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewScheduledSparkApplicationInformer(client versioned.Interface, namespace string, resyncPeriod time.Duration, indexers cache.Indexers) cache.SharedIndexInformer {
	return NewFilteredScheduledSparkApplicationInformer(client, namespace, resyncPeriod, indexers, nil)
}

// NewFilteredScheduledSparkApplicationInformer constructs a new informer for ScheduledSparkApplication type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewFilteredScheduledSparkApplicationInformer(client versioned.Interface, namespace string, resyncPeriod time.Duration, indexers cache.Indexers, tweakListOptions internalinterfaces.TweakListOptionsFunc) cache.SharedIndexInformer {
	return cache.NewSharedIndexInformer(
		&cache.ListWatch{
			ListFunc: func(options v1.ListOptions) (runtime.Object, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.SparkoperatorV1beta1().ScheduledSparkApplications(namespace).List(options)
			},
			WatchFunc: func(options v1.ListOptions) (watch.Interface, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.SparkoperatorV1beta1().ScheduledSparkApplications(namespace).Watch(options)
			},
		},
		&sparkoperatork8siov1beta1.ScheduledSparkApplication{},
		resyncPeriod,
		indexers,
	)
}

func (f *scheduledSparkApplicationInformer) defaultInformer(client versioned.Interface, resyncPeriod time.Duration) cache.SharedIndexInformer {
	return NewFilteredScheduledSparkApplicationInformer(client, f.namespace, resyncPeriod, cache.Indexers{cache.NamespaceIndex: cache.MetaNamespaceIndexFunc}, f.tweakListOptions)
}

func (f *scheduledSparkApplicationInformer) Informer() cache.SharedIndexInformer {
	return f.factory.InformerFor(&sparkoperatork8siov1beta1.ScheduledSparkApplication{}, f.defaultInformer)
}

func (f *scheduledSparkApplicationInformer) Lister() v1beta1.ScheduledSparkApplicationLister {
	return v1beta1.NewScheduledSparkApplicationLister(f.Informer().GetIndexer())
}