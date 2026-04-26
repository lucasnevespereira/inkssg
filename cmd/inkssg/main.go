package main

import (
	"fmt"
	"os"
)

const version = "0.0.1"

func main() {
	if len(os.Args) < 2 {
		usage()
		os.Exit(1)
	}

	switch os.Args[1] {
	case "version", "--version", "-v":
		fmt.Println("inkssg", version)
	case "build":
		fmt.Println("not implemented yet")
		os.Exit(1)
	default:
		usage()
		os.Exit(1)
	}
}

func usage() {
	fmt.Println("inkssg — a small static site generator")
	fmt.Println()
	fmt.Println("Usage:")
	fmt.Println("  inkssg build      build the site into ./public")
	fmt.Println("  inkssg version    print version")
}
