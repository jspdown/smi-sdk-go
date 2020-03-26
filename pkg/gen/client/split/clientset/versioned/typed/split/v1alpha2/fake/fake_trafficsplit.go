/*
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

	v1alpha2 "github.com/servicemeshinterface/smi-sdk-go/pkg/apis/split/v1alpha2"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
)

// FakeTrafficSplits implements TrafficSplitInterface
type FakeTrafficSplits struct {
	Fake *FakeSplitV1alpha2
	ns   string
}

var trafficsplitsResource = schema.GroupVersionResource{Group: "split.smi-spec.io", Version: "v1alpha2", Resource: "trafficsplits"}

var trafficsplitsKind = schema.GroupVersionKind{Group: "split.smi-spec.io", Version: "v1alpha2", Kind: "TrafficSplit"}

// Get takes name of the trafficSplit, and returns the corresponding trafficSplit object, and an error if there is any.
func (c *FakeTrafficSplits) Get(ctx context.Context, name string, options v1.GetOptions) (result *v1alpha2.TrafficSplit, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(trafficsplitsResource, c.ns, name), &v1alpha2.TrafficSplit{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha2.TrafficSplit), err
}

// List takes label and field selectors, and returns the list of TrafficSplits that match those selectors.
func (c *FakeTrafficSplits) List(ctx context.Context, opts v1.ListOptions) (result *v1alpha2.TrafficSplitList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(trafficsplitsResource, trafficsplitsKind, c.ns, opts), &v1alpha2.TrafficSplitList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha2.TrafficSplitList{ListMeta: obj.(*v1alpha2.TrafficSplitList).ListMeta}
	for _, item := range obj.(*v1alpha2.TrafficSplitList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested trafficSplits.
func (c *FakeTrafficSplits) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(trafficsplitsResource, c.ns, opts))

}

// Create takes the representation of a trafficSplit and creates it.  Returns the server's representation of the trafficSplit, and an error, if there is any.
func (c *FakeTrafficSplits) Create(ctx context.Context, trafficSplit *v1alpha2.TrafficSplit, opts v1.CreateOptions) (result *v1alpha2.TrafficSplit, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(trafficsplitsResource, c.ns, trafficSplit), &v1alpha2.TrafficSplit{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha2.TrafficSplit), err
}

// Update takes the representation of a trafficSplit and updates it. Returns the server's representation of the trafficSplit, and an error, if there is any.
func (c *FakeTrafficSplits) Update(ctx context.Context, trafficSplit *v1alpha2.TrafficSplit, opts v1.UpdateOptions) (result *v1alpha2.TrafficSplit, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(trafficsplitsResource, c.ns, trafficSplit), &v1alpha2.TrafficSplit{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha2.TrafficSplit), err
}

// Delete takes name of the trafficSplit and deletes it. Returns an error if one occurs.
func (c *FakeTrafficSplits) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteAction(trafficsplitsResource, c.ns, name), &v1alpha2.TrafficSplit{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeTrafficSplits) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(trafficsplitsResource, c.ns, listOpts)

	_, err := c.Fake.Invokes(action, &v1alpha2.TrafficSplitList{})
	return err
}

// Patch applies the patch and returns the patched trafficSplit.
func (c *FakeTrafficSplits) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha2.TrafficSplit, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(trafficsplitsResource, c.ns, name, pt, data, subresources...), &v1alpha2.TrafficSplit{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha2.TrafficSplit), err
}
