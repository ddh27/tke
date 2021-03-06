/*
 * Tencent is pleased to support the open source community by making TKEStack
 * available.
 *
 * Copyright (C) 2012-2019 Tencent. All Rights Reserved.
 *
 * Licensed under the Apache License, Version 2.0 (the "License"); you may not use
 * this file except in compliance with the License. You may obtain a copy of the
 * License at
 *
 * https://opensource.org/licenses/Apache-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS, WITHOUT
 * WARRANTIES OF ANY KIND, either express or implied.  See the License for the
 * specific language governing permissions and limitations under the License.
 */

// Code generated by client-gen. DO NOT EDIT.

package v1

import (
	"time"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	rest "k8s.io/client-go/rest"
	v1 "tkestack.io/tke/api/business/v1"
	scheme "tkestack.io/tke/api/client/clientset/versioned/scheme"
)

// ChartGroupsGetter has a method to return a ChartGroupInterface.
// A group's client should implement this interface.
type ChartGroupsGetter interface {
	ChartGroups(namespace string) ChartGroupInterface
}

// ChartGroupInterface has methods to work with ChartGroup resources.
type ChartGroupInterface interface {
	Create(*v1.ChartGroup) (*v1.ChartGroup, error)
	Update(*v1.ChartGroup) (*v1.ChartGroup, error)
	UpdateStatus(*v1.ChartGroup) (*v1.ChartGroup, error)
	Delete(name string, options *metav1.DeleteOptions) error
	DeleteCollection(options *metav1.DeleteOptions, listOptions metav1.ListOptions) error
	Get(name string, options metav1.GetOptions) (*v1.ChartGroup, error)
	List(opts metav1.ListOptions) (*v1.ChartGroupList, error)
	Watch(opts metav1.ListOptions) (watch.Interface, error)
	Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1.ChartGroup, err error)
	ChartGroupExpansion
}

// chartGroups implements ChartGroupInterface
type chartGroups struct {
	client rest.Interface
	ns     string
}

// newChartGroups returns a ChartGroups
func newChartGroups(c *BusinessV1Client, namespace string) *chartGroups {
	return &chartGroups{
		client: c.RESTClient(),
		ns:     namespace,
	}
}

// Get takes name of the chartGroup, and returns the corresponding chartGroup object, and an error if there is any.
func (c *chartGroups) Get(name string, options metav1.GetOptions) (result *v1.ChartGroup, err error) {
	result = &v1.ChartGroup{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("chartgroups").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do().
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of ChartGroups that match those selectors.
func (c *chartGroups) List(opts metav1.ListOptions) (result *v1.ChartGroupList, err error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	result = &v1.ChartGroupList{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("chartgroups").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Do().
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested chartGroups.
func (c *chartGroups) Watch(opts metav1.ListOptions) (watch.Interface, error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	opts.Watch = true
	return c.client.Get().
		Namespace(c.ns).
		Resource("chartgroups").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Watch()
}

// Create takes the representation of a chartGroup and creates it.  Returns the server's representation of the chartGroup, and an error, if there is any.
func (c *chartGroups) Create(chartGroup *v1.ChartGroup) (result *v1.ChartGroup, err error) {
	result = &v1.ChartGroup{}
	err = c.client.Post().
		Namespace(c.ns).
		Resource("chartgroups").
		Body(chartGroup).
		Do().
		Into(result)
	return
}

// Update takes the representation of a chartGroup and updates it. Returns the server's representation of the chartGroup, and an error, if there is any.
func (c *chartGroups) Update(chartGroup *v1.ChartGroup) (result *v1.ChartGroup, err error) {
	result = &v1.ChartGroup{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("chartgroups").
		Name(chartGroup.Name).
		Body(chartGroup).
		Do().
		Into(result)
	return
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().

func (c *chartGroups) UpdateStatus(chartGroup *v1.ChartGroup) (result *v1.ChartGroup, err error) {
	result = &v1.ChartGroup{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("chartgroups").
		Name(chartGroup.Name).
		SubResource("status").
		Body(chartGroup).
		Do().
		Into(result)
	return
}

// Delete takes name of the chartGroup and deletes it. Returns an error if one occurs.
func (c *chartGroups) Delete(name string, options *metav1.DeleteOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("chartgroups").
		Name(name).
		Body(options).
		Do().
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *chartGroups) DeleteCollection(options *metav1.DeleteOptions, listOptions metav1.ListOptions) error {
	var timeout time.Duration
	if listOptions.TimeoutSeconds != nil {
		timeout = time.Duration(*listOptions.TimeoutSeconds) * time.Second
	}
	return c.client.Delete().
		Namespace(c.ns).
		Resource("chartgroups").
		VersionedParams(&listOptions, scheme.ParameterCodec).
		Timeout(timeout).
		Body(options).
		Do().
		Error()
}

// Patch applies the patch and returns the patched chartGroup.
func (c *chartGroups) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1.ChartGroup, err error) {
	result = &v1.ChartGroup{}
	err = c.client.Patch(pt).
		Namespace(c.ns).
		Resource("chartgroups").
		SubResource(subresources...).
		Name(name).
		Body(data).
		Do().
		Into(result)
	return
}
