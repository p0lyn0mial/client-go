//go:build !ignore_autogenerated
// +build !ignore_autogenerated

/*
Copyright The KCP Authors.

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

// Code generated by kcp code-generator. DO NOT EDIT.

package v1

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/kcp-dev/logicalcluster/v2"

	eventsv1 "k8s.io/api/events/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/apimachinery/pkg/runtime/schema"
	"k8s.io/apimachinery/pkg/types"
	"k8s.io/apimachinery/pkg/watch"
	applyconfigurationseventsv1 "k8s.io/client-go/applyconfigurations/events/v1"
	eventsv1client "k8s.io/client-go/kubernetes/typed/events/v1"
	"k8s.io/client-go/testing"

	kcpeventsv1 "github.com/kcp-dev/client-go/clients/clientset/versioned/typed/events/v1"
	kcptesting "github.com/kcp-dev/client-go/third_party/k8s.io/client-go/testing"
)

var eventsResource = schema.GroupVersionResource{Group: "events.k8s.io", Version: "v1", Resource: "events"}
var eventsKind = schema.GroupVersionKind{Group: "events.k8s.io", Version: "v1", Kind: "Event"}

type eventsClusterClient struct {
	*kcptesting.Fake
}

// Cluster scopes the client down to a particular cluster.
func (c *eventsClusterClient) Cluster(cluster logicalcluster.Name) kcpeventsv1.EventsNamespacer {
	if cluster == logicalcluster.Wildcard {
		panic("A specific cluster must be provided when scoping, not the wildcard.")
	}

	return &eventsNamespacer{Fake: c.Fake, Cluster: cluster}
}

// List takes label and field selectors, and returns the list of Events that match those selectors across all clusters.
func (c *eventsClusterClient) List(ctx context.Context, opts metav1.ListOptions) (*eventsv1.EventList, error) {
	obj, err := c.Fake.Invokes(kcptesting.NewListAction(eventsResource, eventsKind, logicalcluster.Wildcard, metav1.NamespaceAll, opts), &eventsv1.EventList{})
	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &eventsv1.EventList{ListMeta: obj.(*eventsv1.EventList).ListMeta}
	for _, item := range obj.(*eventsv1.EventList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested Events across all clusters.
func (c *eventsClusterClient) Watch(ctx context.Context, opts metav1.ListOptions) (watch.Interface, error) {
	return c.Fake.InvokesWatch(kcptesting.NewWatchAction(eventsResource, logicalcluster.Wildcard, metav1.NamespaceAll, opts))
}

type eventsNamespacer struct {
	*kcptesting.Fake
	Cluster logicalcluster.Name
}

func (n *eventsNamespacer) Namespace(namespace string) eventsv1client.EventInterface {
	return &eventsClient{Fake: n.Fake, Cluster: n.Cluster, Namespace: namespace}
}

type eventsClient struct {
	*kcptesting.Fake
	Cluster   logicalcluster.Name
	Namespace string
}

func (c *eventsClient) Create(ctx context.Context, event *eventsv1.Event, opts metav1.CreateOptions) (*eventsv1.Event, error) {
	obj, err := c.Fake.Invokes(kcptesting.NewCreateAction(eventsResource, c.Cluster, c.Namespace, event), &eventsv1.Event{})
	if obj == nil {
		return nil, err
	}
	return obj.(*eventsv1.Event), err
}

func (c *eventsClient) Update(ctx context.Context, event *eventsv1.Event, opts metav1.UpdateOptions) (*eventsv1.Event, error) {
	obj, err := c.Fake.Invokes(kcptesting.NewUpdateAction(eventsResource, c.Cluster, c.Namespace, event), &eventsv1.Event{})
	if obj == nil {
		return nil, err
	}
	return obj.(*eventsv1.Event), err
}

func (c *eventsClient) UpdateStatus(ctx context.Context, event *eventsv1.Event, opts metav1.UpdateOptions) (*eventsv1.Event, error) {
	obj, err := c.Fake.Invokes(kcptesting.NewUpdateSubresourceAction(eventsResource, c.Cluster, "status", c.Namespace, event), &eventsv1.Event{})
	if obj == nil {
		return nil, err
	}
	return obj.(*eventsv1.Event), err
}

func (c *eventsClient) Delete(ctx context.Context, name string, opts metav1.DeleteOptions) error {
	_, err := c.Fake.Invokes(kcptesting.NewDeleteActionWithOptions(eventsResource, c.Cluster, c.Namespace, name, opts), &eventsv1.Event{})
	return err
}

func (c *eventsClient) DeleteCollection(ctx context.Context, opts metav1.DeleteOptions, listOpts metav1.ListOptions) error {
	action := kcptesting.NewDeleteCollectionAction(eventsResource, c.Cluster, c.Namespace, listOpts)

	_, err := c.Fake.Invokes(action, &eventsv1.EventList{})
	return err
}

func (c *eventsClient) Get(ctx context.Context, name string, options metav1.GetOptions) (*eventsv1.Event, error) {
	obj, err := c.Fake.Invokes(kcptesting.NewGetAction(eventsResource, c.Cluster, c.Namespace, name), &eventsv1.Event{})
	if obj == nil {
		return nil, err
	}
	return obj.(*eventsv1.Event), err
}

// List takes label and field selectors, and returns the list of Events that match those selectors.
func (c *eventsClient) List(ctx context.Context, opts metav1.ListOptions) (*eventsv1.EventList, error) {
	obj, err := c.Fake.Invokes(kcptesting.NewListAction(eventsResource, eventsKind, c.Cluster, c.Namespace, opts), &eventsv1.EventList{})
	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &eventsv1.EventList{ListMeta: obj.(*eventsv1.EventList).ListMeta}
	for _, item := range obj.(*eventsv1.EventList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

func (c *eventsClient) Watch(ctx context.Context, opts metav1.ListOptions) (watch.Interface, error) {
	return c.Fake.InvokesWatch(kcptesting.NewWatchAction(eventsResource, c.Cluster, c.Namespace, opts))
}

func (c *eventsClient) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts metav1.PatchOptions, subresources ...string) (*eventsv1.Event, error) {
	obj, err := c.Fake.Invokes(kcptesting.NewPatchSubresourceAction(eventsResource, c.Cluster, c.Namespace, name, pt, data, subresources...), &eventsv1.Event{})
	if obj == nil {
		return nil, err
	}
	return obj.(*eventsv1.Event), err
}

func (c *eventsClient) Apply(ctx context.Context, applyConfiguration *applyconfigurationseventsv1.EventApplyConfiguration, opts metav1.ApplyOptions) (*eventsv1.Event, error) {
	if applyConfiguration == nil {
		return nil, fmt.Errorf("applyConfiguration provided to Apply must not be nil")
	}
	data, err := json.Marshal(applyConfiguration)
	if err != nil {
		return nil, err
	}
	name := applyConfiguration.Name
	if name == nil {
		return nil, fmt.Errorf("applyConfiguration.Name must be provided to Apply")
	}
	obj, err := c.Fake.Invokes(kcptesting.NewPatchSubresourceAction(eventsResource, c.Cluster, c.Namespace, *name, types.ApplyPatchType, data), &eventsv1.Event{})
	if obj == nil {
		return nil, err
	}
	return obj.(*eventsv1.Event), err
}

func (c *eventsClient) ApplyStatus(ctx context.Context, applyConfiguration *applyconfigurationseventsv1.EventApplyConfiguration, opts metav1.ApplyOptions) (*eventsv1.Event, error) {
	if applyConfiguration == nil {
		return nil, fmt.Errorf("applyConfiguration provided to Apply must not be nil")
	}
	data, err := json.Marshal(applyConfiguration)
	if err != nil {
		return nil, err
	}
	name := applyConfiguration.Name
	if name == nil {
		return nil, fmt.Errorf("applyConfiguration.Name must be provided to Apply")
	}
	obj, err := c.Fake.Invokes(kcptesting.NewPatchSubresourceAction(eventsResource, c.Cluster, c.Namespace, *name, types.ApplyPatchType, data, "status"), &eventsv1.Event{})
	if obj == nil {
		return nil, err
	}
	return obj.(*eventsv1.Event), err
}
