package mwaaenvironment


type MwaaEnvironmentLoggingConfigurationWebserverLogs struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.33.0/docs/resources/mwaa_environment#enabled MwaaEnvironment#enabled}.
	Enabled interface{} `field:"optional" json:"enabled" yaml:"enabled"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.33.0/docs/resources/mwaa_environment#log_level MwaaEnvironment#log_level}.
	LogLevel *string `field:"optional" json:"logLevel" yaml:"logLevel"`
}

