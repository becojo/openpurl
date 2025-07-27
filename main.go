package main

import (
	_ "embed"
	"fmt"
	"log"
	"os"
	"os/exec"
	"runtime"

	"github.com/urfave/cli/v2"
)

var usage = "usage: openpurl [--print|-p] [pkg:...]"

func main() {
	var print bool

	app := &cli.App{
		Name:  "openpurl",
		Usage: "Open canonical web URL of a package URL",
		Flags: []cli.Flag{
			&cli.BoolFlag{
				Name:        "print",
				Usage:       "print the URL instead of opening it",
				Aliases:     []string{"p"},
				Destination: &print,
			},
		},
		Action: func(ctx *cli.Context) error {
			var arg = ctx.Args().Get(0)

			if arg == "" {
				fmt.Println(usage)
				return nil
			}

			purl, err := Parse(arg)

			if err != nil {
				fmt.Fprintln(os.Stderr, err)
				return nil
			}

			url, err := DefaultConfig.Render(purl)
			if err != nil {
				fmt.Fprintln(os.Stderr, err)
				return nil
			}

			if print {
				fmt.Println(url)
			} else {
				open := "xdg-open"

				switch runtime.GOOS {
				case "windows":
					open = "start"
				case "darwin":
					open = "open"
				}

				exec.Command(open, url).Run()
			}

			return nil
		},
	}

	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}
