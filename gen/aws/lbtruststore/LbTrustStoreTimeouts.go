package lbtruststore


type LbTrustStoreTimeouts struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.33.0/docs/resources/lb_trust_store#create LbTrustStore#create}.
	Create *string `field:"optional" json:"create" yaml:"create"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.33.0/docs/resources/lb_trust_store#delete LbTrustStore#delete}.
	Delete *string `field:"optional" json:"delete" yaml:"delete"`
}
