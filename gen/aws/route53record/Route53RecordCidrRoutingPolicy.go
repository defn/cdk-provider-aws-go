package route53record


type Route53RecordCidrRoutingPolicy struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.33.0/docs/resources/route53_record#collection_id Route53Record#collection_id}.
	CollectionId *string `field:"required" json:"collectionId" yaml:"collectionId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.33.0/docs/resources/route53_record#location_name Route53Record#location_name}.
	LocationName *string `field:"required" json:"locationName" yaml:"locationName"`
}
