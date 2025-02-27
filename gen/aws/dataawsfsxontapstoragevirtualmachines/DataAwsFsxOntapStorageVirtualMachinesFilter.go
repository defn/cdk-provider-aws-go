package dataawsfsxontapstoragevirtualmachines


type DataAwsFsxOntapStorageVirtualMachinesFilter struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.33.0/docs/data-sources/fsx_ontap_storage_virtual_machines#name DataAwsFsxOntapStorageVirtualMachines#name}.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.33.0/docs/data-sources/fsx_ontap_storage_virtual_machines#values DataAwsFsxOntapStorageVirtualMachines#values}.
	Values *[]*string `field:"required" json:"values" yaml:"values"`
}

