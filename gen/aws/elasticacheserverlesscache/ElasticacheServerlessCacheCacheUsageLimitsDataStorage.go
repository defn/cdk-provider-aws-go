package elasticacheserverlesscache


type ElasticacheServerlessCacheCacheUsageLimitsDataStorage struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.33.0/docs/resources/elasticache_serverless_cache#maximum ElasticacheServerlessCache#maximum}.
	Maximum *float64 `field:"required" json:"maximum" yaml:"maximum"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.33.0/docs/resources/elasticache_serverless_cache#unit ElasticacheServerlessCache#unit}.
	Unit *string `field:"required" json:"unit" yaml:"unit"`
}

