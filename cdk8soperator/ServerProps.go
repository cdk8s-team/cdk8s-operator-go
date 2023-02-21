// Create Kubernetes CRD Operators using CDK8s Constructs
package cdk8soperator


type ServerProps struct {
	// The command to execute in order to synthesize the CDK app.
	AppCommand *string `field:"required" json:"appCommand" yaml:"appCommand"`
}

