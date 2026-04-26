package main

import (
	"fmt"
	"os"

	"github.com/snowztech/inkssg"
)

func main() {
	if err := inkssg.Build("."); err != nil {
		fmt.Fprintf(os.Stderr, "build failed: %v\n", err)
		os.Exit(1)
	}
}
