package ssmincidentsresponseplan


type SsmincidentsResponsePlanActionSsmAutomationParameter struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.33.0/docs/resources/ssmincidents_response_plan#name SsmincidentsResponsePlan#name}.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.33.0/docs/resources/ssmincidents_response_plan#values SsmincidentsResponsePlan#values}.
	Values *[]*string `field:"required" json:"values" yaml:"values"`
}
