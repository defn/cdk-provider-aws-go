package apigatewayusageplan


type ApiGatewayUsagePlanQuotaSettings struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.33.0/docs/resources/api_gateway_usage_plan#limit ApiGatewayUsagePlan#limit}.
	Limit *float64 `field:"required" json:"limit" yaml:"limit"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.33.0/docs/resources/api_gateway_usage_plan#period ApiGatewayUsagePlan#period}.
	Period *string `field:"required" json:"period" yaml:"period"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.33.0/docs/resources/api_gateway_usage_plan#offset ApiGatewayUsagePlan#offset}.
	Offset *float64 `field:"optional" json:"offset" yaml:"offset"`
}

