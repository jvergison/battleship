package main

import (
	"flag"
	"fmt"
	"html/template"
	"log"
	"net/http"
)

var addr = flag.String("addr", "localhost:8585", "http service address")
var homeTemplate = template.Must((template.ParseFiles("./static/index.html")))

func main() {
	fmt.Println("Start")

	flag.Parse()
	log.SetFlags(0)

	initMessageHandler()

	router := newRouter()
	log.Fatal(http.ListenAndServe(*addr, router))
}

func home(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		http.Error(w, "Not found", 404)
		return
	}
	if r.Method != "GET" {
		http.Error(w, "Method not allowed", 405)
		return
	}
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	w.WriteHeader(http.StatusOK)
	homeTemplate.Execute(w, r.Host)
}

func battleship(w http.ResponseWriter, r *http.Request) {
	makeConnection(w, r)

}
