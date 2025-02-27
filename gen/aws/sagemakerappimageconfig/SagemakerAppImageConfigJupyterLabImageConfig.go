package sagemakerappimageconfig


type SagemakerAppImageConfigJupyterLabImageConfig struct {
	// container_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.33.0/docs/resources/sagemaker_app_image_config#container_config SagemakerAppImageConfig#container_config}
	ContainerConfig *SagemakerAppImageConfigJupyterLabImageConfigContainerConfig `field:"optional" json:"containerConfig" yaml:"containerConfig"`
}

