package sagemakerdomain


type SagemakerDomainDefaultUserSettingsCustomFileSystemConfigEfsFileSystemConfig struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.33.0/docs/resources/sagemaker_domain#file_system_id SagemakerDomain#file_system_id}.
	FileSystemId *string `field:"required" json:"fileSystemId" yaml:"fileSystemId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.33.0/docs/resources/sagemaker_domain#file_system_path SagemakerDomain#file_system_path}.
	FileSystemPath *string `field:"required" json:"fileSystemPath" yaml:"fileSystemPath"`
}
