package main

import (
	"html/template"
	"io/ioutil"
	"log"
	"net/http"
	"regexp"
)

type Page struct {
	Title string
	Body  []byte
}

func (p *Page) save() error {
	filename := p.Title + ".txt"
	return ioutil.WriteFile(filename, p.Body, 0600)
}

func loadPage(title string) (*Page, error) {
	filename := title + ".txt"
	body, err := ioutil.ReadFile(filename)
	if err != nil {
		return nil, err
	}
	return &Page{Title: title, Body: body}, nil
}

func viewHandler(w http.ResponseWriter, r *http.Request, title string) {
	p, err := loadPage(title)
	if err != nil {
		http.Redirect(w, r, "/app4/"+title, http.StatusFound)
		return
	}
	renderTemplate(w, "view", p)
}

func app4Handler(w http.ResponseWriter, r *http.Request, title string) {
	p, err := loadPage(title)
	if err != nil {
		p = &Page{Title: title}
	}
	renderTemplate(w, "app4", p)
}

var templates = template.Must(template.ParseFiles("app4.html", "view.html"))

func renderTemplate(w http.ResponseWriter, tmpl string, p *Page) {
	err := templates.ExecuteTemplate(w, tmpl+".html", p)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

// Compile regular expressions for path validation.
var validPath = regexp.MustCompile("^/(app4|view)/([a-zA-Z0-9]+)$")

func makeHandler(fn func(http.ResponseWriter, *http.Request, string)) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// If the title is invalid, an error will be written
		// to the ResponseWriter using the http.NotFound function.
		// If the title is valid, the enclosed handler function
		// fn will be called with the ResponseWriter, Request, and title as arguments.
		m := validPath.FindStringSubmatch(r.URL.Path)
		if m == nil {
			http.NotFound(w, r)
			return
		}
		fn(w, r, m[2])
	}
}

func main() {
	//fs := justFilesFilesystem{http.Dir("GO_Excercise/web_apps/static")}
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))
	http.HandleFunc("/view/", makeHandler(viewHandler))
	http.HandleFunc("/app4/", makeHandler(app4Handler))
	//http.HandleFunc("/", makeHandler())

	log.Fatal(http.ListenAndServe(":5000", nil))
}
