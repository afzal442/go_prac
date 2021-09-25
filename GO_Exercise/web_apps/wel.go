package main

import (
	"html/template"
	"net/http"
	"time"
)

type app struct {
	title string
	Time  string
}

func main() {
	ax := app{
		"My First Web App",
		time.Now().String(),
	}
	t := template.Must(template.ParseFiles("wel.html"))
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		if s := r.FormValue("title"); s != "" {
			ax.title = s
		}
		if err := t.Execute(w, ax); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}

	})
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))
	http.ListenAndServe(":8080", nil)
}
