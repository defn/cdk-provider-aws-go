package appmeshvirtualrouter


type AppmeshVirtualRouterSpec struct {
	// listener block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.33.0/docs/resources/appmesh_virtual_router#listener AppmeshVirtualRouter#listener}
	Listener interface{} `field:"optional" json:"listener" yaml:"listener"`
}

