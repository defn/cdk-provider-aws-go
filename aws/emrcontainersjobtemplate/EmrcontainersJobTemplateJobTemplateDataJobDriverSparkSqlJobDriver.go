package emrcontainersjobtemplate


type EmrcontainersJobTemplateJobTemplateDataJobDriverSparkSqlJobDriver struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.33.0/docs/resources/emrcontainers_job_template#entry_point EmrcontainersJobTemplate#entry_point}.
	EntryPoint *string `field:"optional" json:"entryPoint" yaml:"entryPoint"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.33.0/docs/resources/emrcontainers_job_template#spark_sql_parameters EmrcontainersJobTemplate#spark_sql_parameters}.
	SparkSqlParameters *string `field:"optional" json:"sparkSqlParameters" yaml:"sparkSqlParameters"`
}
