package spotinstancerequest


type SpotInstanceRequestTimeouts struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.33.0/docs/resources/spot_instance_request#create SpotInstanceRequest#create}.
	Create *string `field:"optional" json:"create" yaml:"create"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.33.0/docs/resources/spot_instance_request#delete SpotInstanceRequest#delete}.
	Delete *string `field:"optional" json:"delete" yaml:"delete"`
}

