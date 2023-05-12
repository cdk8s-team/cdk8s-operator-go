package cdk8soperator

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdk8s-team/cdk8s-operator-go/cdk8soperator/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/cdk8s-team/cdk8s-core-go/cdk8s/v2"
	"github.com/cdk8s-team/cdk8s-operator-go/cdk8soperator/internal"
)

// A CDK8s app which allows implementing Kubernetes operators using CDK8s constructs.
type Operator interface {
	cdk8s.App
	// Returns all the charts in this app, sorted topologically.
	Charts() *[]cdk8s.Chart
	// The tree node.
	Node() constructs.Node
	// The output directory into which manifests will be synthesized.
	Outdir() *string
	// The file extension to use for rendered YAML files.
	OutputFileExtension() *string
	// How to divide the YAML output into files.
	YamlOutputType() cdk8s.YamlOutputType
	// Adds a custom resource provider to this operator.
	AddProvider(provider *CustomResourceProvider)
	// Reads a Kubernetes manifest in JSON format from STDIN or the file specified as the first positional command-line argument.
	//
	// This manifest is expected to
	// include a single Kubernetes resource. Then, we match `apiVersion` and
	// `kind` to one of the registered providers and if we do, we invoke
	// `apply()`, passing it the `spec` of the input manifest and a chart as a
	// scope. The chart is then synthesized and the output manifest is written to
	// STDOUT.
	Synth()
	// Synthesizes the app into a YAML string.
	//
	// Returns: A string with all YAML objects across all charts in this app.
	SynthYaml() *string
	// Returns a string representation of this construct.
	ToString() *string
}

// The jsii proxy struct for Operator
type jsiiProxy_Operator struct {
	internal.Type__cdk8sApp
}

func (j *jsiiProxy_Operator) Charts() *[]cdk8s.Chart {
	var returns *[]cdk8s.Chart
	_jsii_.Get(
		j,
		"charts",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Operator) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Operator) Outdir() *string {
	var returns *string
	_jsii_.Get(
		j,
		"outdir",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Operator) OutputFileExtension() *string {
	var returns *string
	_jsii_.Get(
		j,
		"outputFileExtension",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Operator) YamlOutputType() cdk8s.YamlOutputType {
	var returns cdk8s.YamlOutputType
	_jsii_.Get(
		j,
		"yamlOutputType",
		&returns,
	)
	return returns
}


func NewOperator(props *OperatorProps) Operator {
	_init_.Initialize()

	if err := validateNewOperatorParameters(props); err != nil {
		panic(err)
	}
	j := jsiiProxy_Operator{}

	_jsii_.Create(
		"cdk8s-operator.Operator",
		[]interface{}{props},
		&j,
	)

	return &j
}

func NewOperator_Override(o Operator, props *OperatorProps) {
	_init_.Initialize()

	_jsii_.Create(
		"cdk8s-operator.Operator",
		[]interface{}{props},
		o,
	)
}

// Checks if `x` is a construct.
//
// Use this method instead of `instanceof` to properly detect `Construct`
// instances, even when the construct library is symlinked.
//
// Explanation: in JavaScript, multiple copies of the `constructs` library on
// disk are seen as independent, completely different libraries. As a
// consequence, the class `Construct` in each copy of the `constructs` library
// is seen as a different class, and an instance of one class will not test as
// `instanceof` the other class. `npm install` will not create installations
// like this, but users may manually symlink construct libraries together or
// use a monorepo tool: in those cases, multiple copies of the `constructs`
// library can be accidentally installed, and `instanceof` will behave
// unpredictably. It is safest to avoid using `instanceof`, and using
// this type-testing method instead.
//
// Returns: true if `x` is an object created from a class which extends `Construct`.
func Operator_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateOperator_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"cdk8s-operator.Operator",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func (o *jsiiProxy_Operator) AddProvider(provider *CustomResourceProvider) {
	if err := o.validateAddProviderParameters(provider); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		o,
		"addProvider",
		[]interface{}{provider},
	)
}

func (o *jsiiProxy_Operator) Synth() {
	_jsii_.InvokeVoid(
		o,
		"synth",
		nil, // no parameters
	)
}

func (o *jsiiProxy_Operator) SynthYaml() *string {
	var returns *string

	_jsii_.Invoke(
		o,
		"synthYaml",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (o *jsiiProxy_Operator) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		o,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

