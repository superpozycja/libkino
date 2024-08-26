package main

import (
	"fmt"
	"log"
	"net/http"
	"html/template"
)

type Author struct {
	Name string
	Link string
}

type Entry struct {
	Title string
	Authors []Author
	Link string
	Description []string
}

var Entries []Entry

func init() {
	Entries = []Entry{
		{Title: "The C Programming Language",
		Authors: []Author {
			{"Brian Kernighan", "google.com"},
			{"Dennis Ritchie", "google.com"},
		},
		Link: "google.com",
		Description: []string {"Cult classic book for learning C. (i wouldn't recommend it for beginners though - anna)",
			     "\"The authors present the complete guide to ANSI standard C language programming. Written by the developers of C, this new version helps readers keep up with the finalized ANSI standard for C while showing how to take advantage of C's rich set of operators, economy of expression, improved control flow, and data structures. The 2/E has been completely rewritten with additional examples and problem sets to clarify the implementation of difficult language constructs. For years, C programmers have let K&R guide them to building well-structured and efficient programs. Now this same help is available to those working with ANSI compilers. Includes detailed coverage of the C language plus the official C language reference manual for at-a-glance help with syntax notation, declarations, ANSI changes, scope rules, and the list goes on and on.\""},
		     },

		{Title: "How Linux Works",
		Authors: []Author {
			{"Brian Ward", "google.com"},
		},
		Link: "google.com",
		Description: []string {"some linux book"}},
	}
}

func mainHandler (w http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(template.ParseFiles("index.html"))
	tmpl.Execute(w, Entries)
	/*
	for _, i := range(Entries["Entries"]) {
		tmpl.ExecuteTemplate(w, "main.html", i)
	}
	*/
}

func aboutHandler (w http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(template.ParseFiles("about.html"))
	tmpl.Execute(w, nil)
}

func submissionsHandler (w http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(template.ParseFiles("submissions.html"))
	tmpl.Execute(w, nil)
}

func main() {
	fmt.Println("boo")

	http.HandleFunc("/", mainHandler)
	http.HandleFunc("/about", aboutHandler)
	http.HandleFunc("/submissions", submissionsHandler)
	http.Handle("/css/", http.StripPrefix("/css", http.FileServer(http.Dir("css"))))
	log.Fatal(http.ListenAndServe(":8000", nil))
}
