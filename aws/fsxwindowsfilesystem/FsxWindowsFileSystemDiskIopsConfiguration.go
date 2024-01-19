package fsxwindowsfilesystem


type FsxWindowsFileSystemDiskIopsConfiguration struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.33.0/docs/resources/fsx_windows_file_system#iops FsxWindowsFileSystem#iops}.
	Iops *float64 `field:"optional" json:"iops" yaml:"iops"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.33.0/docs/resources/fsx_windows_file_system#mode FsxWindowsFileSystem#mode}.
	Mode *string `field:"optional" json:"mode" yaml:"mode"`
}
