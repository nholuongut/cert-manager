/*
Copyright 2020 The cert-manager Authors.

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

package controller

import (
	"reflect"
	"strings"
	"time"

	"github.com/go-logr/logr"
	corev1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/api/errors"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime/schema"
	"k8s.io/apimachinery/pkg/types"
	"k8s.io/apimachinery/pkg/util/runtime"
	"k8s.io/client-go/tools/cache"
	"k8s.io/client-go/util/workqueue"

	logf "github.com/nholuongut/cert-manager/pkg/logs"
)

// DefaultItemBasedRateLimiter returns a new rate limiter with base delay of 5
// seconds, max delay of 5 minutes.
func DefaultItemBasedRateLimiter() workqueue.TypedRateLimiter[types.NamespacedName] {
	return workqueue.NewTypedItemExponentialFailureRateLimiter[types.NamespacedName](time.Second*5, time.Minute*5)
}

// DefaultCertificateRateLimiter returns a new rate limiter with base delay of 1
// seconds, max delay of 30 seconds.
func DefaultCertificateRateLimiter() workqueue.TypedRateLimiter[types.NamespacedName] {
	return workqueue.NewTypedItemExponentialFailureRateLimiter[types.NamespacedName](time.Second*1, time.Second*30)
}

// DefaultCertificateRateLimiter returns a new rate limiter with base delay of 5
// seconds, max delay of 30 minutes.
func DefaultACMERateLimiter() workqueue.TypedRateLimiter[types.NamespacedName] {
	return workqueue.NewTypedItemExponentialFailureRateLimiter[types.NamespacedName](time.Second*5, time.Minute*30)
}

// HandleOwnedResourceNamespacedFunc returns a function that accepts a
// Kubernetes object and adds its owner references to the workqueue.
// https://kubernetes.io/docs/concepts/workloads/controllers/garbage-collection/#owners-and-dependents
func HandleOwnedResourceNamespacedFunc[T metav1.Object](
	log logr.Logger,
	queue workqueue.TypedRateLimitingInterface[types.NamespacedName],
	ownerGVK schema.GroupVersionKind,
	get func(namespace, name string) (T, error),
) func(obj interface{}) {
	return func(obj interface{}) {
		log := log.WithName("handleOwnedResource")

		metaobj, ok := obj.(metav1.Object)
		if !ok {
			log.Error(nil, "item passed to handleOwnedResource does not implement metav1.Object")
			return
		}
		log = logf.WithResource(log, metaobj)

		ownerRefs := metaobj.GetOwnerReferences()
		for _, ref := range ownerRefs {
			log := log.WithValues(
				logf.RelatedResourceNamespaceKey, metaobj.GetNamespace(),
				logf.RelatedResourceNameKey, ref.Name,
				logf.RelatedResourceKindKey, ref.Kind,
			)

			// Parse the Group out of the OwnerReference to compare it to what was parsed out of the requested OwnerType
			refGV, err := schema.ParseGroupVersion(ref.APIVersion)
			if err != nil {
				log.Error(err, "could not parse OwnerReference GroupVersion")
				continue
			}

			if refGV.Group == ownerGVK.Group && ref.Kind == ownerGVK.Kind {
				obj, err := get(metaobj.GetNamespace(), ref.Name)
				// This function is always called with a getter
				// that gets from informers cache. Because this
				// is also called on cache sync it may be that
				// the owner is not yet in the cache.
				if err != nil && errors.IsNotFound(err) {
					log.Info("owning resource not found in cache")
					continue
				}
				if err != nil {
					log.Error(err, "error getting referenced owning resource from cache")
					continue
				}
				queue.Add(types.NamespacedName{
					Name:      obj.GetName(),
					Namespace: obj.GetNamespace(),
				})
			}
		}
	}
}

// QueuingEventHandler is an implementation of cache.ResourceEventHandler that
// simply queues objects that are added/updated/deleted.
type QueuingEventHandler struct {
	Queue workqueue.TypedRateLimitingInterface[types.NamespacedName]
}

// Enqueue adds a key for an object to the workqueue.
func (q *QueuingEventHandler) Enqueue(obj interface{}) {
	objectName, err := cache.DeletionHandlingObjectToName(obj)
	if err != nil {
		runtime.HandleError(err)
		return
	}
	q.Queue.Add(types.NamespacedName{
		Name:      objectName.Name,
		Namespace: objectName.Namespace,
	})
}

// OnAdd adds a newly created object to the workqueue.
func (q *QueuingEventHandler) OnAdd(obj interface{}, isInInitialList bool) {
	q.Enqueue(obj)
}

// OnUpdate adds an updated object to the workqueue.
func (q *QueuingEventHandler) OnUpdate(oldObj, newObj interface{}) {
	if reflect.DeepEqual(oldObj, newObj) {
		return
	}
	q.Enqueue(newObj)
}

// OnDelete adds a deleted object to the workqueue for processing.
func (q *QueuingEventHandler) OnDelete(obj interface{}) {
	tombstone, ok := obj.(cache.DeletedFinalStateUnknown)
	if ok {
		obj = tombstone.Obj
	}
	q.Enqueue(obj)
}

// BlockingEventHandler is an implementation of cache.ResourceEventHandler that
// simply synchronously calls it's WorkFunc upon calls to OnAdd, OnUpdate or
// OnDelete.
type BlockingEventHandler struct {
	WorkFunc func(obj interface{})
}

// Enqueue synchronously adds a key for an object to the workqueue.
func (b *BlockingEventHandler) Enqueue(obj interface{}) {
	b.WorkFunc(obj)
}

// OnAdd synchronously adds a newly created object to the workqueue.
func (b *BlockingEventHandler) OnAdd(obj interface{}, isInInitialList bool) {
	b.WorkFunc(obj)
}

// OnUpdate synchronously adds an updated object to the workqueue.
func (b *BlockingEventHandler) OnUpdate(oldObj, newObj interface{}) {
	if reflect.DeepEqual(oldObj, newObj) {
		return
	}
	b.WorkFunc(newObj)
}

// OnDelete synchronously adds a deleted object to the workqueue.
func (b *BlockingEventHandler) OnDelete(obj interface{}) {
	tombstone, ok := obj.(cache.DeletedFinalStateUnknown)
	if ok {
		obj = tombstone.Obj
	}
	b.WorkFunc(obj)
}

// BuildAnnotationsToCopy takes a map of annotations and a list of prefix
// filters and builds a filtered map of annotations. It is used to filter
// annotations to be copied from Certificate to CertificateRequest and from
// CertificateSigningRequest to Order.
func BuildAnnotationsToCopy(allAnnotations map[string]string, prefixes []string) map[string]string {
	filteredAnnotations := make(map[string]string)
	includeAll := false
	for _, v := range prefixes {
		if v == "*" {
			includeAll = true
		}
	}
	for _, annotation := range prefixes {
		prefix := strings.TrimPrefix(annotation, "-")
		for k, v := range allAnnotations {
			if strings.HasPrefix(annotation, "-") {
				if strings.HasPrefix(k, prefix) {
					// If this is an annotation to not be copied.
					delete(filteredAnnotations, k)
				}
			} else if includeAll || strings.HasPrefix(k, annotation) {
				// If this is an annotation to be copied or if 'all' should be copied.
				filteredAnnotations[k] = v
			}
		}
	}
	return filteredAnnotations
}

func ToSecret(obj interface{}) (*corev1.Secret, bool) {
	secret, ok := obj.(*corev1.Secret)
	if !ok {
		meta, ok := obj.(*metav1.PartialObjectMetadata)
		if !ok {
			// TODO: I wasn't able to get GVK from PartialMetadata,
			// however perhaps this should be possible and then we
			// could verify that this really is a Secret. At the
			// moment this is okay as there is no path how any
			// reconcile loop would receive PartialObjectMetadata
			// for any other type.
			return nil, false
		}
		secret = &corev1.Secret{}
		secret.SetName(meta.Name)
		secret.SetNamespace(meta.Namespace)
	}
	return secret, true
}
