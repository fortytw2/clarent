package env

import (
	"fmt"
	"os"
	"strings"
	"github.com/codegangsta/cli"
	"github.com/fortytw2/clarent/util"
)

func Env(args []string) error {
	app := cli.NewApp()
	cli.AppHelpTemplate = util.AppletHelpTemplate
	app.Name = "env"
	app.Usage = "view environment variables"
	app.Action = list

	app.Run(args)

	return nil
}

func list(c *cli.Context) {
	for _, e := range os.Environ() {
        pair := strings.Split(e, "=")
        fmt.Println(pair[0], ":", pair[1])
    }
}

