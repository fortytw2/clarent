package cat

import (
	"fmt"
	"github.com/codegangsta/cli"
	"github.com/fortytw2/clarent/util"
	"io/ioutil"
)

func Cat(args []string) {
	app := cli.NewApp()
	cli.AppHelpTemplate = util.AppletHelpTemplate
	app.Name = "cat"
	app.Usage = "cat [file 1] [file 2] [...]"
	app.Action = concat

	app.Run(args)
}

func concat(c *cli.Context) {
	for _, file := range c.Args() {
		f, err := ioutil.ReadFile(file)
		if err != nil {
			util.DumpError(err)
		}
		// print out the file
		fmt.Print(string(f))
	}
}
