package main

import (
	"fmt"
	"log"
	"os"

	"github.com/urfave/cli/v2"
)

var (
	file_flag = &cli.StringFlag{
		Name:    "file",
		Aliases: []string{"f"},
		Usage:   "Path to YAML file with the responder config",
		Value:   "responders.yml",
	}

	host_flag = &cli.StringFlag{
		Name:    "bind",
		Aliases: []string{"b"},
		Usage:   "HTTP Bind Address",
		Value:   "127.0.0.1",
	}

	port_flag = &cli.IntFlag{
		Name:    "port",
		Aliases: []string{"p"},
		Usage:   "HTTP Bind Port",
		Value:   8080,
	}

	cert_flag = &cli.StringFlag{
		Name:    "cert",
		Aliases: []string{"c"},
		Usage:   "Path to TLS Certificate (Public Key)",
		Value:   "tls.crt",
	}

	key_flag = &cli.StringFlag{
		Name:    "key",
		Aliases: []string{"k"},
		Usage:   "Path to TLS Secret Key (Private Key)",
		Value:   "tls.key",
	}
)

func main() {
	app := &cli.App{
		//Name:  "mock-http-server",
		Usage:       "Mock HTTP Responses",
		Description: "Serve mock HTTP endpoints based on a mapping between requests and responses.",
		Commands: []*cli.Command{
			{
				Name:  "serve",
				Usage: "Serve unencrypted HTTP",
				Flags: []cli.Flag{
					file_flag,
					host_flag,
					port_flag,
				},
				Action: func(c *cli.Context) error {
					fmt.Println("Serve HTTP")
					return nil
				},
				Subcommands: []*cli.Command{
					{
						Name:  "tls",
						Usage: "Serve encrypted HTTP",
						Flags: []cli.Flag{
							cert_flag,
							key_flag,
						},
						Action: func(c *cli.Context) error {
							fmt.Println("Serve HTTPS")
							return nil
						},
					},
				},
			},
			{
				Name:  "list",
				Usage: "Print request -> response mappings",
				Flags: []cli.Flag{
					file_flag,
				},
				Action: func(c *cli.Context) error {
					fmt.Println("List mappings")
					return nil
				},
			},
		},
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}