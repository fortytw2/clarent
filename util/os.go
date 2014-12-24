package util

import (
	"fmt"
	"io"
	"os"
)

// Prints an error to stdout in a "nice" way.
func DumpError(e error) {
	FDumpError(os.Stdout, e)
}

// Prints an error to a writer in a "nice" way.
func FDumpError(w io.Writer, e error) {
	fmt.Fprintf(w, "clarent: Error: %s\n", e.Error())
}
