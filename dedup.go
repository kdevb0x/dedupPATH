// dedupPATH is a simple command-line utility to remove duplicate entries from
// the $PATH environment variable on linux.
package main

import "os"

func getpath() string {
	return os.Getenv("PATH")
}

func main() {

}
