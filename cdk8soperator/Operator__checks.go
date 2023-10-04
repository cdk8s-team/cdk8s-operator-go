//go:build !no_runtime_type_checking

package cdk8soperator

import (
	"fmt"

	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/constructs-go/constructs/v10"
)

func (o *jsiiProxy_Operator) validateAddProviderParameters(provider *CustomResourceProvider) error {
	if provider == nil {
		return fmt.Errorf("parameter provider is required, but nil was provided")
	}
	if err := _jsii_.ValidateStruct(provider, func() string { return "parameter provider" }); err != nil {
		return err
	}

	return nil
}

func validateOperator_IsConstructParameters(x interface{}) error {
	if x == nil {
		return fmt.Errorf("parameter x is required, but nil was provided")
	}

	return nil
}

func validateOperator_OfParameters(c constructs.IConstruct) error {
	if c == nil {
		return fmt.Errorf("parameter c is required, but nil was provided")
	}

	return nil
}

func validateNewOperatorParameters(props *OperatorProps) error {
	if err := _jsii_.ValidateStruct(props, func() string { return "parameter props" }); err != nil {
		return err
	}

	return nil
}

