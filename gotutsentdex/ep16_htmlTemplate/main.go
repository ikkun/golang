package main

import (
	"fmt"
	"html/template"
	"net/http"
)

type NewsAggPage struct {
	Title string
	News  map[string]NewsMap
}

func newsAggHandler(w http.ResponseWriter, r *http.Request) {
	p := NewsAggPage{Title: "Amazing News Aggregator", News: "some news"}
	t, _ := template.ParseFiles("basictemplating.html")
	fmt.Println(t.Execute(w, p))

}

func index_handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "<h1>Hello world.</h1>")

}

func main() {
	http.HandleFunc("/", index_handler)
	http.HandleFunc("/agg/", newsAggHandler)
	http.ListenAndServe(":9000", nil)
}
