/*
Copyright 2019 Compose, Zalando SE

Permission is hereby granted, free of charge, to any person obtaining a copy
of this software and associated documentation files (the "Software"), to deal
in the Software without restriction, including without limitation the rights
to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
copies of the Software, and to permit persons to whom the Software is
furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in all
copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
SOFTWARE.
*/

// Code generated by client-gen. DO NOT EDIT.

package fake

import (
	acidzalandov1 "github.com/zalando-incubator/postgres-operator/pkg/apis/acid.zalan.do/v1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
)

// FakePostgresqls implements PostgresqlInterface
type FakePostgresqls struct {
	Fake *FakeAcidV1
	ns   string
}

var postgresqlsResource = schema.GroupVersionResource{Group: "acid.zalan.do", Version: "v1", Resource: "postgresqls"}

var postgresqlsKind = schema.GroupVersionKind{Group: "acid.zalan.do", Version: "v1", Kind: "Postgresql"}

// Get takes name of the postgresql, and returns the corresponding postgresql object, and an error if there is any.
func (c *FakePostgresqls) Get(name string, options v1.GetOptions) (result *acidzalandov1.Postgresql, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(postgresqlsResource, c.ns, name), &acidzalandov1.Postgresql{})

	if obj == nil {
		return nil, err
	}
	return obj.(*acidzalandov1.Postgresql), err
}

// List takes label and field selectors, and returns the list of Postgresqls that match those selectors.
func (c *FakePostgresqls) List(opts v1.ListOptions) (result *acidzalandov1.PostgresqlList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(postgresqlsResource, postgresqlsKind, c.ns, opts), &acidzalandov1.PostgresqlList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &acidzalandov1.PostgresqlList{ListMeta: obj.(*acidzalandov1.PostgresqlList).ListMeta}
	for _, item := range obj.(*acidzalandov1.PostgresqlList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested postgresqls.
func (c *FakePostgresqls) Watch(opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(postgresqlsResource, c.ns, opts))

}

// Create takes the representation of a postgresql and creates it.  Returns the server's representation of the postgresql, and an error, if there is any.
func (c *FakePostgresqls) Create(postgresql *acidzalandov1.Postgresql) (result *acidzalandov1.Postgresql, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(postgresqlsResource, c.ns, postgresql), &acidzalandov1.Postgresql{})

	if obj == nil {
		return nil, err
	}
	return obj.(*acidzalandov1.Postgresql), err
}

// Update takes the representation of a postgresql and updates it. Returns the server's representation of the postgresql, and an error, if there is any.
func (c *FakePostgresqls) Update(postgresql *acidzalandov1.Postgresql) (result *acidzalandov1.Postgresql, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(postgresqlsResource, c.ns, postgresql), &acidzalandov1.Postgresql{})

	if obj == nil {
		return nil, err
	}
	return obj.(*acidzalandov1.Postgresql), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakePostgresqls) UpdateStatus(postgresql *acidzalandov1.Postgresql) (*acidzalandov1.Postgresql, error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateSubresourceAction(postgresqlsResource, "status", c.ns, postgresql), &acidzalandov1.Postgresql{})

	if obj == nil {
		return nil, err
	}
	return obj.(*acidzalandov1.Postgresql), err
}

// Delete takes name of the postgresql and deletes it. Returns an error if one occurs.
func (c *FakePostgresqls) Delete(name string, options *v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteAction(postgresqlsResource, c.ns, name), &acidzalandov1.Postgresql{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakePostgresqls) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(postgresqlsResource, c.ns, listOptions)

	_, err := c.Fake.Invokes(action, &acidzalandov1.PostgresqlList{})
	return err
}

// Patch applies the patch and returns the patched postgresql.
func (c *FakePostgresqls) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *acidzalandov1.Postgresql, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(postgresqlsResource, c.ns, name, data, subresources...), &acidzalandov1.Postgresql{})

	if obj == nil {
		return nil, err
	}
	return obj.(*acidzalandov1.Postgresql), err
}
