package neptuneclusterinstance


type NeptuneClusterInstanceTimeouts struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.33.0/docs/resources/neptune_cluster_instance#create NeptuneClusterInstance#create}.
	Create *string `field:"optional" json:"create" yaml:"create"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.33.0/docs/resources/neptune_cluster_instance#delete NeptuneClusterInstance#delete}.
	Delete *string `field:"optional" json:"delete" yaml:"delete"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.33.0/docs/resources/neptune_cluster_instance#update NeptuneClusterInstance#update}.
	Update *string `field:"optional" json:"update" yaml:"update"`
}
