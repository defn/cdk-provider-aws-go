package quicksightanalysis


type QuicksightAnalysisPermissions struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.33.0/docs/resources/quicksight_analysis#actions QuicksightAnalysis#actions}.
	Actions *[]*string `field:"required" json:"actions" yaml:"actions"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.33.0/docs/resources/quicksight_analysis#principal QuicksightAnalysis#principal}.
	Principal *string `field:"required" json:"principal" yaml:"principal"`
}

