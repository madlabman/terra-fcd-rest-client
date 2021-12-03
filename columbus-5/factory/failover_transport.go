package factory

import (
	"fmt"

	"github.com/go-openapi/runtime"
	openapiTransport "github.com/go-openapi/runtime/client"
	"github.com/sirupsen/logrus"
)

// failoverTransport is the runtime.ClientTransport implementation that supports multiple endpoints
// and queries them one after another until either request succeeds or all endpoints fail.
type failoverTransport struct {
	logger    *logrus.Logger
	endpoints []runtime.ClientTransport
}

func newFailoverTransport(logger *logrus.Logger, endpoints []Endpoint, basepath string) runtime.ClientTransport {
	out := &failoverTransport{
		logger: logger,
	}

	for _, endpoint := range endpoints {
		out.endpoints = append(
			out.endpoints,
			openapiTransport.New(endpoint.Host, basepath, endpoint.Schemes),
		)
	}
	return out
}

func (t *failoverTransport) Submit(operation *runtime.ClientOperation) (resp interface{}, err error) {
	// This is a naive implementation (successively try all endpoints until one of them works).
	// Nothing fancy, but it does its job.
	for endpointID, transport := range t.endpoints {
		resp, err = transport.Submit(operation)
		if err != nil {
			endpoint := transport.(*openapiTransport.Runtime).Host
			t.logger.Errorf("failed to Submit to endpoint #%d (%s): %s",
				endpointID, endpoint, err)
			continue
		}
		return
	}
	return resp, fmt.Errorf("failed to Submit (all retries failed): %w", err)
}
