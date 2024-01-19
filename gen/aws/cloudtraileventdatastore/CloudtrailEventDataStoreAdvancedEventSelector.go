package cloudtraileventdatastore


type CloudtrailEventDataStoreAdvancedEventSelector struct {
	// field_selector block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.33.0/docs/resources/cloudtrail_event_data_store#field_selector CloudtrailEventDataStore#field_selector}
	FieldSelector interface{} `field:"optional" json:"fieldSelector" yaml:"fieldSelector"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.33.0/docs/resources/cloudtrail_event_data_store#name CloudtrailEventDataStore#name}.
	Name *string `field:"optional" json:"name" yaml:"name"`
}
