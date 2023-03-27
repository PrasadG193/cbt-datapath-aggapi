/*
Copyright 2023.

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

package v1alpha1

import (
	"context"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/runtime/schema"
	"k8s.io/apimachinery/pkg/util/validation/field"
	"sigs.k8s.io/apiserver-runtime/pkg/builder/resource"
	"sigs.k8s.io/apiserver-runtime/pkg/builder/resource/resourcestrategy"
)

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// VolumeSnapshotDeltaToken
// +k8s:openapi-gen=true
type VolumeSnapshotDeltaToken struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   VolumeSnapshotDeltaTokenSpec   `json:"spec,omitempty"`
	Status VolumeSnapshotDeltaTokenStatus `json:"status,omitempty"`
}

// VolumeSnapshotDeltaTokenList
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
type VolumeSnapshotDeltaTokenList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`

	Items []VolumeSnapshotDeltaToken `json:"items"`
}

// VolumeSnapshotDeltaTokenSpec defines the desired state of VolumeSnapshotDeltaToken
type VolumeSnapshotDeltaTokenSpec struct {
	// The name of the base CSI volume snapshot to use for comparison.
	// If not specified, return all changed blocks.
	// +optional
	BaseVolumeSnapshotName string `json:"baseVolumeSnapshotName,omitempty"`

	// The name of the target CSI volume snapshot to use for comparison.
	// Required.
	TargetVolumeSnapshotName string `json:"targetVolumeSnapshotName"`

	// Defines the type of volume. Default to "block".
	// Required.
	Mode string `json:"mode,omitempty"`
}

var _ resource.Object = &VolumeSnapshotDeltaToken{}
var _ resourcestrategy.Validater = &VolumeSnapshotDeltaToken{}

func (in *VolumeSnapshotDeltaToken) GetObjectMeta() *metav1.ObjectMeta {
	return &in.ObjectMeta
}

func (in *VolumeSnapshotDeltaToken) NamespaceScoped() bool {
	return true
}

func (in *VolumeSnapshotDeltaToken) New() runtime.Object {
	return &VolumeSnapshotDeltaToken{}
}

func (in *VolumeSnapshotDeltaToken) NewList() runtime.Object {
	return &VolumeSnapshotDeltaTokenList{}
}

func (in *VolumeSnapshotDeltaToken) GetGroupVersionResource() schema.GroupVersionResource {
	return schema.GroupVersionResource{
		Group:    "cbt.storage.k8s.io",
		Version:  "v1alpha1",
		Resource: "volumesnapshottokens",
	}
}

func (in *VolumeSnapshotDeltaToken) IsStorageVersion() bool {
	return true
}

func (in *VolumeSnapshotDeltaToken) Validate(ctx context.Context) field.ErrorList {
	// TODO(user): Modify it, adding your API validation here.
	return nil
}

var _ resource.ObjectList = &VolumeSnapshotDeltaTokenList{}

func (in *VolumeSnapshotDeltaTokenList) GetListMeta() *metav1.ListMeta {
	return &in.ListMeta
}

// VolumeSnapshotDeltaTokenStatus defines the observed state of VolumeSnapshotDeltaToken
type VolumeSnapshotDeltaTokenStatus struct {
	// Captures any error encountered.
	Error string `json:"error,omitempty"`

	// CABundle client side CA used for server validation
	CABundle []byte `json:"cabundle,omitempty"`

	// Token cbt server token for validation
	Token []byte `json:"token,omitempty"`

	// URL to get CBT metadata from
	URL string `json:"url,omitempty"`
}
