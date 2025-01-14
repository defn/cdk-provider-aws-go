package securitylakedatalake


type SecuritylakeDataLakeConfigurationReplicationConfiguration struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.33.0/docs/resources/securitylake_data_lake#regions SecuritylakeDataLake#regions}.
	Regions *[]*string `field:"optional" json:"regions" yaml:"regions"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.33.0/docs/resources/securitylake_data_lake#role_arn SecuritylakeDataLake#role_arn}.
	RoleArn *string `field:"optional" json:"roleArn" yaml:"roleArn"`
}

