package imagebuilderimage


type ImagebuilderImageImageScanningConfigurationEcrConfiguration struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.33.0/docs/resources/imagebuilder_image#container_tags ImagebuilderImage#container_tags}.
	ContainerTags *[]*string `field:"optional" json:"containerTags" yaml:"containerTags"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.33.0/docs/resources/imagebuilder_image#repository_name ImagebuilderImage#repository_name}.
	RepositoryName *string `field:"optional" json:"repositoryName" yaml:"repositoryName"`
}

