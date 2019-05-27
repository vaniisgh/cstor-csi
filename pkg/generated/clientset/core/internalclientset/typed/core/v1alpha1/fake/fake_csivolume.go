/*
Copyright 2019 The OpenEBS Authors

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
	v1alpha1 "github.com/openebs/csi/pkg/apis/openebs.io/core/v1alpha1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
)

// FakeCSIVolumes implements CSIVolumeInterface
type FakeCSIVolumes struct {
	Fake *FakeOpenebsV1alpha1
	ns   string
}

var csivolumesResource = schema.GroupVersionResource{Group: "openebs.io", Version: "v1alpha1", Resource: "csivolumes"}

var csivolumesKind = schema.GroupVersionKind{Group: "openebs.io", Version: "v1alpha1", Kind: "CSIVolume"}

// Get takes name of the cSIVolume, and returns the corresponding cSIVolume object, and an error if there is any.
func (c *FakeCSIVolumes) Get(name string, options v1.GetOptions) (result *v1alpha1.CSIVolume, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(csivolumesResource, c.ns, name), &v1alpha1.CSIVolume{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.CSIVolume), err
}

// List takes label and field selectors, and returns the list of CSIVolumes that match those selectors.
func (c *FakeCSIVolumes) List(opts v1.ListOptions) (result *v1alpha1.CSIVolumeList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(csivolumesResource, csivolumesKind, c.ns, opts), &v1alpha1.CSIVolumeList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.CSIVolumeList{ListMeta: obj.(*v1alpha1.CSIVolumeList).ListMeta}
	for _, item := range obj.(*v1alpha1.CSIVolumeList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested cSIVolumes.
func (c *FakeCSIVolumes) Watch(opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(csivolumesResource, c.ns, opts))

}

// Create takes the representation of a cSIVolume and creates it.  Returns the server's representation of the cSIVolume, and an error, if there is any.
func (c *FakeCSIVolumes) Create(cSIVolume *v1alpha1.CSIVolume) (result *v1alpha1.CSIVolume, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(csivolumesResource, c.ns, cSIVolume), &v1alpha1.CSIVolume{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.CSIVolume), err
}

// Update takes the representation of a cSIVolume and updates it. Returns the server's representation of the cSIVolume, and an error, if there is any.
func (c *FakeCSIVolumes) Update(cSIVolume *v1alpha1.CSIVolume) (result *v1alpha1.CSIVolume, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(csivolumesResource, c.ns, cSIVolume), &v1alpha1.CSIVolume{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.CSIVolume), err
}

// Delete takes name of the cSIVolume and deletes it. Returns an error if one occurs.
func (c *FakeCSIVolumes) Delete(name string, options *v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteAction(csivolumesResource, c.ns, name), &v1alpha1.CSIVolume{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeCSIVolumes) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(csivolumesResource, c.ns, listOptions)

	_, err := c.Fake.Invokes(action, &v1alpha1.CSIVolumeList{})
	return err
}

// Patch applies the patch and returns the patched cSIVolume.
func (c *FakeCSIVolumes) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.CSIVolume, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(csivolumesResource, c.ns, name, pt, data, subresources...), &v1alpha1.CSIVolume{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.CSIVolume), err
}