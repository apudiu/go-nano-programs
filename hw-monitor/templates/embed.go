package templates

import (
	"embed"
	"html/template"
)

//go:embed main.gohtml
var MainFile string

//go:embed components/*.gohtml
var Components embed.FS

func GetTemplate(name string, funcMap ...template.FuncMap) (*template.Template, error) {
	fb, err := Components.ReadFile(name)
	if err != nil {
		return nil, err
	}

	t := template.New(name)

	if funcMap != nil {
		t.Funcs(funcMap[0])
	}

	return t.Parse(string(fb))
}
