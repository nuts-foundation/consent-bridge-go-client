package main

import (
	"github.com/nuts-foundation/consent-bridge-go-client/engine"
	"github.com/nuts-foundation/nuts-go-core/docs"
)

func main() {
	docs.GenerateConfigOptionsDocs("README_options.rst", engine.NewConsentBridgeClientEngine().FlagSet)
}
