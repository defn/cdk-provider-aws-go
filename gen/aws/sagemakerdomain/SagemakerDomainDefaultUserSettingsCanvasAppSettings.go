package sagemakerdomain


type SagemakerDomainDefaultUserSettingsCanvasAppSettings struct {
	// direct_deploy_settings block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.33.0/docs/resources/sagemaker_domain#direct_deploy_settings SagemakerDomain#direct_deploy_settings}
	DirectDeploySettings *SagemakerDomainDefaultUserSettingsCanvasAppSettingsDirectDeploySettings `field:"optional" json:"directDeploySettings" yaml:"directDeploySettings"`
	// identity_provider_oauth_settings block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.33.0/docs/resources/sagemaker_domain#identity_provider_oauth_settings SagemakerDomain#identity_provider_oauth_settings}
	IdentityProviderOauthSettings interface{} `field:"optional" json:"identityProviderOauthSettings" yaml:"identityProviderOauthSettings"`
	// kendra_settings block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.33.0/docs/resources/sagemaker_domain#kendra_settings SagemakerDomain#kendra_settings}
	KendraSettings *SagemakerDomainDefaultUserSettingsCanvasAppSettingsKendraSettings `field:"optional" json:"kendraSettings" yaml:"kendraSettings"`
	// model_register_settings block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.33.0/docs/resources/sagemaker_domain#model_register_settings SagemakerDomain#model_register_settings}
	ModelRegisterSettings *SagemakerDomainDefaultUserSettingsCanvasAppSettingsModelRegisterSettings `field:"optional" json:"modelRegisterSettings" yaml:"modelRegisterSettings"`
	// time_series_forecasting_settings block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.33.0/docs/resources/sagemaker_domain#time_series_forecasting_settings SagemakerDomain#time_series_forecasting_settings}
	TimeSeriesForecastingSettings *SagemakerDomainDefaultUserSettingsCanvasAppSettingsTimeSeriesForecastingSettings `field:"optional" json:"timeSeriesForecastingSettings" yaml:"timeSeriesForecastingSettings"`
	// workspace_settings block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.33.0/docs/resources/sagemaker_domain#workspace_settings SagemakerDomain#workspace_settings}
	WorkspaceSettings *SagemakerDomainDefaultUserSettingsCanvasAppSettingsWorkspaceSettings `field:"optional" json:"workspaceSettings" yaml:"workspaceSettings"`
}
