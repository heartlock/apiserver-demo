// +build !ignore_autogenerated

/*
Copyright 2017 The Kubernetes Authors.

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

// This file was autogenerated by conversion-gen. Do not edit it manually!

package v1beta1

import (
	batch "github.com/heartlock/apiserver-demo/pkg/apis/batch"
	conversion "k8s.io/apimachinery/pkg/conversion"
	runtime "k8s.io/apimachinery/pkg/runtime"
	unsafe "unsafe"
)

func init() {
	SchemeBuilder.Register(RegisterConversions)
}

// RegisterConversions adds conversion functions to the given scheme.
// Public to allow building arbitrary schemes.
func RegisterConversions(scheme *runtime.Scheme) error {
	return scheme.AddGeneratedConversionFuncs(
		Convert_v1beta1_Network_To_batch_Network,
		Convert_batch_Network_To_v1beta1_Network,
		Convert_v1beta1_NetworkList_To_batch_NetworkList,
		Convert_batch_NetworkList_To_v1beta1_NetworkList,
		Convert_v1beta1_NetworkSpec_To_batch_NetworkSpec,
		Convert_batch_NetworkSpec_To_v1beta1_NetworkSpec,
		Convert_v1beta1_NetworkStatus_To_batch_NetworkStatus,
		Convert_batch_NetworkStatus_To_v1beta1_NetworkStatus,
		Convert_v1beta1_NetworkStatusStrategy_To_batch_NetworkStatusStrategy,
		Convert_batch_NetworkStatusStrategy_To_v1beta1_NetworkStatusStrategy,
		Convert_v1beta1_NetworkStrategy_To_batch_NetworkStrategy,
		Convert_batch_NetworkStrategy_To_v1beta1_NetworkStrategy,
	)
}

func autoConvert_v1beta1_Network_To_batch_Network(in *Network, out *batch.Network, s conversion.Scope) error {
	out.ObjectMeta = in.ObjectMeta
	if err := Convert_v1beta1_NetworkSpec_To_batch_NetworkSpec(&in.Spec, &out.Spec, s); err != nil {
		return err
	}
	if err := Convert_v1beta1_NetworkStatus_To_batch_NetworkStatus(&in.Status, &out.Status, s); err != nil {
		return err
	}
	return nil
}

// Convert_v1beta1_Network_To_batch_Network is an autogenerated conversion function.
func Convert_v1beta1_Network_To_batch_Network(in *Network, out *batch.Network, s conversion.Scope) error {
	return autoConvert_v1beta1_Network_To_batch_Network(in, out, s)
}

func autoConvert_batch_Network_To_v1beta1_Network(in *batch.Network, out *Network, s conversion.Scope) error {
	out.ObjectMeta = in.ObjectMeta
	if err := Convert_batch_NetworkSpec_To_v1beta1_NetworkSpec(&in.Spec, &out.Spec, s); err != nil {
		return err
	}
	if err := Convert_batch_NetworkStatus_To_v1beta1_NetworkStatus(&in.Status, &out.Status, s); err != nil {
		return err
	}
	return nil
}

// Convert_batch_Network_To_v1beta1_Network is an autogenerated conversion function.
func Convert_batch_Network_To_v1beta1_Network(in *batch.Network, out *Network, s conversion.Scope) error {
	return autoConvert_batch_Network_To_v1beta1_Network(in, out, s)
}

func autoConvert_v1beta1_NetworkList_To_batch_NetworkList(in *NetworkList, out *batch.NetworkList, s conversion.Scope) error {
	out.ListMeta = in.ListMeta
	out.Items = *(*[]batch.Network)(unsafe.Pointer(&in.Items))
	return nil
}

// Convert_v1beta1_NetworkList_To_batch_NetworkList is an autogenerated conversion function.
func Convert_v1beta1_NetworkList_To_batch_NetworkList(in *NetworkList, out *batch.NetworkList, s conversion.Scope) error {
	return autoConvert_v1beta1_NetworkList_To_batch_NetworkList(in, out, s)
}

func autoConvert_batch_NetworkList_To_v1beta1_NetworkList(in *batch.NetworkList, out *NetworkList, s conversion.Scope) error {
	out.ListMeta = in.ListMeta
	if in.Items == nil {
		out.Items = make([]Network, 0)
	} else {
		out.Items = *(*[]Network)(unsafe.Pointer(&in.Items))
	}
	return nil
}

// Convert_batch_NetworkList_To_v1beta1_NetworkList is an autogenerated conversion function.
func Convert_batch_NetworkList_To_v1beta1_NetworkList(in *batch.NetworkList, out *NetworkList, s conversion.Scope) error {
	return autoConvert_batch_NetworkList_To_v1beta1_NetworkList(in, out, s)
}

func autoConvert_v1beta1_NetworkSpec_To_batch_NetworkSpec(in *NetworkSpec, out *batch.NetworkSpec, s conversion.Scope) error {
	return nil
}

// Convert_v1beta1_NetworkSpec_To_batch_NetworkSpec is an autogenerated conversion function.
func Convert_v1beta1_NetworkSpec_To_batch_NetworkSpec(in *NetworkSpec, out *batch.NetworkSpec, s conversion.Scope) error {
	return autoConvert_v1beta1_NetworkSpec_To_batch_NetworkSpec(in, out, s)
}

func autoConvert_batch_NetworkSpec_To_v1beta1_NetworkSpec(in *batch.NetworkSpec, out *NetworkSpec, s conversion.Scope) error {
	return nil
}

// Convert_batch_NetworkSpec_To_v1beta1_NetworkSpec is an autogenerated conversion function.
func Convert_batch_NetworkSpec_To_v1beta1_NetworkSpec(in *batch.NetworkSpec, out *NetworkSpec, s conversion.Scope) error {
	return autoConvert_batch_NetworkSpec_To_v1beta1_NetworkSpec(in, out, s)
}

func autoConvert_v1beta1_NetworkStatus_To_batch_NetworkStatus(in *NetworkStatus, out *batch.NetworkStatus, s conversion.Scope) error {
	return nil
}

// Convert_v1beta1_NetworkStatus_To_batch_NetworkStatus is an autogenerated conversion function.
func Convert_v1beta1_NetworkStatus_To_batch_NetworkStatus(in *NetworkStatus, out *batch.NetworkStatus, s conversion.Scope) error {
	return autoConvert_v1beta1_NetworkStatus_To_batch_NetworkStatus(in, out, s)
}

func autoConvert_batch_NetworkStatus_To_v1beta1_NetworkStatus(in *batch.NetworkStatus, out *NetworkStatus, s conversion.Scope) error {
	return nil
}

// Convert_batch_NetworkStatus_To_v1beta1_NetworkStatus is an autogenerated conversion function.
func Convert_batch_NetworkStatus_To_v1beta1_NetworkStatus(in *batch.NetworkStatus, out *NetworkStatus, s conversion.Scope) error {
	return autoConvert_batch_NetworkStatus_To_v1beta1_NetworkStatus(in, out, s)
}

func autoConvert_v1beta1_NetworkStatusStrategy_To_batch_NetworkStatusStrategy(in *NetworkStatusStrategy, out *batch.NetworkStatusStrategy, s conversion.Scope) error {
	out.DefaultStatusStorageStrategy = in.DefaultStatusStorageStrategy
	return nil
}

// Convert_v1beta1_NetworkStatusStrategy_To_batch_NetworkStatusStrategy is an autogenerated conversion function.
func Convert_v1beta1_NetworkStatusStrategy_To_batch_NetworkStatusStrategy(in *NetworkStatusStrategy, out *batch.NetworkStatusStrategy, s conversion.Scope) error {
	return autoConvert_v1beta1_NetworkStatusStrategy_To_batch_NetworkStatusStrategy(in, out, s)
}

func autoConvert_batch_NetworkStatusStrategy_To_v1beta1_NetworkStatusStrategy(in *batch.NetworkStatusStrategy, out *NetworkStatusStrategy, s conversion.Scope) error {
	out.DefaultStatusStorageStrategy = in.DefaultStatusStorageStrategy
	return nil
}

// Convert_batch_NetworkStatusStrategy_To_v1beta1_NetworkStatusStrategy is an autogenerated conversion function.
func Convert_batch_NetworkStatusStrategy_To_v1beta1_NetworkStatusStrategy(in *batch.NetworkStatusStrategy, out *NetworkStatusStrategy, s conversion.Scope) error {
	return autoConvert_batch_NetworkStatusStrategy_To_v1beta1_NetworkStatusStrategy(in, out, s)
}

func autoConvert_v1beta1_NetworkStrategy_To_batch_NetworkStrategy(in *NetworkStrategy, out *batch.NetworkStrategy, s conversion.Scope) error {
	out.DefaultStorageStrategy = in.DefaultStorageStrategy
	return nil
}

// Convert_v1beta1_NetworkStrategy_To_batch_NetworkStrategy is an autogenerated conversion function.
func Convert_v1beta1_NetworkStrategy_To_batch_NetworkStrategy(in *NetworkStrategy, out *batch.NetworkStrategy, s conversion.Scope) error {
	return autoConvert_v1beta1_NetworkStrategy_To_batch_NetworkStrategy(in, out, s)
}

func autoConvert_batch_NetworkStrategy_To_v1beta1_NetworkStrategy(in *batch.NetworkStrategy, out *NetworkStrategy, s conversion.Scope) error {
	out.DefaultStorageStrategy = in.DefaultStorageStrategy
	return nil
}

// Convert_batch_NetworkStrategy_To_v1beta1_NetworkStrategy is an autogenerated conversion function.
func Convert_batch_NetworkStrategy_To_v1beta1_NetworkStrategy(in *batch.NetworkStrategy, out *NetworkStrategy, s conversion.Scope) error {
	return autoConvert_batch_NetworkStrategy_To_v1beta1_NetworkStrategy(in, out, s)
}