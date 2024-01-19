package lambdafunction


type LambdaFunctionEphemeralStorage struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.33.0/docs/resources/lambda_function#size LambdaFunction#size}.
	Size *float64 `field:"optional" json:"size" yaml:"size"`
}

