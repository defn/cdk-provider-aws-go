package medialivemultiplex


type MedialiveMultiplexTimeouts struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.33.0/docs/resources/medialive_multiplex#create MedialiveMultiplex#create}.
	Create *string `field:"optional" json:"create" yaml:"create"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.33.0/docs/resources/medialive_multiplex#delete MedialiveMultiplex#delete}.
	Delete *string `field:"optional" json:"delete" yaml:"delete"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.33.0/docs/resources/medialive_multiplex#update MedialiveMultiplex#update}.
	Update *string `field:"optional" json:"update" yaml:"update"`
}
