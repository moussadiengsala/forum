package router

import (
	"fmt"
	"golang-rest-api-starter/handlers"
	"log"
	"net/http"
	"strings"

	"github.com/gorilla/websocket"
)

type Middleware func(http.HandlerFunc) http.HandlerFunc
type NewRouter struct {
	Middlewares []Middleware
}

var clients = make(map[*websocket.Conn]bool)
var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
}

func broadcastMessage(message []byte) {
	for client := range clients {
		err := client.WriteMessage(websocket.TextMessage, message)
		if err != nil {
			log.Println("Error broadcasting message:", err)
			client.Close()
			delete(clients, client) // Remove the disconnected client from the map
		}
	}
}

func ws(w http.ResponseWriter, r *http.Request) {
	conn, _ := upgrader.Upgrade(w, r, nil) // error ignored for sake of simplicity
	clients[conn] = true
	for {
		// Read message from browser
		_, msg, err := conn.ReadMessage()
		if err != nil {
			return
		}

		// Print the message to the console
		fmt.Printf("%s sent: %s\n", conn.RemoteAddr(), string(msg))

		// Write message back to browser
		// if err = conn.WriteMessage(msgType, msg); err != nil {
		// 	return
		// }
		broadcastMessage(msg)
	}
}

func (router *NewRouter) ServeHTTP(w http.ResponseWriter, r *http.Request) {

	handler := func(w http.ResponseWriter, r *http.Request) {
		switch {
		// Home Page
		case r.URL.Path == "/":
			handlers.HomeHandler(w, r)
		case r.URL.Path == "/echo":
			ws(w, r)
		// Static files
		case strings.HasPrefix(r.URL.Path, "/static/"):
			fs := http.StripPrefix("/static/", http.FileServer(http.Dir("./static")))
			fs.ServeHTTP(w, r)

		// When the URL doesn't exist
		default:
		}
	}

	for _, mw := range router.Middlewares {
		handler = mw(handler)
	}
	handler(w, r)
}
