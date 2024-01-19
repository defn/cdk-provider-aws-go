package prometheusscraper


type PrometheusScraperSource struct {
	// eks block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.33.0/docs/resources/prometheus_scraper#eks PrometheusScraper#eks}
	Eks interface{} `field:"optional" json:"eks" yaml:"eks"`
}

