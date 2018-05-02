package main

import (
	"fmt"
	"super-procrastinator/consumer"
)

func main() {

	for _, v := range consumer.Stories() {
		fmt.Println(v.URL)
	}

	/*
		tmpl := template.Must(template.ParseFiles("templates/layout.html"))

		http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
			tmpl.Execute(w, data)
		})

		http.ListenAndServe(":8000", nil)
	*/
}
