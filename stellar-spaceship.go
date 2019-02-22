package main

import (
	"fmt"
	"github.com/stellar/go/keypair"
	"github.com/urfave/cli"
	"log"
	"os"
)

func main() {
	app := cli.NewApp()

	app.Name = "stellar-spaceship"
	app.Usage = ""
	app.Commands = []cli.Command{
		{
			Name: "gen",
			Aliases: []string{"g"},
			Usage: "generate a new valid Stellar keypair",
			Action: func(c *cli.Context) error {
				k, err := keypair.Random()

				if err != nil {
					return fmt.Errorf("there was an error generating the keypair, try again")
				}

				fmt.Printf("New Seed:    %s\n", k.Seed())
				fmt.Printf("New Address: %s\n", k.Address())

				return nil
			},
		},
		{
			Name: "parse",
			Aliases: []string{"p"},
			Usage: "return the corresponding Stellar address from a valid `SEED` string",
			Action: func(c *cli.Context) error {
				k, err := keypair.Parse(string(c.Args()[0]))

				if err != nil {
					return fmt.Errorf("the string provided was not a valid Stellar seed or address")
				}

				fmt.Printf("Address: %s\n", k.Address())

				return nil
			},
		},
	}

	err := app.Run(os.Args)

	if err != nil {
		log.Fatal(err)
	}
}
