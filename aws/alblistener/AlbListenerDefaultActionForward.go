package alblistener


type AlbListenerDefaultActionForward struct {
	// target_group block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.33.0/docs/resources/alb_listener#target_group AlbListener#target_group}
	TargetGroup interface{} `field:"required" json:"targetGroup" yaml:"targetGroup"`
	// stickiness block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.33.0/docs/resources/alb_listener#stickiness AlbListener#stickiness}
	Stickiness *AlbListenerDefaultActionForwardStickiness `field:"optional" json:"stickiness" yaml:"stickiness"`
}
