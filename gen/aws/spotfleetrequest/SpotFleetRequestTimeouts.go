package spotfleetrequest


type SpotFleetRequestTimeouts struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.33.0/docs/resources/spot_fleet_request#create SpotFleetRequest#create}.
	Create *string `field:"optional" json:"create" yaml:"create"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.33.0/docs/resources/spot_fleet_request#delete SpotFleetRequest#delete}.
	Delete *string `field:"optional" json:"delete" yaml:"delete"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.33.0/docs/resources/spot_fleet_request#update SpotFleetRequest#update}.
	Update *string `field:"optional" json:"update" yaml:"update"`
}
