//go:build no_runtime_type_checking

package cdk8soperator

// Building without runtime type checking enabled, so all the below just return nil

func (o *jsiiProxy_Operator) validateAddProviderParameters(provider *CustomResourceProvider) error {
	return nil
}

func validateOperator_IsConstructParameters(x interface{}) error {
	return nil
}

func validateNewOperatorParameters(props *OperatorProps) error {
	return nil
}

