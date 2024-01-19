package securityhubinsight


type SecurityhubInsightFiltersUserDefinedValues struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.33.0/docs/resources/securityhub_insight#comparison SecurityhubInsight#comparison}.
	Comparison *string `field:"required" json:"comparison" yaml:"comparison"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.33.0/docs/resources/securityhub_insight#key SecurityhubInsight#key}.
	Key *string `field:"required" json:"key" yaml:"key"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.33.0/docs/resources/securityhub_insight#value SecurityhubInsight#value}.
	Value *string `field:"required" json:"value" yaml:"value"`
}
