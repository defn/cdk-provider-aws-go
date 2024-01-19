package comprehendentityrecognizer


type ComprehendEntityRecognizerVpcConfig struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.33.0/docs/resources/comprehend_entity_recognizer#security_group_ids ComprehendEntityRecognizer#security_group_ids}.
	SecurityGroupIds *[]*string `field:"required" json:"securityGroupIds" yaml:"securityGroupIds"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.33.0/docs/resources/comprehend_entity_recognizer#subnets ComprehendEntityRecognizer#subnets}.
	Subnets *[]*string `field:"required" json:"subnets" yaml:"subnets"`
}
