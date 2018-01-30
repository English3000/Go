// https://golang.org/doc/articles/wiki/
package main

import (
	"html/template"
	"io/ioutil"
	"net/http"
	"regexp"
)

// Path ...of an HTTP request
type Path struct {
	Name string
}

// Page ...a wiki consists of many, accessed via a Path
type Page struct {
	Title string
	Body  []byte
	// The Body element is a []byte rather than string because
	// that is the type expected by the io libraries we will use
}

// When parsing multiple files with the same name in different directories,
// the last one mentioned will be the one that results.
var templates = template.Must(template.ParseFiles("edit.html", "view.html"))

// MustCompile is distinct from Compile in that
// it will panic if the expression compilation fails,
// while Compile returns an error as a second parameter.
var validPath = regexp.MustCompile("^/(edit|save|view)/([a-zA-Z0-9]+)$")

/*  You can compile and run the program like this:

$ go build wiki.go
$ ./wiki
*/
func main() {
	view := "/view/"
	edit := "/edit/"
	save := "/save/"

	http.HandleFunc(view, makeHandler(Path{Name: view}.handler))
	http.HandleFunc(edit, makeHandler(Path{Name: edit}.handler))
	http.HandleFunc(save, makeHandler(saveHandler))

	http.ListenAndServe(":8080", nil)
}

func makeHandler(fn func(http.ResponseWriter, *http.Request, string)) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		p := validPath.FindStringSubmatch(r.URL.Path)
		if p == nil {
			http.NotFound(w, r)
			return
		}
		fn(w, r, p[2])
	}
}

func (path Path) handler(w http.ResponseWriter, r *http.Request, title string) {
	p, err := loadPage(title)
	switch err != nil {
	case path.Name[:5] == "/view":
		http.Redirect(w, r, "/edit/"+title, http.StatusFound) //redirect error for edit/ExistingPage
		return
	default:
		p = &Page{Title: title}
	}

	renderTemplate(w, path.Name[1:len(path.Name)-1], p)
}

func loadPage(title string) (*Page, error) {
	filename := title + ".txt"
	body, err := ioutil.ReadFile(filename) //returns byte, error

	if err != nil {
		return nil, err
	}
	return &Page{Title: title, Body: body}, nil
	// return a pointer to the newly constructed Page
}

func renderTemplate(w http.ResponseWriter, t string, p *Page) {
	err := templates.ExecuteTemplate(w, t+".html", p)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func saveHandler(w http.ResponseWriter, r *http.Request, title string) {
	body := r.FormValue("body") //string
	p := &Page{Title: title, Body: []byte(body)}

	err := p.save()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	http.Redirect(w, r, "/view/"+title, http.StatusFound)
}

func (p *Page) save() error {
	filename := p.Title + ".txt"

	// The save method returns the error value,
	// to let the application handle it
	// should anything go wrong while writing the file.
	return ioutil.WriteFile(filename, p.Body, 0600)
	// The octal integer literal 0600,
	// passed as the third parameter to WriteFile,
	// indicates that the file should be created with
	// read-write permissions for the current user only.
}

// err = tt.Execute(w, p)
// func (t *Template) Execute(wr io.Writer, data interface{}) error
// applies a parsed template to the specified data object,
// writing the output to wr.
// if err != nil {
// 	http.Error(w, err.Error(), http.StatusInternalServerError)
// }
// If an error occurs executing the template or writing its output, execution stops,
// but partial results may already have been written to the output writer.
// ...if parallel executions share a Writer the output may be interleaved.

// fmt.Fprintf(w, `
// 	html`,
// 	p.Title, p.Title, p.Body)
/* func Fprintf(w io.Writer, format string, a ...interface{}) (n int, err error)
Fprintf formats according to a format specifier and writes to w.
It returns the number of bytes written and any write error encountered. */
