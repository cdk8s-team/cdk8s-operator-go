package cdk8soperator

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/constructs-go/constructs/v10"
)

// The handler for this custom resource provider.
type ICustomResourceProviderHandler interface {
	Apply(scope constructs.Construct, id *string, spec interface{}) constructs.Construct
}

// The jsii proxy for ICustomResourceProviderHandler
type jsiiProxy_ICustomResourceProviderHandler struct {
	_ byte // padding
}

func (i *jsiiProxy_ICustomResourceProviderHandler) Apply(scope constructs.Construct, id *string, spec interface{}) constructs.Construct {
	if err := i.validateApplyParameters(scope, id, spec); err != nil {
		panic(err)
	}
	var returns constructs.Construct

	_jsii_.Invoke(
		i,
		"apply",
		[]interface{}{scope, id, spec},
		&returns,
	)

	return returns
}

