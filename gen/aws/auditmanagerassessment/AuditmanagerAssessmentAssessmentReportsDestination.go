package auditmanagerassessment


type AuditmanagerAssessmentAssessmentReportsDestination struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.33.0/docs/resources/auditmanager_assessment#destination AuditmanagerAssessment#destination}.
	Destination *string `field:"required" json:"destination" yaml:"destination"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.33.0/docs/resources/auditmanager_assessment#destination_type AuditmanagerAssessment#destination_type}.
	DestinationType *string `field:"required" json:"destinationType" yaml:"destinationType"`
}
