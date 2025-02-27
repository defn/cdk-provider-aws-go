package signersigningjob


type SignerSigningJobDestination struct {
	// s3 block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.33.0/docs/resources/signer_signing_job#s3 SignerSigningJob#s3}
	S3 *SignerSigningJobDestinationS3 `field:"required" json:"s3" yaml:"s3"`
}

