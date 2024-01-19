package rdsclusterroleassociation


type RdsClusterRoleAssociationTimeouts struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.33.0/docs/resources/rds_cluster_role_association#create RdsClusterRoleAssociation#create}.
	Create *string `field:"optional" json:"create" yaml:"create"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.33.0/docs/resources/rds_cluster_role_association#delete RdsClusterRoleAssociation#delete}.
	Delete *string `field:"optional" json:"delete" yaml:"delete"`
}
