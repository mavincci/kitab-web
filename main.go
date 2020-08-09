package main

import "net/http"
import "github.com/mavincci/Kitab-web/api"

func main() {
	fs := http.FileServer(http.Dir("./assets"))
	http.Handle("/assets/", http.StripPrefix("/assets/", fs))
	http.HandleFunc("/", api.Index)
	_ = http.ListenAndServe(":90", nil)
}
