package quicksighttheme


type QuicksightThemePermissions struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.33.0/docs/resources/quicksight_theme#actions QuicksightTheme#actions}.
	Actions *[]*string `field:"required" json:"actions" yaml:"actions"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.33.0/docs/resources/quicksight_theme#principal QuicksightTheme#principal}.
	Principal *string `field:"required" json:"principal" yaml:"principal"`
}

