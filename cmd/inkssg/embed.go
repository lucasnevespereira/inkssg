package main

import (
	"embed"

	"github.com/snowztech/inkssg"
)

//go:embed themes/minimal/layout.html
//go:embed themes/minimal/styles.css
var minimalTheme embed.FS

func init() {
	inkssg.SetMinimalTheme(minimalTheme)
}
