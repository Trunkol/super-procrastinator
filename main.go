package main

import (
	"super-procrastinator/reddit"
)

func main() {
	reddit.TopStories()
}

/*
import (
	"html/template"
	"net/http"
	"super-procrastinator/consumer"
)

func main() {
	tmpl := template.Must(template.ParseFiles("templates/layout.html"))

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		tmpl.Execute(w, consumer.Stories())
	})

	http.ListenAndServe(":8000", nil)
}
*/
