// dedupPATH is a simple command-line utility to remove duplicate entries from
// the $PATH environment variable on linux.
package main

import (
	"os"
	"strings"
)

func getpath() string {
	return os.Getenv("PATH")
}

func pathdedup(path string) string {
	var sep = strings.Split(path, ":")

	for i := 0; i < len(sep); i++ {
		for j, val := range sep[i+1:] {
			if sep[i] == val {
				sep = append(sep[i:], sep[j+1:]...)
			} else {
				continue
			}
		}
	}
	// concat the non-duplicate paths and return the deduped path
	return strings.Join(sep, ":")
}

func main() {

}
