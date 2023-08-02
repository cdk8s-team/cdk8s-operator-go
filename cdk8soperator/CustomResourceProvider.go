package cdk8soperator


type CustomResourceProvider struct {
	// API version of the custom resource.
	// Default: "v1".
	//
	ApiVersion *string `field:"required" json:"apiVersion" yaml:"apiVersion"`
	// The construct handler.
	Handler ICustomResourceProviderHandler `field:"required" json:"handler" yaml:"handler"`
	// Kind of this custom resource.
	Kind *string `field:"required" json:"kind" yaml:"kind"`
}

