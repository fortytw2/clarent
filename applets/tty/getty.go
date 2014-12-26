package tty

import (
	"fmt"
	"github.com/codegangsta/cli"
	"github.com/fortytw2/clarent/util"
	"os"
	"os/exec"
)

// weirdness here - all getty really has to do for us is spawn a process with a
// certain tty as std{in,out,err} - seems to accomplish as much as mingetty.
// this is probably because serial consoles aren't a thing anymore.
// granted, this might be completely the wrong thing to do, but it seems to work
func Getty(args []string) {
	app := cli.NewApp()
	cli.AppHelpTemplate = util.AppletHelpTemplate
	app.Name = "getty"
	app.Usage = "getty [/dev/tty*] [command to run]"
	app.Action = getty

	app.Run(args)
}

func getty(c *cli.Context) {

	if len(c.Args()) < 2 {
		cli.ShowAppHelp(c)
	}

	// open our character device
	tty, err := os.OpenFile(c.Args().Get(0), os.O_RDWR, 0)
	if err != nil {
		fmt.Println("unable to obtain tty device, try running as root")
	}

	// swap std{in,out,err}
	os.Stdout = tty
	os.Stderr = tty
	os.Stdin = tty

	hostname, _ := os.Hostname()

	fmt.Println("Welcome to ", hostname, "running at ", c.Args().Get(0))
	// run the command on the tty. init is responsible for respawning getty if
	// it dies, getty doesn't do any respawning
	cmd := exec.Command(c.Args().Get(1))
	cmd.Stdin = tty
	cmd.Stdout = tty
	cmd.Stderr = tty
	cmd.Run()

}
