package templates

import (
	"embed"
)

//go:embed main.gohtml
var MainFile string

//go:embed components/*.gohtml
var Components embed.FS
