package app

import (
	"log"
	"net/http"
	"text/template"
)

type app struct{}

func New() *app {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		tmpl := template.Must(template.ParseFiles("./internal/app/views/index.html"))
		tmpl.Execute(w, nil)
	})

	return &app{}
}

func (*app) Run() {
	log.Println("App Running on 8000...")
	log.Fatalln(http.ListenAndServe(":8000", nil))
}
