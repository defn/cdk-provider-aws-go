package ecrrepository


type EcrRepositoryTimeouts struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.33.0/docs/resources/ecr_repository#delete EcrRepository#delete}.
	Delete *string `field:"optional" json:"delete" yaml:"delete"`
}

