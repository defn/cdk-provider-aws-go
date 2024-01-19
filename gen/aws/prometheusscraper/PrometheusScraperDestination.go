package prometheusscraper


type PrometheusScraperDestination struct {
	// amp block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.33.0/docs/resources/prometheus_scraper#amp PrometheusScraper#amp}
	Amp interface{} `field:"optional" json:"amp" yaml:"amp"`
}

