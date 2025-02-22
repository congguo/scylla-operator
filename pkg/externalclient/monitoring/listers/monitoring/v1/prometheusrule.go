// Code generated by lister-gen. DO NOT EDIT.

package v1

import (
	v1 "github.com/scylladb/scylla-operator/pkg/externalapi/monitoring/v1"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/client-go/listers"
	"k8s.io/client-go/tools/cache"
)

// PrometheusRuleLister helps list PrometheusRules.
// All objects returned here must be treated as read-only.
type PrometheusRuleLister interface {
	// List lists all PrometheusRules in the indexer.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1.PrometheusRule, err error)
	// PrometheusRules returns an object that can list and get PrometheusRules.
	PrometheusRules(namespace string) PrometheusRuleNamespaceLister
	PrometheusRuleListerExpansion
}

// prometheusRuleLister implements the PrometheusRuleLister interface.
type prometheusRuleLister struct {
	listers.ResourceIndexer[*v1.PrometheusRule]
}

// NewPrometheusRuleLister returns a new PrometheusRuleLister.
func NewPrometheusRuleLister(indexer cache.Indexer) PrometheusRuleLister {
	return &prometheusRuleLister{listers.New[*v1.PrometheusRule](indexer, v1.Resource("prometheusrule"))}
}

// PrometheusRules returns an object that can list and get PrometheusRules.
func (s *prometheusRuleLister) PrometheusRules(namespace string) PrometheusRuleNamespaceLister {
	return prometheusRuleNamespaceLister{listers.NewNamespaced[*v1.PrometheusRule](s.ResourceIndexer, namespace)}
}

// PrometheusRuleNamespaceLister helps list and get PrometheusRules.
// All objects returned here must be treated as read-only.
type PrometheusRuleNamespaceLister interface {
	// List lists all PrometheusRules in the indexer for a given namespace.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1.PrometheusRule, err error)
	// Get retrieves the PrometheusRule from the indexer for a given namespace and name.
	// Objects returned here must be treated as read-only.
	Get(name string) (*v1.PrometheusRule, error)
	PrometheusRuleNamespaceListerExpansion
}

// prometheusRuleNamespaceLister implements the PrometheusRuleNamespaceLister
// interface.
type prometheusRuleNamespaceLister struct {
	listers.ResourceIndexer[*v1.PrometheusRule]
}
