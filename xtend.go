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
	Name  string          `json:"name"`
	Conn  *websocket.Conn `json:"-"`
	X     int             `json:"x"`
	Y     int             `json:"y"`
	Color string          `json:"color"`
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

				users[0].X = 100
				users[0].Y = 300
				users[0].Color = "0xFFFF0B"
				users[1].X = 700
				users[1].Y = 300
				users[1].Color = "0xBC1C22"

				firstuser := Event{
					Action: "render_base",
					Data: map[string]interface{}{
						"players": users,
					},
				}

				if err := websocket.JSON.Send(users[0].Conn, firstuser); err != nil {
					fmt.Println(err.Error())
				}
				if err := websocket.JSON.Send(users[1].Conn, firstuser); err != nil {
					fmt.Println(err.Error())
				}
			} else {
				websocket.JSON.Send(player.Conn, Event{Action: "init"})
			}

		case "":
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
