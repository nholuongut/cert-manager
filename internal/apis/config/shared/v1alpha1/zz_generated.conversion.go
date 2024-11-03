//go:build !ignore_autogenerated
// +build !ignore_autogenerated

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

// Code generated by conversion-gen. DO NOT EDIT.

package v1alpha1

import (
	time "time"
	unsafe "unsafe"

	shared "github.com/nholuongut/cert-manager/internal/apis/config/shared"
	v1alpha1 "github.com/nholuongut/cert-manager/pkg/apis/config/shared/v1alpha1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	conversion "k8s.io/apimachinery/pkg/conversion"
	runtime "k8s.io/apimachinery/pkg/runtime"
)

func init() {
	localSchemeBuilder.Register(RegisterConversions)
}

// RegisterConversions adds conversion functions to the given scheme.
// Public to allow building arbitrary schemes.
func RegisterConversions(s *runtime.Scheme) error {
	if err := s.AddGeneratedConversionFunc((*v1alpha1.DynamicServingConfig)(nil), (*shared.DynamicServingConfig)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1alpha1_DynamicServingConfig_To_shared_DynamicServingConfig(a.(*v1alpha1.DynamicServingConfig), b.(*shared.DynamicServingConfig), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*shared.DynamicServingConfig)(nil), (*v1alpha1.DynamicServingConfig)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_shared_DynamicServingConfig_To_v1alpha1_DynamicServingConfig(a.(*shared.DynamicServingConfig), b.(*v1alpha1.DynamicServingConfig), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*v1alpha1.FilesystemServingConfig)(nil), (*shared.FilesystemServingConfig)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1alpha1_FilesystemServingConfig_To_shared_FilesystemServingConfig(a.(*v1alpha1.FilesystemServingConfig), b.(*shared.FilesystemServingConfig), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*shared.FilesystemServingConfig)(nil), (*v1alpha1.FilesystemServingConfig)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_shared_FilesystemServingConfig_To_v1alpha1_FilesystemServingConfig(a.(*shared.FilesystemServingConfig), b.(*v1alpha1.FilesystemServingConfig), scope)
	}); err != nil {
		return err
	}
	if err := s.AddConversionFunc((**float32)(nil), (*float32)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_Pointer_float32_To_float32(a.(**float32), b.(*float32), scope)
	}); err != nil {
		return err
	}
	if err := s.AddConversionFunc((**int32)(nil), (*int)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_Pointer_int32_To_int(a.(**int32), b.(*int), scope)
	}); err != nil {
		return err
	}
	if err := s.AddConversionFunc((**v1alpha1.Duration)(nil), (*time.Duration)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_Pointer_v1alpha1_Duration_To_time_Duration(a.(**v1alpha1.Duration), b.(*time.Duration), scope)
	}); err != nil {
		return err
	}
	if err := s.AddConversionFunc((*float32)(nil), (**float32)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_float32_To_Pointer_float32(a.(*float32), b.(**float32), scope)
	}); err != nil {
		return err
	}
	if err := s.AddConversionFunc((*int)(nil), (**int32)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_int_To_Pointer_int32(a.(*int), b.(**int32), scope)
	}); err != nil {
		return err
	}
	if err := s.AddConversionFunc((*shared.LeaderElectionConfig)(nil), (*v1alpha1.LeaderElectionConfig)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_shared_LeaderElectionConfig_To_v1alpha1_LeaderElectionConfig(a.(*shared.LeaderElectionConfig), b.(*v1alpha1.LeaderElectionConfig), scope)
	}); err != nil {
		return err
	}
	if err := s.AddConversionFunc((*shared.TLSConfig)(nil), (*v1alpha1.TLSConfig)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_shared_TLSConfig_To_v1alpha1_TLSConfig(a.(*shared.TLSConfig), b.(*v1alpha1.TLSConfig), scope)
	}); err != nil {
		return err
	}
	if err := s.AddConversionFunc((*time.Duration)(nil), (**v1alpha1.Duration)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_time_Duration_To_Pointer_v1alpha1_Duration(a.(*time.Duration), b.(**v1alpha1.Duration), scope)
	}); err != nil {
		return err
	}
	if err := s.AddConversionFunc((*v1alpha1.LeaderElectionConfig)(nil), (*shared.LeaderElectionConfig)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1alpha1_LeaderElectionConfig_To_shared_LeaderElectionConfig(a.(*v1alpha1.LeaderElectionConfig), b.(*shared.LeaderElectionConfig), scope)
	}); err != nil {
		return err
	}
	if err := s.AddConversionFunc((*v1alpha1.TLSConfig)(nil), (*shared.TLSConfig)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1alpha1_TLSConfig_To_shared_TLSConfig(a.(*v1alpha1.TLSConfig), b.(*shared.TLSConfig), scope)
	}); err != nil {
		return err
	}
	return nil
}

func autoConvert_v1alpha1_DynamicServingConfig_To_shared_DynamicServingConfig(in *v1alpha1.DynamicServingConfig, out *shared.DynamicServingConfig, s conversion.Scope) error {
	out.SecretNamespace = in.SecretNamespace
	out.SecretName = in.SecretName
	out.DNSNames = *(*[]string)(unsafe.Pointer(&in.DNSNames))
	if err := Convert_Pointer_v1alpha1_Duration_To_time_Duration(&in.LeafDuration, &out.LeafDuration, s); err != nil {
		return err
	}
	return nil
}

// Convert_v1alpha1_DynamicServingConfig_To_shared_DynamicServingConfig is an autogenerated conversion function.
func Convert_v1alpha1_DynamicServingConfig_To_shared_DynamicServingConfig(in *v1alpha1.DynamicServingConfig, out *shared.DynamicServingConfig, s conversion.Scope) error {
	return autoConvert_v1alpha1_DynamicServingConfig_To_shared_DynamicServingConfig(in, out, s)
}

func autoConvert_shared_DynamicServingConfig_To_v1alpha1_DynamicServingConfig(in *shared.DynamicServingConfig, out *v1alpha1.DynamicServingConfig, s conversion.Scope) error {
	out.SecretNamespace = in.SecretNamespace
	out.SecretName = in.SecretName
	out.DNSNames = *(*[]string)(unsafe.Pointer(&in.DNSNames))
	if err := Convert_time_Duration_To_Pointer_v1alpha1_Duration(&in.LeafDuration, &out.LeafDuration, s); err != nil {
		return err
	}
	return nil
}

// Convert_shared_DynamicServingConfig_To_v1alpha1_DynamicServingConfig is an autogenerated conversion function.
func Convert_shared_DynamicServingConfig_To_v1alpha1_DynamicServingConfig(in *shared.DynamicServingConfig, out *v1alpha1.DynamicServingConfig, s conversion.Scope) error {
	return autoConvert_shared_DynamicServingConfig_To_v1alpha1_DynamicServingConfig(in, out, s)
}

func autoConvert_v1alpha1_FilesystemServingConfig_To_shared_FilesystemServingConfig(in *v1alpha1.FilesystemServingConfig, out *shared.FilesystemServingConfig, s conversion.Scope) error {
	out.CertFile = in.CertFile
	out.KeyFile = in.KeyFile
	return nil
}

// Convert_v1alpha1_FilesystemServingConfig_To_shared_FilesystemServingConfig is an autogenerated conversion function.
func Convert_v1alpha1_FilesystemServingConfig_To_shared_FilesystemServingConfig(in *v1alpha1.FilesystemServingConfig, out *shared.FilesystemServingConfig, s conversion.Scope) error {
	return autoConvert_v1alpha1_FilesystemServingConfig_To_shared_FilesystemServingConfig(in, out, s)
}

func autoConvert_shared_FilesystemServingConfig_To_v1alpha1_FilesystemServingConfig(in *shared.FilesystemServingConfig, out *v1alpha1.FilesystemServingConfig, s conversion.Scope) error {
	out.CertFile = in.CertFile
	out.KeyFile = in.KeyFile
	return nil
}

// Convert_shared_FilesystemServingConfig_To_v1alpha1_FilesystemServingConfig is an autogenerated conversion function.
func Convert_shared_FilesystemServingConfig_To_v1alpha1_FilesystemServingConfig(in *shared.FilesystemServingConfig, out *v1alpha1.FilesystemServingConfig, s conversion.Scope) error {
	return autoConvert_shared_FilesystemServingConfig_To_v1alpha1_FilesystemServingConfig(in, out, s)
}

func autoConvert_v1alpha1_LeaderElectionConfig_To_shared_LeaderElectionConfig(in *v1alpha1.LeaderElectionConfig, out *shared.LeaderElectionConfig, s conversion.Scope) error {
	if err := v1.Convert_Pointer_bool_To_bool(&in.Enabled, &out.Enabled, s); err != nil {
		return err
	}
	out.Namespace = in.Namespace
	if err := Convert_Pointer_v1alpha1_Duration_To_time_Duration(&in.LeaseDuration, &out.LeaseDuration, s); err != nil {
		return err
	}
	if err := Convert_Pointer_v1alpha1_Duration_To_time_Duration(&in.RenewDeadline, &out.RenewDeadline, s); err != nil {
		return err
	}
	if err := Convert_Pointer_v1alpha1_Duration_To_time_Duration(&in.RetryPeriod, &out.RetryPeriod, s); err != nil {
		return err
	}
	return nil
}

func autoConvert_shared_LeaderElectionConfig_To_v1alpha1_LeaderElectionConfig(in *shared.LeaderElectionConfig, out *v1alpha1.LeaderElectionConfig, s conversion.Scope) error {
	if err := v1.Convert_bool_To_Pointer_bool(&in.Enabled, &out.Enabled, s); err != nil {
		return err
	}
	out.Namespace = in.Namespace
	if err := Convert_time_Duration_To_Pointer_v1alpha1_Duration(&in.LeaseDuration, &out.LeaseDuration, s); err != nil {
		return err
	}
	if err := Convert_time_Duration_To_Pointer_v1alpha1_Duration(&in.RenewDeadline, &out.RenewDeadline, s); err != nil {
		return err
	}
	if err := Convert_time_Duration_To_Pointer_v1alpha1_Duration(&in.RetryPeriod, &out.RetryPeriod, s); err != nil {
		return err
	}
	return nil
}

func autoConvert_v1alpha1_TLSConfig_To_shared_TLSConfig(in *v1alpha1.TLSConfig, out *shared.TLSConfig, s conversion.Scope) error {
	out.CipherSuites = *(*[]string)(unsafe.Pointer(&in.CipherSuites))
	out.MinTLSVersion = in.MinTLSVersion
	if err := Convert_v1alpha1_FilesystemServingConfig_To_shared_FilesystemServingConfig(&in.Filesystem, &out.Filesystem, s); err != nil {
		return err
	}
	if err := Convert_v1alpha1_DynamicServingConfig_To_shared_DynamicServingConfig(&in.Dynamic, &out.Dynamic, s); err != nil {
		return err
	}
	return nil
}

func autoConvert_shared_TLSConfig_To_v1alpha1_TLSConfig(in *shared.TLSConfig, out *v1alpha1.TLSConfig, s conversion.Scope) error {
	out.CipherSuites = *(*[]string)(unsafe.Pointer(&in.CipherSuites))
	out.MinTLSVersion = in.MinTLSVersion
	if err := Convert_shared_FilesystemServingConfig_To_v1alpha1_FilesystemServingConfig(&in.Filesystem, &out.Filesystem, s); err != nil {
		return err
	}
	if err := Convert_shared_DynamicServingConfig_To_v1alpha1_DynamicServingConfig(&in.Dynamic, &out.Dynamic, s); err != nil {
		return err
	}
	return nil
}
