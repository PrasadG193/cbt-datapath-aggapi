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

// Code generated by client-gen. DO NOT EDIT.

package fake

import (
	"context"

	v1alpha1 "github.com/PrasadG193/cbt-datapath-aggapi/pkg/apis/cbt/v1alpha1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
)

// FakeVolumeSnapshotDeltaTokens implements VolumeSnapshotDeltaTokenInterface
type FakeVolumeSnapshotDeltaTokens struct {
	Fake *FakeCbtV1alpha1
	ns   string
}

var volumesnapshotdeltatokensResource = schema.GroupVersionResource{Group: "cbt.storage.k8s.io", Version: "v1alpha1", Resource: "volumesnapshotdeltatokens"}

var volumesnapshotdeltatokensKind = schema.GroupVersionKind{Group: "cbt.storage.k8s.io", Version: "v1alpha1", Kind: "VolumeSnapshotDeltaToken"}

// Get takes name of the volumeSnapshotDeltaToken, and returns the corresponding volumeSnapshotDeltaToken object, and an error if there is any.
func (c *FakeVolumeSnapshotDeltaTokens) Get(ctx context.Context, name string, options v1.GetOptions) (result *v1alpha1.VolumeSnapshotDeltaToken, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(volumesnapshotdeltatokensResource, c.ns, name), &v1alpha1.VolumeSnapshotDeltaToken{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.VolumeSnapshotDeltaToken), err
}

// List takes label and field selectors, and returns the list of VolumeSnapshotDeltaTokens that match those selectors.
func (c *FakeVolumeSnapshotDeltaTokens) List(ctx context.Context, opts v1.ListOptions) (result *v1alpha1.VolumeSnapshotDeltaTokenList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(volumesnapshotdeltatokensResource, volumesnapshotdeltatokensKind, c.ns, opts), &v1alpha1.VolumeSnapshotDeltaTokenList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.VolumeSnapshotDeltaTokenList{ListMeta: obj.(*v1alpha1.VolumeSnapshotDeltaTokenList).ListMeta}
	for _, item := range obj.(*v1alpha1.VolumeSnapshotDeltaTokenList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested volumeSnapshotDeltaTokens.
func (c *FakeVolumeSnapshotDeltaTokens) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(volumesnapshotdeltatokensResource, c.ns, opts))

}

// Create takes the representation of a volumeSnapshotDeltaToken and creates it.  Returns the server's representation of the volumeSnapshotDeltaToken, and an error, if there is any.
func (c *FakeVolumeSnapshotDeltaTokens) Create(ctx context.Context, volumeSnapshotDeltaToken *v1alpha1.VolumeSnapshotDeltaToken, opts v1.CreateOptions) (result *v1alpha1.VolumeSnapshotDeltaToken, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(volumesnapshotdeltatokensResource, c.ns, volumeSnapshotDeltaToken), &v1alpha1.VolumeSnapshotDeltaToken{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.VolumeSnapshotDeltaToken), err
}

// Update takes the representation of a volumeSnapshotDeltaToken and updates it. Returns the server's representation of the volumeSnapshotDeltaToken, and an error, if there is any.
func (c *FakeVolumeSnapshotDeltaTokens) Update(ctx context.Context, volumeSnapshotDeltaToken *v1alpha1.VolumeSnapshotDeltaToken, opts v1.UpdateOptions) (result *v1alpha1.VolumeSnapshotDeltaToken, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(volumesnapshotdeltatokensResource, c.ns, volumeSnapshotDeltaToken), &v1alpha1.VolumeSnapshotDeltaToken{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.VolumeSnapshotDeltaToken), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeVolumeSnapshotDeltaTokens) UpdateStatus(ctx context.Context, volumeSnapshotDeltaToken *v1alpha1.VolumeSnapshotDeltaToken, opts v1.UpdateOptions) (*v1alpha1.VolumeSnapshotDeltaToken, error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateSubresourceAction(volumesnapshotdeltatokensResource, "status", c.ns, volumeSnapshotDeltaToken), &v1alpha1.VolumeSnapshotDeltaToken{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.VolumeSnapshotDeltaToken), err
}

// Delete takes name of the volumeSnapshotDeltaToken and deletes it. Returns an error if one occurs.
func (c *FakeVolumeSnapshotDeltaTokens) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteActionWithOptions(volumesnapshotdeltatokensResource, c.ns, name, opts), &v1alpha1.VolumeSnapshotDeltaToken{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeVolumeSnapshotDeltaTokens) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(volumesnapshotdeltatokensResource, c.ns, listOpts)

	_, err := c.Fake.Invokes(action, &v1alpha1.VolumeSnapshotDeltaTokenList{})
	return err
}

// Patch applies the patch and returns the patched volumeSnapshotDeltaToken.
func (c *FakeVolumeSnapshotDeltaTokens) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha1.VolumeSnapshotDeltaToken, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(volumesnapshotdeltatokensResource, c.ns, name, pt, data, subresources...), &v1alpha1.VolumeSnapshotDeltaToken{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.VolumeSnapshotDeltaToken), err
}
