package main

import (
	"fmt"
	"html"
	"html/template"
	"log"
	"net/http"
)

func httpServer() {

	http.HandleFunc("/bar", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello, %q", html.EscapeString(r.URL.Path))
	})

	log.Fatal(http.ListenAndServe(":8080", nil))
}

func staticServer() {

	log.Fatal(http.ListenAndServe(":8080", http.FileServer(http.Dir("D:/code/wing-xue/golang_study/study/html"))))
}

func templateServer() {
	// Define a template.
	const letter = `
		Dear {{.Name}},
		{{if .Attended}}
		It was a pleasure to see you at the wedding.
		{{- else}}
		It is a shame you couldn't make it to the wedding.
		{{- end}}
		{{with .Gift -}}
		Thank you for the lovely {{.}}.
		{{end}}
		Best wishes,
		Josie
		`

	// Prepare some data to insert into the template.
	type Recipient struct {
		Name, Gift string
		Attended   bool
	}
	var recipients = Recipient{"Aunt Mildred", "bone china tea set", true}

	// Create a new template and parse the letter into it.
	t := template.Must(template.New("letter").Parse(letter))

	// Execute the template for each recipient.
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		err := t.Execute(w, recipients)
		if err != nil {
			log.Println("executing template:", err)
		}

	})

	log.Fatal(http.ListenAndServe(":8080", nil))

}
func main() {
	templateServer()
}
