package ch4

import (
	"crypto/sha256"
	"crypto/sha512"
	"fmt"
	"gopkg.in/urfave/cli.v2"
	"os"
	"sort"
)

// SHACalculator is a simple SHA checksum calculator application
func SHACalculator() {
	app := &cli.App{
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:    "input",
				Value:   "",
				Aliases: []string{"i"},
				Usage:   "source that will be sha",
			},
			&cli.IntFlag{
				Name:    "mode",
				Value:   256,
				Aliases: []string{"m"},
				Usage:   "sha mode, 256, 384, 512",
			},
		},
		Name:    "shaprog",
		Usage:   "calculate sha sum",
		Version: "v1.0.0",
		Action: func(c *cli.Context) error {
			mode := c.Int("mode")
			input := c.String("input")
			if input == "" {
				return cli.Exit("input empty", 1)
			}
			var sum interface{}
			switch mode {
			case 256:
				sum = sha256.Sum256([]byte(input))
			case 384:
				sum = sha512.Sum384([]byte(input))
			case 512:
				sum = sha512.Sum512([]byte(input))
			default:
				cli.Exit("invalid mode, must be 256, 384 or 512", 2)
			}
			fmt.Printf("sha%d\n%x\n", mode, sum)
			return nil
		},
	}

	sort.Sort(cli.FlagsByName(app.Flags))
	app.Run(os.Args)
}
