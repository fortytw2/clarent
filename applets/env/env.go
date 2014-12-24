package env

import (
	"fmt"
	"github.com/codegangsta/cli"
	"github.com/fortytw2/clarent/util"
	"os"
	"strings"
)

func Env(args []string) {
	app := cli.NewApp()
	cli.AppHelpTemplate = util.AppletHelpTemplate
	app.Name = "env"
	app.Usage = "view environment variables"
	app.Action = list

	app.Run(args)
}

func list(c *cli.Context) {
	for _, e := range os.Environ() {
		pair := strings.Split(e, "=")
		fmt.Println(pair[0], ":", pair[1])
	}
}
