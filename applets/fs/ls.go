package fs

import (
	"fmt"
	"github.com/codegangsta/cli"
	"github.com/fortytw2/clarent/util"
	"io/ioutil"
)

// ls lists files... should probably print them a bit prettier, but too bad
func Ls(args []string) {
	app := cli.NewApp()
	cli.AppHelpTemplate = util.AppletHelpTemplate
	app.Name = "ls"
	app.Usage = "ls [directory]"
	app.Action = ls

	app.Run(args)

}

func ls(c *cli.Context) {
	if len(c.Args()) > 0 {
		for _, dir := range c.Args() {
			files, _ := ioutil.ReadDir(dir)
			for _, f := range files {
				fmt.Println(f.Name())
			}
		}
	} else {
		files, _ := ioutil.ReadDir("./")
		for _, f := range files {
			fmt.Println(f.Name())
		}
	}
}
