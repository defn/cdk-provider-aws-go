package mskconnectcustomplugin


type MskconnectCustomPluginTimeouts struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.33.0/docs/resources/mskconnect_custom_plugin#create MskconnectCustomPlugin#create}.
	Create *string `field:"optional" json:"create" yaml:"create"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.33.0/docs/resources/mskconnect_custom_plugin#delete MskconnectCustomPlugin#delete}.
	Delete *string `field:"optional" json:"delete" yaml:"delete"`
}
