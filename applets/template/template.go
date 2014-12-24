package template

import (
	"fmt"
	"github.com/codegangsta/cli"
	"github.com/fortytw2/clarent/util"
)

func Clarent(args []string) error {
	app := cli.NewApp()
	cli.AppHelpTemplate = util.AppletHelpTemplate
	app.Name = "template"
	app.Usage = "what is a template?"
	app.Action = cli.ShowAppHelp

	app.Commands = []cli.Command{
		{
			Name:      "example",
			ShortName: "e",
			Usage:     "print some example text, then exit",
			Action:    example,
		},
	}

	app.Run(args)

	return nil
}

func example() {
	fmt.Println("I did things")
}
