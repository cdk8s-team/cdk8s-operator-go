package cdk8soperator

import (
	"reflect"

	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

func init() {
	_jsii_.RegisterStruct(
		"cdk8s-operator.CustomResourceProvider",
		reflect.TypeOf((*CustomResourceProvider)(nil)).Elem(),
	)
	_jsii_.RegisterInterface(
		"cdk8s-operator.ICustomResourceProviderHandler",
		reflect.TypeOf((*ICustomResourceProviderHandler)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "apply", GoMethod: "Apply"},
		},
		func() interface{} {
			return &jsiiProxy_ICustomResourceProviderHandler{}
		},
	)
	_jsii_.RegisterClass(
		"cdk8s-operator.Operator",
		reflect.TypeOf((*Operator)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "addProvider", GoMethod: "AddProvider"},
			_jsii_.MemberProperty{JsiiProperty: "charts", GoGetter: "Charts"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberProperty{JsiiProperty: "outdir", GoGetter: "Outdir"},
			_jsii_.MemberProperty{JsiiProperty: "outputFileExtension", GoGetter: "OutputFileExtension"},
			_jsii_.MemberMethod{JsiiMethod: "synth", GoMethod: "Synth"},
			_jsii_.MemberMethod{JsiiMethod: "synthYaml", GoMethod: "SynthYaml"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
			_jsii_.MemberProperty{JsiiProperty: "yamlOutputType", GoGetter: "YamlOutputType"},
		},
		func() interface{} {
			j := jsiiProxy_Operator{}
			_jsii_.InitJsiiProxy(&j.Type__cdk8sApp)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"cdk8s-operator.OperatorProps",
		reflect.TypeOf((*OperatorProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"cdk8s-operator.Server",
		reflect.TypeOf((*Server)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "close", GoMethod: "Close"},
			_jsii_.MemberMethod{JsiiMethod: "listen", GoMethod: "Listen"},
		},
		func() interface{} {
			return &jsiiProxy_Server{}
		},
	)
	_jsii_.RegisterStruct(
		"cdk8s-operator.ServerProps",
		reflect.TypeOf((*ServerProps)(nil)).Elem(),
	)
}
