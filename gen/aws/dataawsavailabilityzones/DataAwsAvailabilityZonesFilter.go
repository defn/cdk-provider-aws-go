package dataawsavailabilityzones


type DataAwsAvailabilityZonesFilter struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.33.0/docs/data-sources/availability_zones#name DataAwsAvailabilityZones#name}.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.33.0/docs/data-sources/availability_zones#values DataAwsAvailabilityZones#values}.
	Values *[]*string `field:"required" json:"values" yaml:"values"`
}

