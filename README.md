# stellar-spaceship
[![Go Report Card](https://goreportcard.com/badge/github.com/edunuzzi/stellar-spaceship)](https://goreportcard.com/report/github.com/edunuzzi/stellar-spaceship)

Small cli application for Stellar-based utilities.

Released under the terms of the MIT LICENSE.

## Installation

To install, just run 
```bash
go get github.com/edunuzzi/stellar-spaceship
```

Alternatively, you can [download](https://github.com/edunuzzi/stellar-spaceship/releases) a compiled binary for your OS and paste it in a directory [included in your $PATH](https://unix.stackexchange.com/questions/26047/how-to-correctly-add-a-path-to-path). 

## Commands

### `gen` | `g`

Generates a new Keypair and prints its Seed and Address

#### Usage
- `stellar-spaceship gen` 
- `stellar-spaceship g` 

### `parse` | `p`

Parses a Stellar Seed and print its corresponding Address

#### Usage
- `stellar-spaceship parse <valid Stellar SEED>`
- `stellar-spaceship p <valid Stellar SEED>`

### `derive` | `d`

Derives a `RANGE` of keypairs from a desired `SEED`

#### Flags

- `seed`: Sets the `SEED` from which to derive from
- `range`: Sets the `RANGE` of key_pairs to derive (e.g. 0..9) 

#### Usage
- `stellar-spaceship derive -seed <valid Stellar SEED> -range 0..9`
- `stellar-spaceship d -s <valid Stellar SEED> -r 0..9`
