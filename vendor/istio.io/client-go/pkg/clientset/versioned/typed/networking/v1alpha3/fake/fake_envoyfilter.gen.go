// Copyright Istio Authors
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//    http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// Code generated by client-gen. DO NOT EDIT.

package fake

import (
	"context"
	json "encoding/json"
	"fmt"

	v1alpha3 "istio.io/client-go/pkg/apis/networking/v1alpha3"
	networkingv1alpha3 "istio.io/client-go/pkg/applyconfiguration/networking/v1alpha3"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
)

// FakeEnvoyFilters implements EnvoyFilterInterface
type FakeEnvoyFilters struct {
	Fake *FakeNetworkingV1alpha3
	ns   string
}

var envoyfiltersResource = v1alpha3.SchemeGroupVersion.WithResource("envoyfilters")

var envoyfiltersKind = v1alpha3.SchemeGroupVersion.WithKind("EnvoyFilter")

// Get takes name of the envoyFilter, and returns the corresponding envoyFilter object, and an error if there is any.
func (c *FakeEnvoyFilters) Get(ctx context.Context, name string, options v1.GetOptions) (result *v1alpha3.EnvoyFilter, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(envoyfiltersResource, c.ns, name), &v1alpha3.EnvoyFilter{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha3.EnvoyFilter), err
}

// List takes label and field selectors, and returns the list of EnvoyFilters that match those selectors.
func (c *FakeEnvoyFilters) List(ctx context.Context, opts v1.ListOptions) (result *v1alpha3.EnvoyFilterList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(envoyfiltersResource, envoyfiltersKind, c.ns, opts), &v1alpha3.EnvoyFilterList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha3.EnvoyFilterList{ListMeta: obj.(*v1alpha3.EnvoyFilterList).ListMeta}
	for _, item := range obj.(*v1alpha3.EnvoyFilterList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested envoyFilters.
func (c *FakeEnvoyFilters) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(envoyfiltersResource, c.ns, opts))

}

// Create takes the representation of a envoyFilter and creates it.  Returns the server's representation of the envoyFilter, and an error, if there is any.
func (c *FakeEnvoyFilters) Create(ctx context.Context, envoyFilter *v1alpha3.EnvoyFilter, opts v1.CreateOptions) (result *v1alpha3.EnvoyFilter, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(envoyfiltersResource, c.ns, envoyFilter), &v1alpha3.EnvoyFilter{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha3.EnvoyFilter), err
}

// Update takes the representation of a envoyFilter and updates it. Returns the server's representation of the envoyFilter, and an error, if there is any.
func (c *FakeEnvoyFilters) Update(ctx context.Context, envoyFilter *v1alpha3.EnvoyFilter, opts v1.UpdateOptions) (result *v1alpha3.EnvoyFilter, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(envoyfiltersResource, c.ns, envoyFilter), &v1alpha3.EnvoyFilter{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha3.EnvoyFilter), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeEnvoyFilters) UpdateStatus(ctx context.Context, envoyFilter *v1alpha3.EnvoyFilter, opts v1.UpdateOptions) (*v1alpha3.EnvoyFilter, error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateSubresourceAction(envoyfiltersResource, "status", c.ns, envoyFilter), &v1alpha3.EnvoyFilter{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha3.EnvoyFilter), err
}

// Delete takes name of the envoyFilter and deletes it. Returns an error if one occurs.
func (c *FakeEnvoyFilters) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteActionWithOptions(envoyfiltersResource, c.ns, name, opts), &v1alpha3.EnvoyFilter{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeEnvoyFilters) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(envoyfiltersResource, c.ns, listOpts)

	_, err := c.Fake.Invokes(action, &v1alpha3.EnvoyFilterList{})
	return err
}

// Patch applies the patch and returns the patched envoyFilter.
func (c *FakeEnvoyFilters) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha3.EnvoyFilter, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(envoyfiltersResource, c.ns, name, pt, data, subresources...), &v1alpha3.EnvoyFilter{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha3.EnvoyFilter), err
}

// Apply takes the given apply declarative configuration, applies it and returns the applied envoyFilter.
func (c *FakeEnvoyFilters) Apply(ctx context.Context, envoyFilter *networkingv1alpha3.EnvoyFilterApplyConfiguration, opts v1.ApplyOptions) (result *v1alpha3.EnvoyFilter, err error) {
	if envoyFilter == nil {
		return nil, fmt.Errorf("envoyFilter provided to Apply must not be nil")
	}
	data, err := json.Marshal(envoyFilter)
	if err != nil {
		return nil, err
	}
	name := envoyFilter.Name
	if name == nil {
		return nil, fmt.Errorf("envoyFilter.Name must be provided to Apply")
	}
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(envoyfiltersResource, c.ns, *name, types.ApplyPatchType, data), &v1alpha3.EnvoyFilter{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha3.EnvoyFilter), err
}

// ApplyStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating ApplyStatus().
func (c *FakeEnvoyFilters) ApplyStatus(ctx context.Context, envoyFilter *networkingv1alpha3.EnvoyFilterApplyConfiguration, opts v1.ApplyOptions) (result *v1alpha3.EnvoyFilter, err error) {
	if envoyFilter == nil {
		return nil, fmt.Errorf("envoyFilter provided to Apply must not be nil")
	}
	data, err := json.Marshal(envoyFilter)
	if err != nil {
		return nil, err
	}
	name := envoyFilter.Name
	if name == nil {
		return nil, fmt.Errorf("envoyFilter.Name must be provided to Apply")
	}
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(envoyfiltersResource, c.ns, *name, types.ApplyPatchType, data, "status"), &v1alpha3.EnvoyFilter{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha3.EnvoyFilter), err
}
