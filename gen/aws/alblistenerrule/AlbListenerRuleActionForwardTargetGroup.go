package alblistenerrule


type AlbListenerRuleActionForwardTargetGroup struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.33.0/docs/resources/alb_listener_rule#arn AlbListenerRule#arn}.
	Arn *string `field:"required" json:"arn" yaml:"arn"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.33.0/docs/resources/alb_listener_rule#weight AlbListenerRule#weight}.
	Weight *float64 `field:"optional" json:"weight" yaml:"weight"`
}

