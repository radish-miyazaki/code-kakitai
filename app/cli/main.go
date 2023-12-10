package main

import (
	"log"
	"os"

	"github.com/urfave/cli/v2"

	"github.com/radish-miyazaki/code-kakitai/infrastructure/mysql/db/schema"
)

func main() {
	app := &cli.App{
		Commands: []*cli.Command{
			{
				Name:  "migration",
				Usage: "migrate schema to database",
				Action: func(ctx *cli.Context) error {
					schemaFile := ctx.Args().Get(0)
					dryRun := ctx.Args().Get(1) != "apply"

					return schema.Migrate(schemaFile, dryRun)
				},
			},
		},
	}

	if err := app.Run(os.Args); err != nil {
		log.Panic(err)
	}
}
