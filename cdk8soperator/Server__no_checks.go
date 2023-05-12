//go:build no_runtime_type_checking

package cdk8soperator

// Building without runtime type checking enabled, so all the below just return nil

func validateNewServerParameters(props *ServerProps) error {
	return nil
}

