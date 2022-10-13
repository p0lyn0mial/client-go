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

package v1beta2

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/kcp-dev/logicalcluster/v2"

	flowcontrolv1beta2 "k8s.io/api/flowcontrol/v1beta2"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/apimachinery/pkg/runtime/schema"
	"k8s.io/apimachinery/pkg/types"
	"k8s.io/apimachinery/pkg/watch"
	applyconfigurationsflowcontrolv1beta2 "k8s.io/client-go/applyconfigurations/flowcontrol/v1beta2"
	flowcontrolv1beta2client "k8s.io/client-go/kubernetes/typed/flowcontrol/v1beta2"
	"k8s.io/client-go/testing"

	kcptesting "github.com/kcp-dev/client-go/third_party/k8s.io/client-go/testing"
)

var flowSchemasResource = schema.GroupVersionResource{Group: "flowcontrol.apiserver.k8s.io", Version: "v1beta2", Resource: "flowschemas"}
var flowSchemasKind = schema.GroupVersionKind{Group: "flowcontrol.apiserver.k8s.io", Version: "v1beta2", Kind: "FlowSchema"}

type flowSchemasClusterClient struct {
	*kcptesting.Fake
}

// Cluster scopes the client down to a particular cluster.
func (c *flowSchemasClusterClient) Cluster(cluster logicalcluster.Name) flowcontrolv1beta2client.FlowSchemaInterface {
	if cluster == logicalcluster.Wildcard {
		panic("A specific cluster must be provided when scoping, not the wildcard.")
	}

	return &flowSchemasClient{Fake: c.Fake, Cluster: cluster}
}

// List takes label and field selectors, and returns the list of FlowSchemas that match those selectors across all clusters.
func (c *flowSchemasClusterClient) List(ctx context.Context, opts metav1.ListOptions) (*flowcontrolv1beta2.FlowSchemaList, error) {
	obj, err := c.Fake.Invokes(kcptesting.NewRootListAction(flowSchemasResource, flowSchemasKind, logicalcluster.Wildcard, opts), &flowcontrolv1beta2.FlowSchemaList{})
	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &flowcontrolv1beta2.FlowSchemaList{ListMeta: obj.(*flowcontrolv1beta2.FlowSchemaList).ListMeta}
	for _, item := range obj.(*flowcontrolv1beta2.FlowSchemaList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested FlowSchemas across all clusters.
func (c *flowSchemasClusterClient) Watch(ctx context.Context, opts metav1.ListOptions) (watch.Interface, error) {
	return c.Fake.InvokesWatch(kcptesting.NewRootWatchAction(flowSchemasResource, logicalcluster.Wildcard, opts))
}

type flowSchemasClient struct {
	*kcptesting.Fake
	Cluster logicalcluster.Name
}

func (c *flowSchemasClient) Create(ctx context.Context, flowSchema *flowcontrolv1beta2.FlowSchema, opts metav1.CreateOptions) (*flowcontrolv1beta2.FlowSchema, error) {
	obj, err := c.Fake.Invokes(kcptesting.NewRootCreateAction(flowSchemasResource, c.Cluster, flowSchema), &flowcontrolv1beta2.FlowSchema{})
	if obj == nil {
		return nil, err
	}
	return obj.(*flowcontrolv1beta2.FlowSchema), err
}

func (c *flowSchemasClient) Update(ctx context.Context, flowSchema *flowcontrolv1beta2.FlowSchema, opts metav1.UpdateOptions) (*flowcontrolv1beta2.FlowSchema, error) {
	obj, err := c.Fake.Invokes(kcptesting.NewRootUpdateAction(flowSchemasResource, c.Cluster, flowSchema), &flowcontrolv1beta2.FlowSchema{})
	if obj == nil {
		return nil, err
	}
	return obj.(*flowcontrolv1beta2.FlowSchema), err
}

func (c *flowSchemasClient) UpdateStatus(ctx context.Context, flowSchema *flowcontrolv1beta2.FlowSchema, opts metav1.UpdateOptions) (*flowcontrolv1beta2.FlowSchema, error) {
	obj, err := c.Fake.Invokes(kcptesting.NewRootUpdateSubresourceAction(flowSchemasResource, c.Cluster, "status", flowSchema), &flowcontrolv1beta2.FlowSchema{})
	if obj == nil {
		return nil, err
	}
	return obj.(*flowcontrolv1beta2.FlowSchema), err
}

func (c *flowSchemasClient) Delete(ctx context.Context, name string, opts metav1.DeleteOptions) error {
	_, err := c.Fake.Invokes(kcptesting.NewRootDeleteActionWithOptions(flowSchemasResource, c.Cluster, name, opts), &flowcontrolv1beta2.FlowSchema{})
	return err
}

func (c *flowSchemasClient) DeleteCollection(ctx context.Context, opts metav1.DeleteOptions, listOpts metav1.ListOptions) error {
	action := kcptesting.NewRootDeleteCollectionAction(flowSchemasResource, c.Cluster, listOpts)

	_, err := c.Fake.Invokes(action, &flowcontrolv1beta2.FlowSchemaList{})
	return err
}

func (c *flowSchemasClient) Get(ctx context.Context, name string, options metav1.GetOptions) (*flowcontrolv1beta2.FlowSchema, error) {
	obj, err := c.Fake.Invokes(kcptesting.NewRootGetAction(flowSchemasResource, c.Cluster, name), &flowcontrolv1beta2.FlowSchema{})
	if obj == nil {
		return nil, err
	}
	return obj.(*flowcontrolv1beta2.FlowSchema), err
}

// List takes label and field selectors, and returns the list of FlowSchemas that match those selectors.
func (c *flowSchemasClient) List(ctx context.Context, opts metav1.ListOptions) (*flowcontrolv1beta2.FlowSchemaList, error) {
	obj, err := c.Fake.Invokes(kcptesting.NewRootListAction(flowSchemasResource, flowSchemasKind, c.Cluster, opts), &flowcontrolv1beta2.FlowSchemaList{})
	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &flowcontrolv1beta2.FlowSchemaList{ListMeta: obj.(*flowcontrolv1beta2.FlowSchemaList).ListMeta}
	for _, item := range obj.(*flowcontrolv1beta2.FlowSchemaList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

func (c *flowSchemasClient) Watch(ctx context.Context, opts metav1.ListOptions) (watch.Interface, error) {
	return c.Fake.InvokesWatch(kcptesting.NewRootWatchAction(flowSchemasResource, c.Cluster, opts))
}

func (c *flowSchemasClient) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts metav1.PatchOptions, subresources ...string) (*flowcontrolv1beta2.FlowSchema, error) {
	obj, err := c.Fake.Invokes(kcptesting.NewRootPatchSubresourceAction(flowSchemasResource, c.Cluster, name, pt, data, subresources...), &flowcontrolv1beta2.FlowSchema{})
	if obj == nil {
		return nil, err
	}
	return obj.(*flowcontrolv1beta2.FlowSchema), err
}

func (c *flowSchemasClient) Apply(ctx context.Context, applyConfiguration *applyconfigurationsflowcontrolv1beta2.FlowSchemaApplyConfiguration, opts metav1.ApplyOptions) (*flowcontrolv1beta2.FlowSchema, error) {
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
	obj, err := c.Fake.Invokes(kcptesting.NewRootPatchSubresourceAction(flowSchemasResource, c.Cluster, *name, types.ApplyPatchType, data), &flowcontrolv1beta2.FlowSchema{})
	if obj == nil {
		return nil, err
	}
	return obj.(*flowcontrolv1beta2.FlowSchema), err
}

func (c *flowSchemasClient) ApplyStatus(ctx context.Context, applyConfiguration *applyconfigurationsflowcontrolv1beta2.FlowSchemaApplyConfiguration, opts metav1.ApplyOptions) (*flowcontrolv1beta2.FlowSchema, error) {
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
	obj, err := c.Fake.Invokes(kcptesting.NewRootPatchSubresourceAction(flowSchemasResource, c.Cluster, *name, types.ApplyPatchType, data, "status"), &flowcontrolv1beta2.FlowSchema{})
	if obj == nil {
		return nil, err
	}
	return obj.(*flowcontrolv1beta2.FlowSchema), err
}
