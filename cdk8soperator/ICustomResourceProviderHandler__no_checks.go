//go:build no_runtime_type_checking

// Create Kubernetes CRD Operators using CDK8s Constructs
package cdk8soperator

// Building without runtime type checking enabled, so all the below just return nil

func (i *jsiiProxy_ICustomResourceProviderHandler) validateApplyParameters(scope constructs.Construct, id *string, spec interface{}) error {
	return nil
}

