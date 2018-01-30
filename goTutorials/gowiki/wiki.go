// https://golang.org/doc/articles/wiki/
package main

import (
	"html/template"
	"io/ioutil"
	"net/http"
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

/*  You can compile and run the program like this:

$ go build wiki.go
$ ./wiki
*/

func main() {
	view := "/view/"
	edit := "/edit/"
	http.HandleFunc(view, Path{Name: view}.handler) //can serve localhost:8080/view/TestPage.txt
	http.HandleFunc(edit, Path{Name: edit}.handler)
	http.HandleFunc(save, saveHandler)

	http.ListenAndServe(":8080", nil)
	// p1 := &Page{Title: "TestPage", Body: []byte("This is a sample Page.")}
	// p1.save()
	// p2, _ := loadPage("TestPage")
	// fmt.Println(string(p2.Body))
}

func (path Path) handler(w http.ResponseWriter, r *http.Request) {
	title := r.URL.Path[len(path.Name):]
	p, err := loadPage(title)

	switch err != nil {
	case path.Name[:6] == "/view/":
		http.Redirect(w, r, "/edit/"+title, http.StatusFound)
		return
	default:
		p = &Page{Title: title}
	}

	renderTemplate(w, path.Name[1:len(path.Name)-1], p)
}

func loadPage(title string) (*Page, error) {
	filename := title + ".txt"
	body, err := ioutil.ReadFile(filename) // returns byte, error

	if err != nil {
		return nil, err
	}
	return &Page{Title: title, Body: body}, nil
	// return a pointer to the newly constructed Page
}

// Callers of this function can now check the second parameter;
// if it is nil then it has successfully loaded a Page.
// If not, it will be an error that can be handled by the caller

func renderTemplate(w http.ResponseWriter, t string, p *Page) {
	tt, _ := template.ParseFiles(t + ".html") //returns a *template.Template
	//can also pass an array of string -^
	/* When parsing multiple files with the same name in different directories,
	the last one mentioned will be the one that results. */
	tt.Execute(w, p)
	// func (t *Template) Execute(wr io.Writer, data interface{}) error
	// applies a parsed template to the specified data object,
	// writing the output to wr.
	// If an error occurs executing the template or writing its output, execution stops,
	// but partial results may already have been written to the output writer.
	// ...if parallel executions share a Writer the output may be interleaved.

	// fmt.Fprintf(w, `
	// 	html`,
	// 	p.Title, p.Title, p.Body)
	/* func Fprintf(w io.Writer, format string, a ...interface{}) (n int, err error)
	Fprintf formats according to a format specifier and writes to w.
	It returns the number of bytes written and any write error encountered. */
}

func saveHandler(w http.ResponseWriter, r *http.Request) {
	title := r.URL.Path[len("/save/"):]
	body := r.FormValue("body") //string
	p := &Page{Title: title, Body: []byte(body)}
	p.save()
	http.Redirect(w, r, "/view/"+title, http.StatusFound)
}

// This method will save the Page's Body to a text file.
// For simplicity, we will use the Title as the file name.
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
