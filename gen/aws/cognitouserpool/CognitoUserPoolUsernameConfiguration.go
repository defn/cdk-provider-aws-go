package cognitouserpool


type CognitoUserPoolUsernameConfiguration struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.33.0/docs/resources/cognito_user_pool#case_sensitive CognitoUserPool#case_sensitive}.
	CaseSensitive interface{} `field:"required" json:"caseSensitive" yaml:"caseSensitive"`
}

