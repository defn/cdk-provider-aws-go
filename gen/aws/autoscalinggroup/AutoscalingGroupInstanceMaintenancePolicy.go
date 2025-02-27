package autoscalinggroup


type AutoscalingGroupInstanceMaintenancePolicy struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.33.0/docs/resources/autoscaling_group#max_healthy_percentage AutoscalingGroup#max_healthy_percentage}.
	MaxHealthyPercentage *float64 `field:"required" json:"maxHealthyPercentage" yaml:"maxHealthyPercentage"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.33.0/docs/resources/autoscaling_group#min_healthy_percentage AutoscalingGroup#min_healthy_percentage}.
	MinHealthyPercentage *float64 `field:"required" json:"minHealthyPercentage" yaml:"minHealthyPercentage"`
}

