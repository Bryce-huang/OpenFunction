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
	internalinterfaces "github.com/openfunction/pkg/client/informers/externalversions/internalinterfaces"
)

// Interface provides access to all the informers in this group version.
type Interface interface {
	// ClusterEventBuses returns a ClusterEventBusInformer.
	ClusterEventBuses() ClusterEventBusInformer
	// EventBuses returns a EventBusInformer.
	EventBuses() EventBusInformer
	// EventSources returns a EventSourceInformer.
	EventSources() EventSourceInformer
	// Triggers returns a TriggerInformer.
	Triggers() TriggerInformer
}

type version struct {
	factory          internalinterfaces.SharedInformerFactory
	namespace        string
	tweakListOptions internalinterfaces.TweakListOptionsFunc
}

// New returns a new Interface.
func New(f internalinterfaces.SharedInformerFactory, namespace string, tweakListOptions internalinterfaces.TweakListOptionsFunc) Interface {
	return &version{factory: f, namespace: namespace, tweakListOptions: tweakListOptions}
}

// ClusterEventBuses returns a ClusterEventBusInformer.
func (v *version) ClusterEventBuses() ClusterEventBusInformer {
	return &clusterEventBusInformer{factory: v.factory, namespace: v.namespace, tweakListOptions: v.tweakListOptions}
}

// EventBuses returns a EventBusInformer.
func (v *version) EventBuses() EventBusInformer {
	return &eventBusInformer{factory: v.factory, namespace: v.namespace, tweakListOptions: v.tweakListOptions}
}

// EventSources returns a EventSourceInformer.
func (v *version) EventSources() EventSourceInformer {
	return &eventSourceInformer{factory: v.factory, namespace: v.namespace, tweakListOptions: v.tweakListOptions}
}

// Triggers returns a TriggerInformer.
func (v *version) Triggers() TriggerInformer {
	return &triggerInformer{factory: v.factory, namespace: v.namespace, tweakListOptions: v.tweakListOptions}
}
