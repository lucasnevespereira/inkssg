package main

import (
	"fmt"
	"os"
	"runtime/debug"

	"github.com/snowztech/inkssg"
)

var version = "dev"

func resolveVersion() string {
	if version != "dev" {
		return version
	}
	if info, ok := debug.ReadBuildInfo(); ok && info.Main.Version != "" && info.Main.Version != "(devel)" {
		return info.Main.Version
	}
	return "dev"
}

func main() {
	if len(os.Args) < 2 {
		usage()
		os.Exit(1)
	}

	switch os.Args[1] {
	case "version", "--version", "-v":
		fmt.Println("inkssg", resolveVersion())
	case "build":
		path := "."
		if len(os.Args) > 2 {
			path = os.Args[2]
		}
		if err := inkssg.Build(path); err != nil {
			fmt.Fprintf(os.Stderr, "error: %v\n", err)
			os.Exit(1)
		}
	case "new":
		path := "."
		if len(os.Args) > 2 {
			path = os.Args[2]
		}
		if err := inkssg.Scaffold(path); err != nil {
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
	fmt.Println("  inkssg new [path]     scaffold a new site")
	fmt.Println("  inkssg build [path]   build the site into ./public")
	fmt.Println("  inkssg version        print version")
}
