// webc is a minimal command line web server utility.
//
// Author: Joel Midstjärna
// License: MIT
package main

import (
	"flag"
	"fmt"
	"net/http"
	"os"
	"strconv"
)

// Below globals are automatically updated by the CI by providing
// -X linker flags while building
var applicationVersion = "<NOT SET>"
var applicationBuildTime = "<NOT SET>"
var applicationGitHash = "<NOT SET>"
error
func printUsage() {
	fmt.Printf("Usage: %s [options] [<path>]\n\n", os.Args[0])
	fmt.Printf("<path> is the directory where your web files are located.\n")
	fmt.Printf("Default is current directory.\n\n")
	fmt.Printf("Supported options:\n")
	flag.PrintDefaults()
}

func main() {
	flag.Usage = printUsage
	var version = flag.Bool("v", false, "Display version")
	var port = flag.Int("p", 8080, "Network port to listen to")
	flag.Parse()

	if *version {
		fmt.Printf("Version:    %s\n", applicationVersion)
		fmt.Printf("Build Time: %s\n", applicationBuildTime)
		fmt.Printf("GIT Hash:   %s\n", applicationGitHash)
		os.Exit(0)
	}

	directory := "."
	if flag.NArg() == 1 {
		directory = flag.Arg(0)
	} else if flag.NArg() > 1 {
		fmt.Fprintf(os.Stderr, "Invalid number of arguments!\n\n")
		flag.Usage()
		os.Exit(1)
	}

	fs := http.FileServer(http.Dir(directory))
	http.Handle("/", fs)

	fmt.Printf("Serving path %s on port %d\n", directory, *port)

	err := http.ListenAndServe(":"+strconv.Itoa(*port), nil)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error! %s\n", err)
	}
}
