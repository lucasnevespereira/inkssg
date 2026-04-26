package main

import inkssg "github.com/snowztech/inkssg"

func main() {
	if err := inkssg.Build(); err != nil {
		panic(err)
	}
}
