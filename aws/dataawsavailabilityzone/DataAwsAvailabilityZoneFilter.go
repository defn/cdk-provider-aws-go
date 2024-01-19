package dataawsavailabilityzone


type DataAwsAvailabilityZoneFilter struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.33.0/docs/data-sources/availability_zone#name DataAwsAvailabilityZone#name}.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.33.0/docs/data-sources/availability_zone#values DataAwsAvailabilityZone#values}.
	Values *[]*string `field:"required" json:"values" yaml:"values"`
}
