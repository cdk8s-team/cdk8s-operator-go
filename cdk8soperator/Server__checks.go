//go:build !no_runtime_type_checking

// Create Kubernetes CRD Operators using CDK8s Constructs
package cdk8soperator

import (
	"fmt"

	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

func validateNewServerParameters(props *ServerProps) error {
	if props == nil {
		return fmt.Errorf("parameter props is required, but nil was provided")
	}
	if err := _jsii_.ValidateStruct(props, func() string { return "parameter props" }); err != nil {
		return err
	}

	return nil
}

