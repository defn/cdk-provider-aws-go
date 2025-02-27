package emrcluster


type EmrClusterStep struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.33.0/docs/resources/emr_cluster#action_on_failure EmrCluster#action_on_failure}.
	ActionOnFailure *string `field:"optional" json:"actionOnFailure" yaml:"actionOnFailure"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.33.0/docs/resources/emr_cluster#hadoop_jar_step EmrCluster#hadoop_jar_step}.
	HadoopJarStep interface{} `field:"optional" json:"hadoopJarStep" yaml:"hadoopJarStep"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.33.0/docs/resources/emr_cluster#name EmrCluster#name}.
	Name *string `field:"optional" json:"name" yaml:"name"`
}

