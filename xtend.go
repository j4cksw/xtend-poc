package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"

	"golang.org/x/net/websocket"
)

var users []Player

type Player struct {
	Name string `json:"name"`
	Conn *websocket.Conn
}

type Message struct {
	Sender string `json:"sender"`
	Text   string `json:"text"`
}

func Start(ws *websocket.Conn) {
	var player Player
	for websocket.JSON.Receive(ws, &player) != io.EOF {
		player.Conn = ws
		users = append(users, player)
		fmt.Println(player.Name, " logged in")

		if (len(users)) >= 2 {
			fmt.Println("Total users is ", len(users), " game start")
			if err := websocket.JSON.Send(users[0].Conn, users[1].Name); err != nil {
				fmt.Println(err.Error())
			}
			if err := websocket.JSON.Send(users[1].Conn, users[0].Name); err != nil {
				fmt.Println(err.Error())
			}
		}

	}
}

func Boardcast(ws *websocket.Conn) {
	var msg Message
	for websocket.JSON.Receive(ws, &msg) != io.EOF {
		for i := range users {
			if users[i].Name != msg.Sender {
				websocket.JSON.Send(users[i].Conn, msg.Text)
				fmt.Printf("%s Boardcast %s ======> %s\n", msg.Sender, msg.Text, users[i].Name)
			}
		}

	}
}

func optionalEnv(key, defaultValue string) string {
	var v = os.Getenv(key)
	if v == "" {
		return defaultValue
	}
	return v
}

func main() {
	http.Handle("/api/start", websocket.Handler(Start))

	staticPath := optionalEnv("STATIC_PATH", "web")
	http.Handle("/", http.FileServer(http.Dir(staticPath)))

	port := optionalEnv("PORT", "12345")
	log.Println("Listen on port", port)

	err := http.ListenAndServe(":"+port, nil)
	if err != nil {
		panic("ListenAndServe: " + err.Error())
	}
}
