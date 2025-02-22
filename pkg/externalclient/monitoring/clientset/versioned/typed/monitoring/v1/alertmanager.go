// Code generated by client-gen. DO NOT EDIT.

package v1

import (
	"context"

	v1 "github.com/scylladb/scylla-operator/pkg/externalapi/monitoring/v1"
	scheme "github.com/scylladb/scylla-operator/pkg/externalclient/monitoring/clientset/versioned/scheme"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	gentype "k8s.io/client-go/gentype"
)

// AlertmanagersGetter has a method to return a AlertmanagerInterface.
// A group's client should implement this interface.
type AlertmanagersGetter interface {
	Alertmanagers(namespace string) AlertmanagerInterface
}

// AlertmanagerInterface has methods to work with Alertmanager resources.
type AlertmanagerInterface interface {
	Create(ctx context.Context, alertmanager *v1.Alertmanager, opts metav1.CreateOptions) (*v1.Alertmanager, error)
	Update(ctx context.Context, alertmanager *v1.Alertmanager, opts metav1.UpdateOptions) (*v1.Alertmanager, error)
	// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
	UpdateStatus(ctx context.Context, alertmanager *v1.Alertmanager, opts metav1.UpdateOptions) (*v1.Alertmanager, error)
	Delete(ctx context.Context, name string, opts metav1.DeleteOptions) error
	DeleteCollection(ctx context.Context, opts metav1.DeleteOptions, listOpts metav1.ListOptions) error
	Get(ctx context.Context, name string, opts metav1.GetOptions) (*v1.Alertmanager, error)
	List(ctx context.Context, opts metav1.ListOptions) (*v1.AlertmanagerList, error)
	Watch(ctx context.Context, opts metav1.ListOptions) (watch.Interface, error)
	Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts metav1.PatchOptions, subresources ...string) (result *v1.Alertmanager, err error)
	AlertmanagerExpansion
}

// alertmanagers implements AlertmanagerInterface
type alertmanagers struct {
	*gentype.ClientWithList[*v1.Alertmanager, *v1.AlertmanagerList]
}

// newAlertmanagers returns a Alertmanagers
func newAlertmanagers(c *MonitoringV1Client, namespace string) *alertmanagers {
	return &alertmanagers{
		gentype.NewClientWithList[*v1.Alertmanager, *v1.AlertmanagerList](
			"alertmanagers",
			c.RESTClient(),
			scheme.ParameterCodec,
			namespace,
			func() *v1.Alertmanager { return &v1.Alertmanager{} },
			func() *v1.AlertmanagerList { return &v1.AlertmanagerList{} }),
	}
}
