package views

import (
	"html/template"
	"vormstein.eu/gocitygame/util"
)

// Templates-Map
var templates map[string]*template.Template

// ### Private

// Create Templates with inheritance
func loadTemplates() {
	templates["index"] = createInheritedTemplate("base.tmpl", "index.tmpl")
	templates["error"] = createInheritedTemplate("base.tmpl", "error.tmpl")
}

func createInheritedTemplate(base string, child string) *template.Template {
	return template.Must(template.ParseFiles("./app/views/"+util.Conf.Tmpl+"/"+child, "./app/views/"+util.Conf.Tmpl+"/"+base))
}

func init() {
	templates = make(map[string]*template.Template)
	loadTemplates()
}

// ### Public

func GetTemplate(name string) *template.Template {
	return templates[name]
}
