package main

import (
	"os"

	"github.com/passw0rd/cli/cmd/demo"

	"gopkg.in/urfave/cli.v2"
)

func main() {
	app := &cli.App{
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:    "token",
				Aliases: []string{"t"},
				Usage:   "Auth token",
				EnvVars: []string{"PASSW0RD_TOKEN"},
			},
			&cli.StringFlag{
				Name:    "appid",
				Aliases: []string{"appId"},
				Usage:   "App ID",
				EnvVars: []string{"PASSW0RD_APP_ID"},
			},
			&cli.StringFlag{
				Name:    "public_key",
				Aliases: []string{"pk"},
				Usage:   "Service public key",
				EnvVars: []string{"PASSW0RD_PUB"},
			},
			&cli.StringFlag{
				Name:    "secret_key",
				Aliases: []string{"sk"},
				Usage:   "Client private key",
				EnvVars: []string{"PASSW0RD_SECRET"},
			},
		},
		Commands: []*cli.Command{
			{
				Name:  "demo",
				Usage: "Try service without writing code",
				Subcommands: []*cli.Command{
					demo.Enroll(),
				},
			},
		},
	}

	app.Run(os.Args)
}
