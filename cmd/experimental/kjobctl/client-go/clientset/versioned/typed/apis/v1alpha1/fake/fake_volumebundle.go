/*
Copyright 2024 The Kubernetes Authors.

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

	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
	v1alpha1 "sigs.k8s.io/kueue/cmd/experimental/kjobctl/apis/v1alpha1"
)

// FakeVolumeBundles implements VolumeBundleInterface
type FakeVolumeBundles struct {
	Fake *FakeKjobctlV1alpha1
	ns   string
}

var volumebundlesResource = v1alpha1.SchemeGroupVersion.WithResource("volumebundles")

var volumebundlesKind = v1alpha1.SchemeGroupVersion.WithKind("VolumeBundle")

// Get takes name of the volumeBundle, and returns the corresponding volumeBundle object, and an error if there is any.
func (c *FakeVolumeBundles) Get(ctx context.Context, name string, options v1.GetOptions) (result *v1alpha1.VolumeBundle, err error) {
	emptyResult := &v1alpha1.VolumeBundle{}
	obj, err := c.Fake.
		Invokes(testing.NewGetActionWithOptions(volumebundlesResource, c.ns, name, options), emptyResult)

	if obj == nil {
		return emptyResult, err
	}
	return obj.(*v1alpha1.VolumeBundle), err
}

// List takes label and field selectors, and returns the list of VolumeBundles that match those selectors.
func (c *FakeVolumeBundles) List(ctx context.Context, opts v1.ListOptions) (result *v1alpha1.VolumeBundleList, err error) {
	emptyResult := &v1alpha1.VolumeBundleList{}
	obj, err := c.Fake.
		Invokes(testing.NewListActionWithOptions(volumebundlesResource, volumebundlesKind, c.ns, opts), emptyResult)

	if obj == nil {
		return emptyResult, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.VolumeBundleList{ListMeta: obj.(*v1alpha1.VolumeBundleList).ListMeta}
	for _, item := range obj.(*v1alpha1.VolumeBundleList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested volumeBundles.
func (c *FakeVolumeBundles) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchActionWithOptions(volumebundlesResource, c.ns, opts))

}

// Create takes the representation of a volumeBundle and creates it.  Returns the server's representation of the volumeBundle, and an error, if there is any.
func (c *FakeVolumeBundles) Create(ctx context.Context, volumeBundle *v1alpha1.VolumeBundle, opts v1.CreateOptions) (result *v1alpha1.VolumeBundle, err error) {
	emptyResult := &v1alpha1.VolumeBundle{}
	obj, err := c.Fake.
		Invokes(testing.NewCreateActionWithOptions(volumebundlesResource, c.ns, volumeBundle, opts), emptyResult)

	if obj == nil {
		return emptyResult, err
	}
	return obj.(*v1alpha1.VolumeBundle), err
}

// Update takes the representation of a volumeBundle and updates it. Returns the server's representation of the volumeBundle, and an error, if there is any.
func (c *FakeVolumeBundles) Update(ctx context.Context, volumeBundle *v1alpha1.VolumeBundle, opts v1.UpdateOptions) (result *v1alpha1.VolumeBundle, err error) {
	emptyResult := &v1alpha1.VolumeBundle{}
	obj, err := c.Fake.
		Invokes(testing.NewUpdateActionWithOptions(volumebundlesResource, c.ns, volumeBundle, opts), emptyResult)

	if obj == nil {
		return emptyResult, err
	}
	return obj.(*v1alpha1.VolumeBundle), err
}

// Delete takes name of the volumeBundle and deletes it. Returns an error if one occurs.
func (c *FakeVolumeBundles) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteActionWithOptions(volumebundlesResource, c.ns, name, opts), &v1alpha1.VolumeBundle{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeVolumeBundles) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	action := testing.NewDeleteCollectionActionWithOptions(volumebundlesResource, c.ns, opts, listOpts)

	_, err := c.Fake.Invokes(action, &v1alpha1.VolumeBundleList{})
	return err
}

// Patch applies the patch and returns the patched volumeBundle.
func (c *FakeVolumeBundles) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha1.VolumeBundle, err error) {
	emptyResult := &v1alpha1.VolumeBundle{}
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceActionWithOptions(volumebundlesResource, c.ns, name, pt, data, opts, subresources...), emptyResult)

	if obj == nil {
		return emptyResult, err
	}
	return obj.(*v1alpha1.VolumeBundle), err
}
