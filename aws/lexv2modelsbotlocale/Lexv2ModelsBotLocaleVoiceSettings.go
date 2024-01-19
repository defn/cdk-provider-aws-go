package lexv2modelsbotlocale


type Lexv2ModelsBotLocaleVoiceSettings struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.33.0/docs/resources/lexv2models_bot_locale#voice_id Lexv2ModelsBotLocale#voice_id}.
	VoiceId *string `field:"required" json:"voiceId" yaml:"voiceId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.33.0/docs/resources/lexv2models_bot_locale#engine Lexv2ModelsBotLocale#engine}.
	Engine *string `field:"optional" json:"engine" yaml:"engine"`
}

