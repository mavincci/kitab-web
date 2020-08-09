package api

import "net/http"
import "github.com/mavincci/Kitab-web/templates"

func init() {
}

func Index(w http.ResponseWriter, r *http.Request) {
	templates.Index.Execute(w, nil)
}
