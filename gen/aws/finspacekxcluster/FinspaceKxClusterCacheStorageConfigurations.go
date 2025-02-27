package finspacekxcluster


type FinspaceKxClusterCacheStorageConfigurations struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.33.0/docs/resources/finspace_kx_cluster#size FinspaceKxCluster#size}.
	Size *float64 `field:"required" json:"size" yaml:"size"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.33.0/docs/resources/finspace_kx_cluster#type FinspaceKxCluster#type}.
	Type *string `field:"required" json:"type" yaml:"type"`
}

