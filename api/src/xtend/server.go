package main

import (
	"fmt"
	"golang.org/x/net/websocket"
	"io"
	"net/http"
)

var users []Player

type Player struct {
	Name string `json:"name"`
	Conn *websocket.Conn
}

func Login(ws *websocket.Conn) {
	var player Player
	for websocket.JSON.Receive(ws, &player) != io.EOF {
		player.Conn = ws
		users = append(users, player)

		if (len(users)) == 2 {
			websocket.JSON.Send(users[0].Conn, users[1].Name)
			websocket.JSON.Send(users[1].Conn, users[0].Name)
		}

	}
	fmt.Println(player)
}

func main() {
	http.Handle("/api/start", websocket.Handler(Login))
	err := http.ListenAndServe(":12345", nil)
	if err != nil {
		panic("ListenAndServe: " + err.Error())
	}
}
