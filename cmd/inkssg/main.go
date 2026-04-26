package main

import (
	"fmt"
	"os"

	"github.com/snowztech/inkssg"
)

const version = "0.1.0"

func main() {
	if len(os.Args) < 2 {
		usage()
		os.Exit(1)
	}

	switch os.Args[1] {
	case "version", "--version", "-v":
		fmt.Println("inkssg", version)
	case "build":
		path := "."
		if len(os.Args) > 2 {
			path = os.Args[2]
		}
		if err := inkssg.Build(path); err != nil {
			fmt.Fprintf(os.Stderr, "error: %v\n", err)
			os.Exit(1)
		}
	default:
		usage()
		os.Exit(1)
	}
}

func usage() {
	fmt.Println("inkssg — a small static site generator")
	fmt.Println()
	fmt.Println("Usage:")
	fmt.Println("  inkssg build [path]   build the site into ./public")
	fmt.Println("  inkssg version        print version")
}
