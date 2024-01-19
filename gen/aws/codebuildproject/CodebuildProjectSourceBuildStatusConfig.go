package codebuildproject


type CodebuildProjectSourceBuildStatusConfig struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.33.0/docs/resources/codebuild_project#context CodebuildProject#context}.
	Context *string `field:"optional" json:"context" yaml:"context"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.33.0/docs/resources/codebuild_project#target_url CodebuildProject#target_url}.
	TargetUrl *string `field:"optional" json:"targetUrl" yaml:"targetUrl"`
}
