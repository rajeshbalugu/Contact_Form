package main

import (
	"fmt"
	"html/template"
	"net/http"
)

func main() {

	// Serve the HTML file
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		tmpl := template.Must(template.ParseFiles("contact.html"))
		tmpl.Execute(w, nil)
	})

	// Handle form submission
	http.HandleFunc("/submit", func(w http.ResponseWriter, r *http.Request) {
		data := map[string]string{
			"Name":  r.FormValue("name"),
			"Email": r.FormValue("email"),
			"Msg":   r.FormValue("msg"),
		}
		tmpl := template.Must(template.ParseFiles("result.html"))
		tmpl.Execute(w, data)
	})

	fmt.Println("Server running at http://localhost:8080")
	http.ListenAndServe(":8080", nil)
}
