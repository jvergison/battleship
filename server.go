package main

import (
	"flag"
	"fmt"
	"html/template"
	"log"
	"net/http"

	"github.com/gorilla/websocket"
)

var addr = flag.String("addr", "localhost:8585", "http service address")
var homeTemplate = template.Must((template.ParseFiles("index.html")))
var connCount = 1
var upgrader = websocket.Upgrader{
//default options
}

type Connection struct {
	socket *websocket.Conn
	id     int
}

var conns []Connection

func main() {
	fmt.Println("Start")

	flag.Parse()
	log.SetFlags(0)
	http.HandleFunc("/battleshipServer", battleship)
	http.HandleFunc("/", home)
	log.Fatal(http.ListenAndServe(*addr, nil))
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
	homeTemplate.Execute(w, r.Host)
}

func battleship(w http.ResponseWriter, r *http.Request) {
	fmt.Println("new connection")
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println(err)
		return
	}

	defer conn.Close() //clean up if we ever exit this function

	//TODO: add connection to connection list
	var c Connection = Connection{id: connCount, socket: conn}
	conns = append(conns, c)
	connCount = connCount + 1

	for {
		mt, message, err := conn.ReadMessage() //keep listening
		if err != nil {
			log.Println("read:", err)
			break
		}
		log.Printf("recv: %s", message)
		err = conn.WriteMessage(mt, message)
		if err != nil {
			log.Println("write:", err)
			break
		}
	}

}
