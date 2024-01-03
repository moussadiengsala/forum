package ws

import (
	"fmt"
	"net/http"

	"github.com/gorilla/websocket"
)

// Here, an Upgrader is created with specified read and write buffer sizes.
var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
}

func Ws(w http.ResponseWriter, r *http.Request) {

	// return true, allowing connections from any origin.
	upgrader.CheckOrigin = func(r *http.Request) bool {
		return true
	}

	conn, err := upgrader.Upgrade(w, r, nil) // error ignored for sake of simplicity

	if err != nil {
		fmt.Println(err)
		return
	}

	defer conn.Close()

	for {
		// Read message from browser
		msgType, msg, err := conn.ReadMessage()
		if err != nil {
			return
		}

		// Print the message to the console
		fmt.Printf("%s sent: %s\n", conn.RemoteAddr(), string(msg))

		// Write message back to browser
		if err = conn.WriteMessage(msgType, msg); err != nil {
			return
		}
	}
}
