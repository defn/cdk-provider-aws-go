package prometheusworkspace


type PrometheusWorkspaceLoggingConfiguration struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.33.0/docs/resources/prometheus_workspace#log_group_arn PrometheusWorkspace#log_group_arn}.
	LogGroupArn *string `field:"required" json:"logGroupArn" yaml:"logGroupArn"`
}

