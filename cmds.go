package main

import (
	"fmt"
	"github.com/stellar/go/exp/crypto/derivation"
	"github.com/stellar/go/keypair"
	"github.com/urfave/cli"
	"strconv"
	"strings"
)

func genKeypair(c *cli.Context) error {
	k, err := keypair.Random()

	if err != nil {
		return fmt.Errorf("there was an error generating the keypair, try again")
	}

	fmt.Printf("New Seed:    %s\n", k.Seed())
	fmt.Printf("New Address: %s\n", k.Address())

	return nil
}

func parseAddressOrSeed(c *cli.Context) error {
	k, err := keypair.Parse(string(c.Args()[0]))

	if err != nil {
		return fmt.Errorf("the string provided was not a valid Stellar seed")
	}

	fmt.Printf("Address: %s\n", k.Address())

	return nil
}

func deriveKeypairsFromSeed(c *cli.Context) error {
	seed := c.String("seed")

	if len(seed) == 0 {
		return fmt.Errorf("please inform a valid `SEED`")
	}

	rangeStr := strings.Split(c.String("range"), "..")
	start, err := strconv.Atoi(rangeStr[0])

	if err != nil {
		return fmt.Errorf("invalid start for range")
	}

	end, err := strconv.Atoi(rangeStr[1])

	if err != nil {
		return fmt.Errorf("invalid end for range")
	}

	for i := start; i <= end; i++ {

		path := fmt.Sprintf(derivation.StellarAccountPathFormat, i)

		key, err := derivation.DeriveForPath(path, []byte(fmt.Sprintf("%s", seed)))
		if err != nil {
			return fmt.Errorf("there was an error trying to derive ")
		}

		kp, err := keypair.FromRawSeed(key.RawSeed())

		fmt.Printf("KeyPair %d:\nSeed: %s\nAddress: %s\n-----------------\n", i, kp.Seed(), kp.Address())
	}

	return nil
}
