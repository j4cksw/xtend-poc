package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"

	"github.com/mitchellh/mapstructure"
	"golang.org/x/net/websocket"
)

var users []Player

type Event struct {
	Action string                 `json:"action"`
	Data   map[string]interface{} `json:"data,omitempty"`
}

type Player struct {
	Name string `json:"name"`
	Conn *websocket.Conn
}

func Start(ws *websocket.Conn) {
	var event Event
	//for websocket.JSON.Receive(ws, &player) != io.EOF {
	for websocket.JSON.Receive(ws, &event) != io.EOF {

		switch event.Action {
		case "new":
			var player Player
			err := mapstructure.Decode(event.Data, &player)
			if err != nil {

			}
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
			} else {
				websocket.JSON.Send(player.Conn, Event{Action: "init"})
			}

		case "":
		}

	}
}

func main() {
	http.Handle("/api/start", websocket.Handler(Start))
	http.Handle("/", http.FileServer(http.Dir(os.Getenv("STATIC_PATH"))))

	port := os.Getenv("PORT")
	if port == "" {
		port = "12345"
	}
	log.Println("Listen on port", port)
	err := http.ListenAndServe(":"+port, nil)
	if err != nil {
		panic("ListenAndServe: " + err.Error())
	}
}
