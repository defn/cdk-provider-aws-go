package launchtemplate


type LaunchTemplateInstanceRequirementsMemoryMib struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.33.0/docs/resources/launch_template#min LaunchTemplate#min}.
	Min *float64 `field:"required" json:"min" yaml:"min"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.33.0/docs/resources/launch_template#max LaunchTemplate#max}.
	Max *float64 `field:"optional" json:"max" yaml:"max"`
}
