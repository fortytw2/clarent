package main

import (
	"errors"
	"os"
	"path/filepath"
)

// neccessary to prevent initialization loop
func init() {
	Applets["clarent"] = Clarent
}

// perform the neccessary trickery to call applets based on Os.Args[0]
func main() {
	callname := filepath.Base(os.Args[0])

	applet, ok := Applets[callname]
	if !ok {
		panic(errors.New("Could not find \"" + callname + "\""))
	}

	// If the executable is called
	// check, if the second parameter is an applet name.
	// If so, call that applet instead
	args := os.Args
	if callname == "clarent" && len(args) >= 2 {
		subapplet, ok := Applets[args[1]]
		if ok {
			applet = subapplet
			args = args[1:]
		}
	}

	// run the applet according to callname
	applet(args)
}
