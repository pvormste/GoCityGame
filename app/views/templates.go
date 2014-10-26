package views

import (
	"html/template"
	"vormstein.eu/gocitygame/util"
)

// Templates-Map
var Templates map[string]*template.Template

// Create Templates with inheritance
func loadTemplates() {
	Templates["index.html"] = createInheritedTemplate("base.tmpl", "index.tmpl")
}

func createInheritedTemplate(base string, child string) (t *template.Template) {
	return template.Must(template.ParseFiles("./app/views/"+util.Conf.Tmpl+"/"+child, "./app/views/"+util.Conf.Tmpl+"/"+base))
}

func init() {
	Templates = make(map[string]*template.Template)
	loadTemplates()
}
