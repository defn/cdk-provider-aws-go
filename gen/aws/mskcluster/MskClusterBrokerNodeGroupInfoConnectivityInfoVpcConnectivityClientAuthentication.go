package mskcluster


type MskClusterBrokerNodeGroupInfoConnectivityInfoVpcConnectivityClientAuthentication struct {
	// sasl block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.33.0/docs/resources/msk_cluster#sasl MskCluster#sasl}
	Sasl *MskClusterBrokerNodeGroupInfoConnectivityInfoVpcConnectivityClientAuthenticationSasl `field:"optional" json:"sasl" yaml:"sasl"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.33.0/docs/resources/msk_cluster#tls MskCluster#tls}.
	Tls interface{} `field:"optional" json:"tls" yaml:"tls"`
}

