package views

import (
	"github.com/pvormste/atempgo"
	"html/template"
	"net/http"
	"vormstein.eu/gocitygame/app"
)

func init() {
	parseOptions := &atempgo.ParseOptions{Ext: "tmpl"}

	atempgo.LoadTemplates("app/views/"+app.Config.Tmpl, parseOptions)
}

func GetTemplate(key string) *template.Template {
	return atempgo.GetTemplate(key)
}

func NotAllowed(w http.ResponseWriter) {
	GetTemplate("#error").ExecuteTemplate(w, "base", app.GetHTTPError(405))
}
