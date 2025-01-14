package finspacekxdataview


type FinspaceKxDataviewSegmentConfigurations struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.33.0/docs/resources/finspace_kx_dataview#db_paths FinspaceKxDataview#db_paths}.
	DbPaths *[]*string `field:"required" json:"dbPaths" yaml:"dbPaths"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.33.0/docs/resources/finspace_kx_dataview#volume_name FinspaceKxDataview#volume_name}.
	VolumeName *string `field:"required" json:"volumeName" yaml:"volumeName"`
}

