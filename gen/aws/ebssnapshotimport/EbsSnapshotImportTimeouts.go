package ebssnapshotimport


type EbsSnapshotImportTimeouts struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.33.0/docs/resources/ebs_snapshot_import#create EbsSnapshotImport#create}.
	Create *string `field:"optional" json:"create" yaml:"create"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.33.0/docs/resources/ebs_snapshot_import#delete EbsSnapshotImport#delete}.
	Delete *string `field:"optional" json:"delete" yaml:"delete"`
}

