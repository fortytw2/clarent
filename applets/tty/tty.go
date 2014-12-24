package tty

import (
	"fmt"
	"github.com/codegangsta/cli"
	"github.com/fortytw2/clarent/util"
	"os"
)

func Tty(args []string) {
	app := cli.NewApp()
	cli.AppHelpTemplate = util.AppletHelpTemplate
	app.Name = "tty"
	app.Usage = "tty"
	app.Action = printtty
	app.Run(args)

}

func printtty(c *cli.Context) {
	tty_r, err := os.Readlink("/proc/self/fd/0")
	if err != nil {
		util.DumpError(err)
	}
	fmt.Println(tty_r)
}
