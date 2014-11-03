package views

import (
	"github.com/pvormste/atempgo"
	"html/template"
	"vormstein.eu/gocitygame/app"
)

func init() {
	parseOptions := &atempgo.ParseOptions{Ext: "tmpl"}

	atempgo.LoadTemplates("app/views/"+app.Config.Tmpl, parseOptions)
}

func GetTemplate(key string) *template.Template {
	return atempgo.GetTemplate(key)
}
