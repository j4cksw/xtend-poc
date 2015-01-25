package main

import (
	"fmt"
	"io"
	"log"
	"math"
	"math/rand"
	"net/http"
	"os"
	"time"

	"github.com/mitchellh/mapstructure"
	"golang.org/x/net/websocket"
)

var players []Player

type Event struct {
	Action string                 `json:"action"`
	Data   map[string]interface{} `json:"data,omitempty"`
}

type Minion struct {
	X     float64 `json:"x"`
	Y     float64 `json:"y"`
	Color string  `json:"color"`
}

type Player struct {
	Name    string          `json:"name"`
	Conn    *websocket.Conn `json:"-"`
	X       int             `json:"x"`
	Y       int             `json:"y"`
	Color   string          `json:"color"`
	Minions []Minion
}

func broadcast(evt Event, players []Player) {
	for _, p := range players {
		if err := websocket.JSON.Send(p.Conn, evt); err != nil {
			log.Fatal(err)
		}
	}
}

func Start(ws *websocket.Conn) {
	var event Event
	for websocket.JSON.Receive(ws, &event) != io.EOF {

		switch event.Action {
		case "new":
			var player Player
			err := mapstructure.Decode(event.Data, &player)
			if err != nil {
				log.Fatal(err)
			}
			player.Conn = ws
			players = append(players, player)
			fmt.Println(player.Name, " logged in")

			if len(players) >= 2 {
				fmt.Println("Total players is ", len(players), " game start")

				players[0].X = 100
				players[0].Y = 300
				players[0].Color = "0xFFFF0B"
				players[1].X = 700
				players[1].Y = 300
				players[1].Color = "0xBC1C22"

				event := Event{
					Action: "render_base",
					Data: map[string]interface{}{
						"players": players,
					},
				}

				broadcast(event, players)
			} else {
				websocket.JSON.Send(player.Conn, Event{Action: "init"})
			}
		case "request_minion":
			var player Player
			mapstructure.Decode(event.Data, &player)

			fmt.Println(player)

			for i, p := range players {
				if player.Name != p.Name {
					continue
				}

				var minion = Minion{X: RandomPositionX(p.X, 60), Y: RandomPositionY(p.Y, 60), Color: p.Color}
				p.Minions = append(p.Minions, minion)

				players[i] = p // set it back because it pass by value
				fmt.Println(players)
			}

			var renderMinionEvt = Event{
				Action: "render_minion",
				Data: map[string]interface{}{
					"minions": append(players[0].Minions, players[1].Minions...),
				},
			}
			broadcast(renderMinionEvt, players)
		}
	}
}

func RandomPositionX(x, radius int) float64 {
	// random degree [0, 360]
	var a = rand.New(rand.NewSource(time.Now().UTC().UnixNano())).Intn(361)
	return float64(x) + float64(radius) + math.Cos(float64(a))
}

func RandomPositionY(y, radius int) float64 {
	// random degree [0, 360]
	var a = rand.New(rand.NewSource(time.Now().UTC().UnixNano())).Intn(361)
	return float64(y) + float64(radius) + math.Sin(float64(a))
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
