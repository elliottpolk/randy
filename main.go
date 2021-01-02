package main

import (
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"time"

	cli "github.com/urfave/cli/v2"
)

const (
	all      = "all"
	alpha    = "alpha"
	numeric  = "numeric"
	alphanum = "alphanum"

	letters  = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
	numbers  = "1234567890"
	specials = "!@#$%^&*_"
)

var (
	version  string
	compiled string = fmt.Sprint(time.Now().Unix())
	githash  string

	countFlag = &cli.IntFlag{
		Name:    "character-count",
		Aliases: []string{"count", "cnt", "c"},
		Value:   32,
		Usage:   "number of characters to produce",
	}

	alphasFlag = &cli.BoolFlag{
		Name:    "alpha-characters",
		Aliases: []string{"alphas", "a"},
		Value:   true,
		Usage:   "include (English) alphabetical characters",
	}

	numericsFlag = &cli.BoolFlag{
		Name:    "numerical-characters",
		Aliases: []string{"numerics", "n"},
		Value:   true,
		Usage:   "include numerical characters",
	}

	specialFlag = &cli.BoolFlag{
		Name:    "special-characters",
		Aliases: []string{"specials", "s"},
		Value:   true,
		Usage:   "include special characters",
	}
)

func main() {
	ct, err := strconv.ParseInt(compiled, 0, 0)
	if err != nil {
		panic(err)
	}

	app := cli.App{
		Name:      "randy",
		Copyright: "Copyright © 2021",
		Version:   fmt.Sprintf("%s | compiled %s | commit %s", version, time.Unix(ct, -1).Format(time.RFC3339), githash),
		Compiled:  time.Unix(ct, -1),
		Usage:     "generates a random string of characters",
		UsageText: "randy [options] [arguments...]",
		Flags: []cli.Flag{
			countFlag,
			alphasFlag,
			numericsFlag,
			specialFlag,
		},
		Action: func(ctx *cli.Context) error {
			var (
				cnt = ctx.Int(countFlag.Name)

				chars = []rune{}
				out   = []rune{}
			)

			if cnt < 1 {
				return cli.Exit("a minimum of 1 character must be specified", 1)
			}

			if ctx.Bool(alphasFlag.Name) {
				chars = append(chars, []rune(letters)...)
			}

			if ctx.Bool(numericsFlag.Name) {
				chars = append(chars, []rune(numbers)...)
			}

			if ctx.Bool(specialFlag.Name) {
				chars = append(chars, []rune(specials)...)
			}

			if len(chars) < 1 {
				return cli.Exit("at least 1 of letters, numbers, or, special characters must be allowed", 1)
			}

			rand.Seed(time.Now().UnixNano())
			for len(out) < cnt {
				out = append(out, chars[rand.Intn(len(chars))])
			}

			fmt.Print(string(out))

			return nil
		},
	}

	app.Run(os.Args)
}
