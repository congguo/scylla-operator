// Code generated by client-gen. DO NOT EDIT.

package fake

import (
	"context"

	v1alpha1 "github.com/scylladb/scylla-operator/pkg/api/scylla/v1alpha1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
)

// FakeRemoteOwners implements RemoteOwnerInterface
type FakeRemoteOwners struct {
	Fake *FakeScyllaV1alpha1
	ns   string
}

var remoteownersResource = v1alpha1.SchemeGroupVersion.WithResource("remoteowners")

var remoteownersKind = v1alpha1.SchemeGroupVersion.WithKind("RemoteOwner")

// Get takes name of the remoteOwner, and returns the corresponding remoteOwner object, and an error if there is any.
func (c *FakeRemoteOwners) Get(ctx context.Context, name string, options v1.GetOptions) (result *v1alpha1.RemoteOwner, err error) {
	emptyResult := &v1alpha1.RemoteOwner{}
	obj, err := c.Fake.
		Invokes(testing.NewGetActionWithOptions(remoteownersResource, c.ns, name, options), emptyResult)

	if obj == nil {
		return emptyResult, err
	}
	return obj.(*v1alpha1.RemoteOwner), err
}

// List takes label and field selectors, and returns the list of RemoteOwners that match those selectors.
func (c *FakeRemoteOwners) List(ctx context.Context, opts v1.ListOptions) (result *v1alpha1.RemoteOwnerList, err error) {
	emptyResult := &v1alpha1.RemoteOwnerList{}
	obj, err := c.Fake.
		Invokes(testing.NewListActionWithOptions(remoteownersResource, remoteownersKind, c.ns, opts), emptyResult)

	if obj == nil {
		return emptyResult, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.RemoteOwnerList{ListMeta: obj.(*v1alpha1.RemoteOwnerList).ListMeta}
	for _, item := range obj.(*v1alpha1.RemoteOwnerList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested remoteOwners.
func (c *FakeRemoteOwners) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchActionWithOptions(remoteownersResource, c.ns, opts))

}

// Create takes the representation of a remoteOwner and creates it.  Returns the server's representation of the remoteOwner, and an error, if there is any.
func (c *FakeRemoteOwners) Create(ctx context.Context, remoteOwner *v1alpha1.RemoteOwner, opts v1.CreateOptions) (result *v1alpha1.RemoteOwner, err error) {
	emptyResult := &v1alpha1.RemoteOwner{}
	obj, err := c.Fake.
		Invokes(testing.NewCreateActionWithOptions(remoteownersResource, c.ns, remoteOwner, opts), emptyResult)

	if obj == nil {
		return emptyResult, err
	}
	return obj.(*v1alpha1.RemoteOwner), err
}

// Update takes the representation of a remoteOwner and updates it. Returns the server's representation of the remoteOwner, and an error, if there is any.
func (c *FakeRemoteOwners) Update(ctx context.Context, remoteOwner *v1alpha1.RemoteOwner, opts v1.UpdateOptions) (result *v1alpha1.RemoteOwner, err error) {
	emptyResult := &v1alpha1.RemoteOwner{}
	obj, err := c.Fake.
		Invokes(testing.NewUpdateActionWithOptions(remoteownersResource, c.ns, remoteOwner, opts), emptyResult)

	if obj == nil {
		return emptyResult, err
	}
	return obj.(*v1alpha1.RemoteOwner), err
}

// Delete takes name of the remoteOwner and deletes it. Returns an error if one occurs.
func (c *FakeRemoteOwners) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteActionWithOptions(remoteownersResource, c.ns, name, opts), &v1alpha1.RemoteOwner{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeRemoteOwners) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	action := testing.NewDeleteCollectionActionWithOptions(remoteownersResource, c.ns, opts, listOpts)

	_, err := c.Fake.Invokes(action, &v1alpha1.RemoteOwnerList{})
	return err
}

// Patch applies the patch and returns the patched remoteOwner.
func (c *FakeRemoteOwners) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha1.RemoteOwner, err error) {
	emptyResult := &v1alpha1.RemoteOwner{}
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceActionWithOptions(remoteownersResource, c.ns, name, pt, data, opts, subresources...), emptyResult)

	if obj == nil {
		return emptyResult, err
	}
	return obj.(*v1alpha1.RemoteOwner), err
}
