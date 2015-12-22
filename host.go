package main

import (
	"fmt"
	"github.com/gorilla/websocket"
	//"net/http"
)

var hostList map[string]*host

type host struct {
	conn *websocket.Conn
	id   string
}

func initHost() {
}

func hostHandler(conn *websocket.Conn) {
	var identifier string

	for {
		messageType, p, err := conn.ReadMessage()
		if err != nil {
			return
		}
		fmt.Println(string(p)[1:])
		if p[0] == '_' {
			if _, ok := hostList[string(p)[1:]]; !ok {
				hostList[string(p)[1:]] = &host{conn, string(p)[1:]}
				identifier = string(p)[1:]
				conn.WriteMessage(messageType, p[1:])
			} else {
				conn.WriteMessage(messageType, []byte(identifier))
			}
		}
	}

}
