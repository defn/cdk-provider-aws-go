package lambdafunction


type LambdaFunctionTracingConfig struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.33.0/docs/resources/lambda_function#mode LambdaFunction#mode}.
	Mode *string `field:"required" json:"mode" yaml:"mode"`
}

