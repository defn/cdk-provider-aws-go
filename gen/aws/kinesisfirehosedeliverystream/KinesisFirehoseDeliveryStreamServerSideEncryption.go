package kinesisfirehosedeliverystream


type KinesisFirehoseDeliveryStreamServerSideEncryption struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.33.0/docs/resources/kinesis_firehose_delivery_stream#enabled KinesisFirehoseDeliveryStream#enabled}.
	Enabled interface{} `field:"optional" json:"enabled" yaml:"enabled"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.33.0/docs/resources/kinesis_firehose_delivery_stream#key_arn KinesisFirehoseDeliveryStream#key_arn}.
	KeyArn *string `field:"optional" json:"keyArn" yaml:"keyArn"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.33.0/docs/resources/kinesis_firehose_delivery_stream#key_type KinesisFirehoseDeliveryStream#key_type}.
	KeyType *string `field:"optional" json:"keyType" yaml:"keyType"`
}

