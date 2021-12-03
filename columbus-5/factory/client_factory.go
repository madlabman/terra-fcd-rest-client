package factory

import (
	client "github.com/lidofinance/terra-fcd-rest-client/columbus-5/client"

	"github.com/sirupsen/logrus"
)

// NewDefaultClient creates a new REST API client to https://fcd.terra.dev.
func NewDefaultClient() *client.TerraRESTApis {
	return NewClient(FCDEndpoint, client.DefaultBasePath)
}

// NewClient creates a new REST API client to a columbus-5 Terra node.
func NewClient(endpoint Endpoint, basepath string) *client.TerraRESTApis {
	return client.NewHTTPClientWithConfig(nil, &client.TransportConfig{
		Host:     endpoint.Host,
		BasePath: basepath,
		Schemes:  endpoint.Schemes,
	})
}

// NewFailoverClient creates a new columbus-5 Terra REST API client that can have addresses of backup
// nodes to interact with in order to improve accessibility.
func NewFailoverClient(logger *logrus.Logger, endpoints []Endpoint, basepath string) *client.TerraRESTApis {
	return client.New(newFailoverTransport(logger, endpoints, basepath), nil)
}
