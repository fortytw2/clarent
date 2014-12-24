package main

// Applet imports
import (
	"github.com/fortytw2/clarent/applets/cat"
	"github.com/fortytw2/clarent/applets/env"
)

// This map contains the mappings from callname
// to applet function.
var Applets map[string]Applet = map[string]Applet{
	"cat": cat.Cat,
	"env": env.Env,
}

// Signature of applet functions.
// call is like os.Argv, and therefore contains the
// name of the applet itself in call[0].
// If the returned error is not nil, it is printed
// to stdout.
type Applet func(call []string) error
