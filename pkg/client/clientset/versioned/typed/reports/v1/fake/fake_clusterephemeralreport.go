/*
Copyright The Kubernetes Authors.

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

	v1 "github.com/kyverno/kyverno/api/reports/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
)

// FakeClusterEphemeralReports implements ClusterEphemeralReportInterface
type FakeClusterEphemeralReports struct {
	Fake *FakeReportsV1
}

var clusterephemeralreportsResource = v1.SchemeGroupVersion.WithResource("clusterephemeralreports")

var clusterephemeralreportsKind = v1.SchemeGroupVersion.WithKind("ClusterEphemeralReport")

// Get takes name of the clusterEphemeralReport, and returns the corresponding clusterEphemeralReport object, and an error if there is any.
func (c *FakeClusterEphemeralReports) Get(ctx context.Context, name string, options metav1.GetOptions) (result *v1.ClusterEphemeralReport, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootGetAction(clusterephemeralreportsResource, name), &v1.ClusterEphemeralReport{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1.ClusterEphemeralReport), err
}

// List takes label and field selectors, and returns the list of ClusterEphemeralReports that match those selectors.
func (c *FakeClusterEphemeralReports) List(ctx context.Context, opts metav1.ListOptions) (result *v1.ClusterEphemeralReportList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootListAction(clusterephemeralreportsResource, clusterephemeralreportsKind, opts), &v1.ClusterEphemeralReportList{})
	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1.ClusterEphemeralReportList{ListMeta: obj.(*v1.ClusterEphemeralReportList).ListMeta}
	for _, item := range obj.(*v1.ClusterEphemeralReportList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested clusterEphemeralReports.
func (c *FakeClusterEphemeralReports) Watch(ctx context.Context, opts metav1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewRootWatchAction(clusterephemeralreportsResource, opts))
}

// Create takes the representation of a clusterEphemeralReport and creates it.  Returns the server's representation of the clusterEphemeralReport, and an error, if there is any.
func (c *FakeClusterEphemeralReports) Create(ctx context.Context, clusterEphemeralReport *v1.ClusterEphemeralReport, opts metav1.CreateOptions) (result *v1.ClusterEphemeralReport, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootCreateAction(clusterephemeralreportsResource, clusterEphemeralReport), &v1.ClusterEphemeralReport{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1.ClusterEphemeralReport), err
}

// Update takes the representation of a clusterEphemeralReport and updates it. Returns the server's representation of the clusterEphemeralReport, and an error, if there is any.
func (c *FakeClusterEphemeralReports) Update(ctx context.Context, clusterEphemeralReport *v1.ClusterEphemeralReport, opts metav1.UpdateOptions) (result *v1.ClusterEphemeralReport, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootUpdateAction(clusterephemeralreportsResource, clusterEphemeralReport), &v1.ClusterEphemeralReport{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1.ClusterEphemeralReport), err
}

// Delete takes name of the clusterEphemeralReport and deletes it. Returns an error if one occurs.
func (c *FakeClusterEphemeralReports) Delete(ctx context.Context, name string, opts metav1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewRootDeleteActionWithOptions(clusterephemeralreportsResource, name, opts), &v1.ClusterEphemeralReport{})
	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeClusterEphemeralReports) DeleteCollection(ctx context.Context, opts metav1.DeleteOptions, listOpts metav1.ListOptions) error {
	action := testing.NewRootDeleteCollectionAction(clusterephemeralreportsResource, listOpts)

	_, err := c.Fake.Invokes(action, &v1.ClusterEphemeralReportList{})
	return err
}

// Patch applies the patch and returns the patched clusterEphemeralReport.
func (c *FakeClusterEphemeralReports) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts metav1.PatchOptions, subresources ...string) (result *v1.ClusterEphemeralReport, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootPatchSubresourceAction(clusterephemeralreportsResource, name, pt, data, subresources...), &v1.ClusterEphemeralReport{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1.ClusterEphemeralReport), err
}