package securitygroup


type SecurityGroupTimeouts struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.33.0/docs/resources/security_group#create SecurityGroup#create}.
	Create *string `field:"optional" json:"create" yaml:"create"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.33.0/docs/resources/security_group#delete SecurityGroup#delete}.
	Delete *string `field:"optional" json:"delete" yaml:"delete"`
}

