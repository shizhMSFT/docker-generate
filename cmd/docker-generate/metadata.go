package main

import (
	"encoding/json"
	"os"

	"github.com/urfave/cli/v2"
)

var pluginMetadata = map[string]string{
	"SchemaVersion": "0.1.0",
	"Vendor":        "github.com/shizhMSFT",
	"Version":       "0.1.0",
}

var metadataCommand = &cli.Command{
	Name: "docker-cli-plugin-metadata",
	Action: func(ctx *cli.Context) error {
		writer := json.NewEncoder(os.Stdout)
		return writer.Encode(pluginMetadata)
	},
	Hidden: true,
}
