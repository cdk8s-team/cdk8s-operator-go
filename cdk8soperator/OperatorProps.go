package cdk8soperator


type OperatorProps struct {
	// A Kubernetes JSON manifest with a single resource that is matched against one of the providers within this operator.
	// Default: - first position command-line argument or "/dev/stdin".
	//
	InputFile *string `field:"optional" json:"inputFile" yaml:"inputFile"`
	// Where to write the synthesized output.
	// Default: "/dev/stdout".
	//
	OutputFile *string `field:"optional" json:"outputFile" yaml:"outputFile"`
}

