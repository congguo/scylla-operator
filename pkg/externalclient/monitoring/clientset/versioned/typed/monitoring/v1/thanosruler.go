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

// ThanosRulersGetter has a method to return a ThanosRulerInterface.
// A group's client should implement this interface.
type ThanosRulersGetter interface {
	ThanosRulers(namespace string) ThanosRulerInterface
}

// ThanosRulerInterface has methods to work with ThanosRuler resources.
type ThanosRulerInterface interface {
	Create(ctx context.Context, thanosRuler *v1.ThanosRuler, opts metav1.CreateOptions) (*v1.ThanosRuler, error)
	Update(ctx context.Context, thanosRuler *v1.ThanosRuler, opts metav1.UpdateOptions) (*v1.ThanosRuler, error)
	// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
	UpdateStatus(ctx context.Context, thanosRuler *v1.ThanosRuler, opts metav1.UpdateOptions) (*v1.ThanosRuler, error)
	Delete(ctx context.Context, name string, opts metav1.DeleteOptions) error
	DeleteCollection(ctx context.Context, opts metav1.DeleteOptions, listOpts metav1.ListOptions) error
	Get(ctx context.Context, name string, opts metav1.GetOptions) (*v1.ThanosRuler, error)
	List(ctx context.Context, opts metav1.ListOptions) (*v1.ThanosRulerList, error)
	Watch(ctx context.Context, opts metav1.ListOptions) (watch.Interface, error)
	Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts metav1.PatchOptions, subresources ...string) (result *v1.ThanosRuler, err error)
	ThanosRulerExpansion
}

// thanosRulers implements ThanosRulerInterface
type thanosRulers struct {
	*gentype.ClientWithList[*v1.ThanosRuler, *v1.ThanosRulerList]
}

// newThanosRulers returns a ThanosRulers
func newThanosRulers(c *MonitoringV1Client, namespace string) *thanosRulers {
	return &thanosRulers{
		gentype.NewClientWithList[*v1.ThanosRuler, *v1.ThanosRulerList](
			"thanosrulers",
			c.RESTClient(),
			scheme.ParameterCodec,
			namespace,
			func() *v1.ThanosRuler { return &v1.ThanosRuler{} },
			func() *v1.ThanosRulerList { return &v1.ThanosRulerList{} }),
	}
}
