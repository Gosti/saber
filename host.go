package main

import (
	"fmt"
	"github.com/gorilla/websocket"
	"io/ioutil"
	"net/http"
)

type host struct {
	conn    *websocket.Conn
	id      string
	players []player
}

func initHost() {
}

func connHandler(w http.ResponseWriter, r *http.Request, string h) {

}
