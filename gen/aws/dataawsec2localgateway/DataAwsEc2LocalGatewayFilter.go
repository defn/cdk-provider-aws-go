package dataawsec2localgateway


type DataAwsEc2LocalGatewayFilter struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.33.0/docs/data-sources/ec2_local_gateway#name DataAwsEc2LocalGateway#name}.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.33.0/docs/data-sources/ec2_local_gateway#values DataAwsEc2LocalGateway#values}.
	Values *[]*string `field:"required" json:"values" yaml:"values"`
}

