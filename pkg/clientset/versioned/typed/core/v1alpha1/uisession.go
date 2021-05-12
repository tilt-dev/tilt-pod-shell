/*
Copyright 2020 The Tilt Dev Authors

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

package v1alpha1

import (
	"context"
	"time"

	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	rest "k8s.io/client-go/rest"

	v1alpha1 "github.com/tilt-dev/tilt/pkg/apis/core/v1alpha1"
	scheme "github.com/tilt-dev/tilt-pod-shell/pkg/clientset/versioned/scheme"
)

// UISessionsGetter has a method to return a UISessionInterface.
// A group's client should implement this interface.
type UISessionsGetter interface {
	UISessions() UISessionInterface
}

// UISessionInterface has methods to work with UISession resources.
type UISessionInterface interface {
	Create(ctx context.Context, uISession *v1alpha1.UISession, opts v1.CreateOptions) (*v1alpha1.UISession, error)
	Update(ctx context.Context, uISession *v1alpha1.UISession, opts v1.UpdateOptions) (*v1alpha1.UISession, error)
	UpdateStatus(ctx context.Context, uISession *v1alpha1.UISession, opts v1.UpdateOptions) (*v1alpha1.UISession, error)
	Delete(ctx context.Context, name string, opts v1.DeleteOptions) error
	DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error
	Get(ctx context.Context, name string, opts v1.GetOptions) (*v1alpha1.UISession, error)
	List(ctx context.Context, opts v1.ListOptions) (*v1alpha1.UISessionList, error)
	Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error)
	Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha1.UISession, err error)
	UISessionExpansion
}

// uISessions implements UISessionInterface
type uISessions struct {
	client rest.Interface
}

// newUISessions returns a UISessions
func newUISessions(c *TiltV1alpha1Client) *uISessions {
	return &uISessions{
		client: c.RESTClient(),
	}
}

// Get takes name of the uISession, and returns the corresponding uISession object, and an error if there is any.
func (c *uISessions) Get(ctx context.Context, name string, options v1.GetOptions) (result *v1alpha1.UISession, err error) {
	result = &v1alpha1.UISession{}
	err = c.client.Get().
		Resource("uisessions").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do(ctx).
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of UISessions that match those selectors.
func (c *uISessions) List(ctx context.Context, opts v1.ListOptions) (result *v1alpha1.UISessionList, err error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	result = &v1alpha1.UISessionList{}
	err = c.client.Get().
		Resource("uisessions").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Do(ctx).
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested uISessions.
func (c *uISessions) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	opts.Watch = true
	return c.client.Get().
		Resource("uisessions").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Watch(ctx)
}

// Create takes the representation of a uISession and creates it.  Returns the server's representation of the uISession, and an error, if there is any.
func (c *uISessions) Create(ctx context.Context, uISession *v1alpha1.UISession, opts v1.CreateOptions) (result *v1alpha1.UISession, err error) {
	result = &v1alpha1.UISession{}
	err = c.client.Post().
		Resource("uisessions").
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(uISession).
		Do(ctx).
		Into(result)
	return
}

// Update takes the representation of a uISession and updates it. Returns the server's representation of the uISession, and an error, if there is any.
func (c *uISessions) Update(ctx context.Context, uISession *v1alpha1.UISession, opts v1.UpdateOptions) (result *v1alpha1.UISession, err error) {
	result = &v1alpha1.UISession{}
	err = c.client.Put().
		Resource("uisessions").
		Name(uISession.Name).
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(uISession).
		Do(ctx).
		Into(result)
	return
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *uISessions) UpdateStatus(ctx context.Context, uISession *v1alpha1.UISession, opts v1.UpdateOptions) (result *v1alpha1.UISession, err error) {
	result = &v1alpha1.UISession{}
	err = c.client.Put().
		Resource("uisessions").
		Name(uISession.Name).
		SubResource("status").
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(uISession).
		Do(ctx).
		Into(result)
	return
}

// Delete takes name of the uISession and deletes it. Returns an error if one occurs.
func (c *uISessions) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	return c.client.Delete().
		Resource("uisessions").
		Name(name).
		Body(&opts).
		Do(ctx).
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *uISessions) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	var timeout time.Duration
	if listOpts.TimeoutSeconds != nil {
		timeout = time.Duration(*listOpts.TimeoutSeconds) * time.Second
	}
	return c.client.Delete().
		Resource("uisessions").
		VersionedParams(&listOpts, scheme.ParameterCodec).
		Timeout(timeout).
		Body(&opts).
		Do(ctx).
		Error()
}

// Patch applies the patch and returns the patched uISession.
func (c *uISessions) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha1.UISession, err error) {
	result = &v1alpha1.UISession{}
	err = c.client.Patch(pt).
		Resource("uisessions").
		Name(name).
		SubResource(subresources...).
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(data).
		Do(ctx).
		Into(result)
	return
}
