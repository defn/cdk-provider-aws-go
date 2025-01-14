package efsfilesystem


type EfsFileSystemProtection struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.33.0/docs/resources/efs_file_system#replication_overwrite EfsFileSystem#replication_overwrite}.
	ReplicationOverwrite *string `field:"optional" json:"replicationOverwrite" yaml:"replicationOverwrite"`
}

