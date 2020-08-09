package templates

import (
	htmlTmpl "html/template"
)

var Index *htmlTmpl.Template

func init() {
	Index = htmlTmpl.Must(htmlTmpl.ParseFiles("templates/index.html"))
}
