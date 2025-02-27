package instance


type InstanceMaintenanceOptions struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.33.0/docs/resources/instance#auto_recovery Instance#auto_recovery}.
	AutoRecovery *string `field:"optional" json:"autoRecovery" yaml:"autoRecovery"`
}

