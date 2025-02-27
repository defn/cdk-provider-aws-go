package appmeshvirtualgateway


type AppmeshVirtualGatewaySpecListenerConnectionPoolHttp2 struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.33.0/docs/resources/appmesh_virtual_gateway#max_requests AppmeshVirtualGateway#max_requests}.
	MaxRequests *float64 `field:"required" json:"maxRequests" yaml:"maxRequests"`
}

