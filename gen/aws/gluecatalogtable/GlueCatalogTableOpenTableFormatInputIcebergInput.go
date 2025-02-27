package gluecatalogtable


type GlueCatalogTableOpenTableFormatInputIcebergInput struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.33.0/docs/resources/glue_catalog_table#metadata_operation GlueCatalogTable#metadata_operation}.
	MetadataOperation *string `field:"required" json:"metadataOperation" yaml:"metadataOperation"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.33.0/docs/resources/glue_catalog_table#version GlueCatalogTable#version}.
	Version *string `field:"optional" json:"version" yaml:"version"`
}

