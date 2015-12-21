package main

import (
	"fmt"
	"github.com/gorilla/websocket"
	"net/http"
)

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
}

var desktop *websocket.Conn
var phone *websocket.Conn

func print_binary(s []byte) {
	fmt.Print(string(s), "\n")
}

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
			desktop = conn
			k = 2
		} else if string(p) == "PHONE" {
			fmt.Println("Phone connected")
			phone = conn
			k = 1
		}

		//print_binary(p)
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
	http.HandleFunc("/echo", echoHandler)
	http.Handle("/", http.FileServer(http.Dir("./")))

	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		panic("Error: " + err.Error())
	}
}
