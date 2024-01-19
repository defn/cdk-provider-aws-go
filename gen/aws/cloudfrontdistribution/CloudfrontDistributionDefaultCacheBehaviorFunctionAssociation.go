package cloudfrontdistribution


type CloudfrontDistributionDefaultCacheBehaviorFunctionAssociation struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.33.0/docs/resources/cloudfront_distribution#event_type CloudfrontDistribution#event_type}.
	EventType *string `field:"required" json:"eventType" yaml:"eventType"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.33.0/docs/resources/cloudfront_distribution#function_arn CloudfrontDistribution#function_arn}.
	FunctionArn *string `field:"required" json:"functionArn" yaml:"functionArn"`
}
