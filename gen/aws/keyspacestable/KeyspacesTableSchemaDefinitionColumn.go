package keyspacestable


type KeyspacesTableSchemaDefinitionColumn struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.33.0/docs/resources/keyspaces_table#name KeyspacesTable#name}.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.33.0/docs/resources/keyspaces_table#type KeyspacesTable#type}.
	Type *string `field:"required" json:"type" yaml:"type"`
}

