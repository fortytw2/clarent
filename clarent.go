package main

import (
	"fmt"
	"github.com/codegangsta/cli"
	"github.com/fortytw2/clarent/util"
	"os"
	"path/filepath"
)

func Clarent(args []string) error {
	app := cli.NewApp()
	app.Name = "clarent"
	app.Usage = "complete userland in a box"
	app.Action = cli.ShowAppHelp
	app.Version = "0.0.1-dev"

	cli.AppHelpTemplate = util.MainHelpTemplate

	app.Commands = []cli.Command{
		{
			Name:      "install",
			ShortName: "i",
			Usage:     "create symlinks in current directory",
			Action:    install,
		},
		{
			Name:      "list",
			ShortName: "l",
			Usage:     "print a list of compiled applets",
			Action:    list,
		},
	}

	app.Run(args)

	return nil
}

func list(c *cli.Context) {
	fmt.Println("APPLETS:")
	for name := range Applets {
		print(name, ", ")
	}
	fmt.Println("")
}

func install(c *cli.Context) {

	wd, e := os.Getwd()
	if e != nil {
		util.DumpError(e)
	}
	relpath := filepath.Join(wd, "clarent")
	binpath, e := filepath.Abs(relpath)
	if e != nil {
		util.DumpError(e)
	}

	for name := range Applets {
		// Don't overwrite the executable
		if name == "clarent" {
			continue
		}

		newpath := filepath.Join(wd, name)
		e = os.Symlink(binpath, newpath)
		if e != nil {
			util.DumpError(e)
		}
	}
}
