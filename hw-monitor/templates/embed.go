package templates

import (
	"embed"
)

//go:embed main.html
var MainFile []byte

//go:embed components/*.gohtml
var Components embed.FS
