package dynamodbtable


type DynamodbTableImportTableInputFormatOptions struct {
	// csv block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.33.0/docs/resources/dynamodb_table#csv DynamodbTable#csv}
	Csv *DynamodbTableImportTableInputFormatOptionsCsv `field:"optional" json:"csv" yaml:"csv"`
}

