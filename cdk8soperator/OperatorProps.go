// Create Kubernetes CRD Operators using CDK8s Constructs
package cdk8soperator


type OperatorProps struct {
	// A Kubernetes JSON manifest with a single resource that is matched against one of the providers within this operator.
	InputFile *string `field:"optional" json:"inputFile" yaml:"inputFile"`
	// Where to write the synthesized output.
	OutputFile *string `field:"optional" json:"outputFile" yaml:"outputFile"`
}

