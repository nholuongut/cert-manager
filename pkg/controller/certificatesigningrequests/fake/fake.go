/*
Copyright 2021 The cert-manager Authors.

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

package fake

import (
	"context"

	certificatesv1 "k8s.io/api/certificates/v1"

	cmapi "github.com/nholuongut/cert-manager/pkg/apis/certmanager/v1"
)

type Signer struct {
	FakeSign func(context.Context, *certificatesv1.CertificateSigningRequest, cmapi.GenericIssuer) error
}

func (s *Signer) Sign(ctx context.Context, csr *certificatesv1.CertificateSigningRequest, issuerObj cmapi.GenericIssuer) error {
	return s.FakeSign(ctx, csr, issuerObj)
}
