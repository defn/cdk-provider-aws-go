package imagebuilderimage


type ImagebuilderImageImageScanningConfiguration struct {
	// ecr_configuration block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.33.0/docs/resources/imagebuilder_image#ecr_configuration ImagebuilderImage#ecr_configuration}
	EcrConfiguration *ImagebuilderImageImageScanningConfigurationEcrConfiguration `field:"optional" json:"ecrConfiguration" yaml:"ecrConfiguration"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.33.0/docs/resources/imagebuilder_image#image_scanning_enabled ImagebuilderImage#image_scanning_enabled}.
	ImageScanningEnabled interface{} `field:"optional" json:"imageScanningEnabled" yaml:"imageScanningEnabled"`
}
