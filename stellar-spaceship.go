package main

import (
	"github.com/urfave/cli"
	"log"
	"os"
)

func main() {
	app := cli.NewApp()

	app.Name = "stellar-spaceship"
	app.Usage = "small terminal application for Stellar-based utilities"
	app.Version = "0.0.1"
	app.Commands = []cli.Command{
		{
			Name:    "gen",
			Aliases: []string{"g"},
			Usage:   "generate a new valid Stellar keypair",
			Action:  genKeypair,
		},
		{
			Name:    "parse",
			Aliases: []string{"p"},
			Usage:   "return the corresponding Stellar address from a valid `SEED` string",
			Action:  parseAddressOrSeed,
		},
		{
			Name:    "derive",
			Aliases: []string{"d"},
			Usage:   "derive a `RANGE` of keypairs from a desired `SEED`",
			Flags: []cli.Flag{
				cli.StringFlag{
					Name:  "seed, s",
					Usage: "set the `SEED` from which to derive from",
				},
				cli.StringFlag{
					Name:  "range, r",
					Usage: "set the `RANGE` of key_pairs to derive (e.g. 0..9)",
				},
			},
			Action: deriveKeypairsFromSeed,
		},
	}

	err := app.Run(os.Args)

	if err != nil {
		log.Fatal(err)
	}
}
