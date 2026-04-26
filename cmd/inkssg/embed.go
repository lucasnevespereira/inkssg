package main

import (
	"embed"

	"github.com/snowztech/inkssg"
)

//go:embed themes/minimal/layout.html
//go:embed themes/minimal/styles.css
//go:embed themes/devtool/layout.html
//go:embed themes/devtool/styles.css
var themes embed.FS

func init() {
	inkssg.SetThemes(themes)
}
