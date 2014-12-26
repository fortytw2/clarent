package sys

import (
	"fmt"
	"github.com/codegangsta/cli"
	"github.com/fortytw2/clarent/util"
	"os"
)

func Hostname(args []string) {
	app := cli.NewApp()
	cli.AppHelpTemplate = util.AppletHelpTemplate
	app.Name = "hostname"
	app.Usage = "hostname"
	app.Action = func(c *cli.Context) {
		hostname, err := os.Hostname()
		if err != nil {
			fmt.Errorf("%!", err)
			os.Exit(1)
		}
		fmt.Println(hostname)
	}

	app.Run(args)
}
