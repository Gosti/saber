package main

import (
	"fmt"
	"github.com/gorilla/websocket"
	"html/template"

	"net/http"
	"os"
)

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
}

type Page struct {
	IP string
	ID string
}

var desktop *websocket.Conn
var phone *websocket.Conn

func echoHandler(w http.ResponseWriter, r *http.Request) {
	var k byte = 0
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		return
	}

	for {
		messageType, p, err := conn.ReadMessage()
		if err != nil {
			return
		}

		if string(p) == "DESKTOP" {
			fmt.Println("Desktop connected")
			hostHandler(conn)
			return
		} else if string(p) == "PHONE" {
			fmt.Println("Phone connected")
			phone = conn
			k = 1
		}

		if k == 1 {
			desktop.WriteMessage(messageType, p)
		}
		err = conn.WriteMessage(messageType, p)
		if err != nil {
			fmt.Println(err)
			return
		}
	}
}

func main() {

	hostList = make(map[string]*host)

	http.HandleFunc("/echo", echoHandler)
	// host, err := ioutil.ReadFile("./host/index.html")
	// if err != nil {
	// 	panic(err)
	// }
	// phone, err := ioutil.ReadFile("./phone/index.html")
	// if err != nil {
	// 	panic(err)
	// }

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path[1:] == "favicon.ico" {
			return
		}
		fmt.Println(r.URL.Path[1:])
		if len(r.URL.Path[1:]) > 0 {
			t, _ := template.ParseFiles("./host/index.html")
			t.Execute(w, &Page{IP: os.Args[1], ID: ""})
		} else {

			t, _ := template.ParseFiles("./host/index.html")
			t.Execute(w, &Page{IP: os.Args[1]})
		}
	})

	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		panic("Error: " + err.Error())
	}
}
