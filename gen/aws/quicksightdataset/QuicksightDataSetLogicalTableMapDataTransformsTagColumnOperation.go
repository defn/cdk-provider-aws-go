package quicksightdataset


type QuicksightDataSetLogicalTableMapDataTransformsTagColumnOperation struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.33.0/docs/resources/quicksight_data_set#column_name QuicksightDataSet#column_name}.
	ColumnName *string `field:"required" json:"columnName" yaml:"columnName"`
	// tags block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.33.0/docs/resources/quicksight_data_set#tags QuicksightDataSet#tags}
	Tags interface{} `field:"required" json:"tags" yaml:"tags"`
}

