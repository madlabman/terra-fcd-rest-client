package factory

const (
	FCDHost   string = "fcd.terra.dev"
	FCDScheme string = "https"
)

// FCDEndpoint is the Full Client Daemon endpoint: https://fcd.terra.dev.
var FCDEndpoint = Endpoint{
	Host:    FCDHost,
	Schemes: []string{FCDScheme},
}

// Endpoint represents a single destination for client requests.
type Endpoint struct {
	Host    string
	Schemes []string
}
