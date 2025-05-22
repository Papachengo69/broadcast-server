package client

import (
	"bufio"
	"fmt"
	"log"
	"os"

	"github.com/gorilla/websocket"
)

func ConnectToServer() {
	conn, _, err := websocket.DefaultDialer.Dial("ws://localhost:8080/ws", nil)
	if err != nil {
		log.Fatal("Dial error:", err)
	}

	defer conn.Close()

	go func() {
		for {
			_, msg, err := conn.ReadMessage()
			if err != nil {
				log.Println("Read error:", err)
				return
			}

			fmt.Println("Received:", string(msg))
		}
	}()

	reader := bufio.NewReader(os.Stdin)

	for {
		fmt.Print("Enter Message: ")
		msg, _ := reader.ReadString('\n')
		err := conn.WriteMessage(websocket.TextMessage, []byte(msg))

		if err != nil {
			log.Println("Write error: ", err)
			break
		}
	}
}
