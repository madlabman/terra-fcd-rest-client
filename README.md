# Terra FCD REST client

A library with Terra FCD REST client generated using the official swagger documentation.

## columbus-5

The columbus-5 folder corresponds to the client and models generated using the columbus-5 swagger specification with a couple of specification mistake fixes.

### Instantiation

To create a client instance, import the factory package and use one of the suggested constructors:
```go
package main

import (
	"github.com/lidofinance/terra-fcd-rest-client/columbus-5/client"
	"github.com/lidofinance/terra-fcd-rest-client/columbus-5/factory"
	"github.com/sirupsen/logrus"
)

const (
	bombayFCDHost   string = "bombay-fcd.terra.dev"
	bombayFCDScheme string = "https"
)

var bombayFCDEndpoint = factory.Endpoint{
	Host:    bombayFCDHost,
	Schemes: []string{bombayFCDScheme},
}

func main() {
	mainnetClient := factory.NewDefaultClient()
	bombayClient := factory.NewClient(bombayFCDEndpoint, client.DefaultBasePath)
	severalFCDsClient := factory.NewFailoverClient(logrus.New(), []factory.Endpoint{{
		Host:    "fcd.ahost.dev",
		Schemes: []string{"https"},
	}, {
		Host:    "fcd.anotherhost.dev",
		Schemes: []string{"https"},
	}, {
		Host:    "fcd.yetanotherhost.dev",
		Schemes: []string{"http", "https"},
		// and so on
	}}, client.DefaultBasePath)
}
}
```

### List of specification fixes

During continuous interaction with Terra FCD we have been encountering different swagger inaccuracies which required to apply fixes to the specification. The list of fixes is presented in the [fixes.md](https://github.com/lidofinance/terra-fcd-rest-client/blob/master/fixes.md). They have been taken from shortened and modified Terra FCD swagger specifications from [terra-monitors](https://github.com/lidofinance/terra-monitors) and [terra-bots](https://github.com/lidofinance/terra-bots) repositories.

Besides, there are other changes that have been applied to the specification, but they haven't modified any data representation. They are about data definitions and references inbetween types that led to compilation errors.

## Generate custom code

If you need to make changes to the code, make changes to the specification and generate a new version of the client:
1. Choose the desired version folder
2. Make changes to the specification
3. Install [go-swagger](https://github.com/go-swagger/go-swagger)
4. Remove the `client` and `models` folders from the version folder
5. Run the generate script:
```
make generate
```
