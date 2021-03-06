// Code generated by failery v1.0.0. DO NOT EDIT.

package disabled

import v1alpha1 "github.com/kyma-project/kyma/components/service-binding-usage-controller/pkg/apis/servicecatalog/v1alpha1"

// ServiceBindingUsageLister is an autogenerated failing mock type for the ServiceBindingUsageLister type
type ServiceBindingUsageLister struct {
	err error
}

// NewServiceBindingUsageLister creates a new ServiceBindingUsageLister type instance
func NewServiceBindingUsageLister(err error) *ServiceBindingUsageLister {
	return &ServiceBindingUsageLister{err: err}
}

// ListByUsageKind provides a failing mock function with given fields: namespace, kind, resourceName
func (_m *ServiceBindingUsageLister) ListByUsageKind(namespace string, kind string, resourceName string) ([]*v1alpha1.ServiceBindingUsage, error) {
	var r0 []*v1alpha1.ServiceBindingUsage
	var r1 error
	r1 = _m.err

	return r0, r1
}
