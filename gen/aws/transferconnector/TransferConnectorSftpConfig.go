package transferconnector


type TransferConnectorSftpConfig struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.33.0/docs/resources/transfer_connector#trusted_host_keys TransferConnector#trusted_host_keys}.
	TrustedHostKeys *[]*string `field:"optional" json:"trustedHostKeys" yaml:"trustedHostKeys"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.33.0/docs/resources/transfer_connector#user_secret_id TransferConnector#user_secret_id}.
	UserSecretId *string `field:"optional" json:"userSecretId" yaml:"userSecretId"`
}
