// dedupPATH is a simple command-line utility to remove duplicate entries from
// the $PATH environment variable on linux.
package main

import (
	"fmt"
	"os"
	"strings"

	"github.com/spf13/pflag"
)

var (
	sep string
)

func getpath() string {
	return os.Getenv("PATH")
}

// removeDuplicates slices env into seperate strings seperated by sep, then ranges
// over the resulting slice, removing duplicates. If sep is is not encountered
// after parsing env, or another error occurs env is returned unmodified.
func removeDuplicates(env, sep string) string {
	paths := strings.Split(env, sep)
	// present represents a 'rollcall'; When a unique string is encountered,
	// it is recorded in present, so that duplicates can be detected and
	// skipped.
	var present = make(map[string]bool)

	for _, p := range paths {
		if _, exists := present[p]; !exists {
			present[p] = true
			continue
		}

		// has been seen already

	}

	// since we know that `present` contains all the needed filepaths,
	// we just range over `present`, and append them all together.

	var newp []string
	for fp := range present {
		newp = append(newp, fp)

	}

	// build new env from the resulting collection seperated by `sep`
	return strings.Join(newp, sep)
}

func doflags() {
	pflag.StringVarP(&sep, "sep", "s", "", "Use sep as a sepatator while parsing input string.")
	pflag.Parse()
}
func main() {
	var initial = os.Args[1]
	// if os.Args[1] is the literal string "path", $PATH env var is used.
	if initial == "path" {
		start := os.Getenv("$PATH")
		final := removeDuplicates(start, ":")
		fmt.Fprintf(os.Stdout, "%s\n", final)
	}
	if len(sep) > 0 {

	}
}
