package lambdaeventsourcemapping


type LambdaEventSourceMappingFilterCriteria struct {
	// filter block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.33.0/docs/resources/lambda_event_source_mapping#filter LambdaEventSourceMapping#filter}
	Filter interface{} `field:"optional" json:"filter" yaml:"filter"`
}

