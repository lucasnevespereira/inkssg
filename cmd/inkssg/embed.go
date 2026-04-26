package main

import (
	"embed"

	"github.com/snowztech/inkssg"
)

//go:embed themes/minimal/layout.html
//go:embed themes/minimal/styles.css
//go:embed themes/landing/layout.html
//go:embed themes/landing/styles.css
var themes embed.FS

func init() {
	inkssg.SetThemes(themes)
}
