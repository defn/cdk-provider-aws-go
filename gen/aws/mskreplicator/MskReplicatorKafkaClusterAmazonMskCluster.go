package mskreplicator


type MskReplicatorKafkaClusterAmazonMskCluster struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.33.0/docs/resources/msk_replicator#msk_cluster_arn MskReplicator#msk_cluster_arn}.
	MskClusterArn *string `field:"required" json:"mskClusterArn" yaml:"mskClusterArn"`
}
