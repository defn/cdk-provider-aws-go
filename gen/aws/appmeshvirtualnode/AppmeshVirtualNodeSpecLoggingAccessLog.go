package appmeshvirtualnode


type AppmeshVirtualNodeSpecLoggingAccessLog struct {
	// file block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.33.0/docs/resources/appmesh_virtual_node#file AppmeshVirtualNode#file}
	File *AppmeshVirtualNodeSpecLoggingAccessLogFile `field:"optional" json:"file" yaml:"file"`
}

