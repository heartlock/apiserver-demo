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

// This file was autogenerated by apiregister-gen. Do not edit it manually!

package v1beta1

import (
	"github.com/heartlock/apiserver-demo/pkg/apis/batch"
	"github.com/kubernetes-incubator/apiserver-builder/pkg/builders"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/runtime/schema"
)

var (
	batchNetworkStorage = builders.NewApiResource( // Resource status endpoint
		batch.InternalNetwork,
		NetworkSchemeFns{},
		func() runtime.Object { return &Network{} },     // Register versioned resource
		func() runtime.Object { return &NetworkList{} }, // Register versioned resource list
		&NetworkStrategy{builders.StorageStrategySingleton},
	)
	ApiVersion = builders.NewApiVersion("batch.heartlock.io", "v1beta1").WithResources(
		batchNetworkStorage,
		builders.NewApiResource( // Resource status endpoint
			batch.InternalNetworkStatus,
			NetworkSchemeFns{},
			func() runtime.Object { return &Network{} },     // Register versioned resource
			func() runtime.Object { return &NetworkList{} }, // Register versioned resource list
			&NetworkStatusStrategy{builders.StatusStorageStrategySingleton},
		))

	// Required by code generated by go2idl
	AddToScheme        = ApiVersion.SchemaBuilder.AddToScheme
	SchemeBuilder      = ApiVersion.SchemaBuilder
	SchemeGroupVersion = ApiVersion.GroupVersion
)

// Required by code generated by go2idl
// Kind takes an unqualified kind and returns a Group qualified GroupKind
func Kind(kind string) schema.GroupKind {
	return SchemeGroupVersion.WithKind(kind).GroupKind()
}

// Required by code generated by go2idl
// Resource takes an unqualified resource and returns a Group qualified GroupResource
func Resource(resource string) schema.GroupResource {
	return SchemeGroupVersion.WithResource(resource).GroupResource()
}

//
// Network Functions and Structs
//
type NetworkSchemeFns struct {
	builders.DefaultSchemeFns
}

type NetworkStrategy struct {
	builders.DefaultStorageStrategy
}

type NetworkStatusStrategy struct {
	builders.DefaultStatusStorageStrategy
}

type NetworkList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Network `json:"items"`
}
