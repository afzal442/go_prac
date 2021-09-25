package main

import (
	"html/template"
	"io/ioutil"
	"log"
	"net/http"
)

type Page struct {
	Title string // title of the page
	Body  []byte // body of the page
}

func (p *Page) save() error {
	filename := p.Title + ".txt"
	return ioutil.WriteFile(filename, p.Body, 0600) // write the file
}

// loadPage returns an error if title is not found.
// It returns the existing Page if title is found.
func loadPage(title string) (*Page, error) {
	filename := title + ".txt"             // "title.txt"
	body, err := ioutil.ReadFile(filename) // read the file
	if err != nil {
		return nil, err
	}
	return &Page{Title: title, Body: body}, nil // return a new Page
}

// renderTemplate renders the specified template with the specified
// page as the data.
func renderTemplate(w http.ResponseWriter, tmpl string, p *Page) {
	t, _ := template.ParseFiles(tmpl + ".html")
	t.Execute(w, p)
}

// An http.ResponseWriter value assembles the HTTP server's response;
// by writing to it, we send data to the HTTP client.
// A ResponseWriter may be used to send different types of data
// (e.g., plain text, HTML, JSON) in an HTTP response.
// A http.Request value contains the client's request.
func viewHandler(w http.ResponseWriter, r *http.Request) {
	title := r.URL.Path[len("/view/"):] // "/view/title"
	p, error := loadPage(title)         // loadPage returns error if title is not found
	if error != nil {
		http.Redirect(w, r, "/edit/"+title, http.StatusFound)
		return
	}
	renderTemplate(w, "view", p) // render the view template
}

func saveHandler(w http.ResponseWriter, r *http.Request) {
	title := r.URL.Path[len("/save/"):]                   // "/save/title"
	body := r.FormValue("body")                           // "body" is the name of the form field, formValue returns the value of the form field
	p := &Page{Title: title, Body: []byte(body)}          // create a new Page
	p.save()                                              // save the page
	http.Redirect(w, r, "/view/"+title, http.StatusFound) // redirect to the view page
}

func editHandler(w http.ResponseWriter, r *http.Request) {
	title := r.URL.Path[len("/edit/"):] // "/edit/title"
	p, err := loadPage(title)           // loadPage returns error if title is not found
	if err != nil {
		p = &Page{Title: title} // create a new Page
	}
	renderTemplate(w, "edit", p) // render the edit template
}

func main() {
	http.HandleFunc("/view/", viewHandler)
	http.HandleFunc("/edit/", editHandler)
	http.HandleFunc("/save/", saveHandler)
	log.Fatal(http.ListenAndServe(":3000", nil))
}
