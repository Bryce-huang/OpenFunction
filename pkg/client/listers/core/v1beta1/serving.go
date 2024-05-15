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
// Code generated by lister-gen. DO NOT EDIT.

package v1beta1

import (
	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/client-go/tools/cache"

	v1beta1 "github.com/openfunction/apis/core/v1beta1"
)

// ServingLister helps list Servings.
// All objects returned here must be treated as read-only.
type ServingLister interface {
	// List lists all Servings in the indexer.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1beta1.Serving, err error)
	// Servings returns an object that can list and get Servings.
	Servings(namespace string) ServingNamespaceLister
	ServingListerExpansion
}

// servingLister implements the ServingLister interface.
type servingLister struct {
	indexer cache.Indexer
}

// NewServingLister returns a new ServingLister.
func NewServingLister(indexer cache.Indexer) ServingLister {
	return &servingLister{indexer: indexer}
}

// List lists all Servings in the indexer.
func (s *servingLister) List(selector labels.Selector) (ret []*v1beta1.Serving, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1beta1.Serving))
	})
	return ret, err
}

// Servings returns an object that can list and get Servings.
func (s *servingLister) Servings(namespace string) ServingNamespaceLister {
	return servingNamespaceLister{indexer: s.indexer, namespace: namespace}
}

// ServingNamespaceLister helps list and get Servings.
// All objects returned here must be treated as read-only.
type ServingNamespaceLister interface {
	// List lists all Servings in the indexer for a given namespace.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1beta1.Serving, err error)
	// Get retrieves the Serving from the indexer for a given namespace and name.
	// Objects returned here must be treated as read-only.
	Get(name string) (*v1beta1.Serving, error)
	ServingNamespaceListerExpansion
}

// servingNamespaceLister implements the ServingNamespaceLister
// interface.
type servingNamespaceLister struct {
	indexer   cache.Indexer
	namespace string
}

// List lists all Servings in the indexer for a given namespace.
func (s servingNamespaceLister) List(selector labels.Selector) (ret []*v1beta1.Serving, err error) {
	err = cache.ListAllByNamespace(s.indexer, s.namespace, selector, func(m interface{}) {
		ret = append(ret, m.(*v1beta1.Serving))
	})
	return ret, err
}

// Get retrieves the Serving from the indexer for a given namespace and name.
func (s servingNamespaceLister) Get(name string) (*v1beta1.Serving, error) {
	obj, exists, err := s.indexer.GetByKey(s.namespace + "/" + name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1beta1.Resource("serving"), name)
	}
	return obj.(*v1beta1.Serving), nil
}
