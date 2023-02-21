// Create Kubernetes CRD Operators using CDK8s Constructs
package cdk8soperator

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdk8s-team/cdk8s-operator-go/cdk8soperator/jsii"
)

type Server interface {
	// Stop server.
	Close()
	// Starts HTTP server.
	Listen(port *float64) *float64
}

// The jsii proxy struct for Server
type jsiiProxy_Server struct {
	_ byte // padding
}

func NewServer(props *ServerProps) Server {
	_init_.Initialize()

	if err := validateNewServerParameters(props); err != nil {
		panic(err)
	}
	j := jsiiProxy_Server{}

	_jsii_.Create(
		"cdk8s-operator.Server",
		[]interface{}{props},
		&j,
	)

	return &j
}

func NewServer_Override(s Server, props *ServerProps) {
	_init_.Initialize()

	_jsii_.Create(
		"cdk8s-operator.Server",
		[]interface{}{props},
		s,
	)
}

func (s *jsiiProxy_Server) Close() {
	_jsii_.InvokeVoid(
		s,
		"close",
		nil, // no parameters
	)
}

func (s *jsiiProxy_Server) Listen(port *float64) *float64 {
	var returns *float64

	_jsii_.Invoke(
		s,
		"listen",
		[]interface{}{port},
		&returns,
	)

	return returns
}

