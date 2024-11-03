/*
Copyright The cert-manager Authors.

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

package v1

import (
	v1 "github.com/nholuongut/cert-manager/pkg/apis/certmanager/v1"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/client-go/listers"
	"k8s.io/client-go/tools/cache"
)

// IssuerLister helps list Issuers.
// All objects returned here must be treated as read-only.
type IssuerLister interface {
	// List lists all Issuers in the indexer.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1.Issuer, err error)
	// Issuers returns an object that can list and get Issuers.
	Issuers(namespace string) IssuerNamespaceLister
	IssuerListerExpansion
}

// issuerLister implements the IssuerLister interface.
type issuerLister struct {
	listers.ResourceIndexer[*v1.Issuer]
}

// NewIssuerLister returns a new IssuerLister.
func NewIssuerLister(indexer cache.Indexer) IssuerLister {
	return &issuerLister{listers.New[*v1.Issuer](indexer, v1.Resource("issuer"))}
}

// Issuers returns an object that can list and get Issuers.
func (s *issuerLister) Issuers(namespace string) IssuerNamespaceLister {
	return issuerNamespaceLister{listers.NewNamespaced[*v1.Issuer](s.ResourceIndexer, namespace)}
}

// IssuerNamespaceLister helps list and get Issuers.
// All objects returned here must be treated as read-only.
type IssuerNamespaceLister interface {
	// List lists all Issuers in the indexer for a given namespace.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1.Issuer, err error)
	// Get retrieves the Issuer from the indexer for a given namespace and name.
	// Objects returned here must be treated as read-only.
	Get(name string) (*v1.Issuer, error)
	IssuerNamespaceListerExpansion
}

// issuerNamespaceLister implements the IssuerNamespaceLister
// interface.
type issuerNamespaceLister struct {
	listers.ResourceIndexer[*v1.Issuer]
}
