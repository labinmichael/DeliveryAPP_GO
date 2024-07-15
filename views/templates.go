// views/templates.go
package views

import (
	"html/template"
	"path/filepath"
)

type Templates struct {
	Index   *template.Template
	New     *template.Template
	Edit    *template.Template
	Partial *template.Template
}

func InitTemplates() *Templates {
	tmpl := &Templates{}
	tmpl.Index = parseTemplate("templates/index.html")
	tmpl.New = parseTemplate("templates/new.html")
	tmpl.Edit = parseTemplate("templates/edit.html")
	tmpl.Partial = parseTemplate("templates/partial.html")
	return tmpl
}

func parseTemplate(path string) *template.Template {
	return template.Must(template.ParseFiles(filepath.Join("views", path)))
}
