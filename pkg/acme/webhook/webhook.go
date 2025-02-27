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

// Package webhook provides a library that can be used to build external ACME
// solver webhooks.
package webhook

import (
	restclient "k8s.io/client-go/rest"

	whapi "github.com/nholuongut/cert-manager/pkg/acme/webhook/apis/acme/v1alpha1"
)

// Solver has the functionality to solve ACME challenges. This interface is
// implemented internally by RFC2136 DNS provider and by external webhook solver
// implementations see https://github.com/cert-manager/webhook-example
type Solver interface {
	// Name is the name of this ACME solver as part of the API group.
	// This must match what you configure in the ACME Issuer's DNS01 config.
	Name() string

	// Present should 'present' the ACME challenge solving parameters as
	// defined in the given challenge resource.
	// TODO: add notes about duplicate records with DNS01
	Present(ch *whapi.ChallengeRequest) error

	// CleanUp should remove any presented challenge records for the given
	// challenge resource
	// TODO: add notes about duplicate records with DNS01
	CleanUp(ch *whapi.ChallengeRequest) error

	// Initialize is called as a post-start hook when the apiserver starts.
	// https://github.com/kubernetes/apiserver/blob/release-1.26/pkg/server/hooks.go#L32-L42
	Initialize(kubeClientConfig *restclient.Config, stopCh <-chan struct{}) error
}
