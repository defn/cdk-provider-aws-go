package customerprofilesdomain


type CustomerprofilesDomainRuleBasedMatchingExportingConfig struct {
	// s3_exporting block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.33.0/docs/resources/customerprofiles_domain#s3_exporting CustomerprofilesDomain#s3_exporting}
	S3Exporting *CustomerprofilesDomainRuleBasedMatchingExportingConfigS3Exporting `field:"optional" json:"s3Exporting" yaml:"s3Exporting"`
}

