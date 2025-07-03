package main

import (
    "fmt"
    "net/http"

	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{
	ReadBufferSize: 1024,
	WriteBufferSize: 1024,
	CheckOrigin: func(r *http.Request) bool { return true }, //allow everyone right now
}

func serveWs(w http.ResponseWriter, r *http.Request){
	
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil{
		fmt.Println(err)
		return
	}
	
	defer conn.Close()

	for{
		messageType, data, err := conn.ReadMessage()
		if err != nil{
			fmt.Println(err)
			return
		}
		fmt.Println(string(data))

		if err:=conn.WriteMessage(messageType, data); err != nil{
			fmt.Println(err)
			return
		}

	}


}

func setupRoutes() {
    http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
        fmt.Fprintf(w, "Simple Server")
    })
	http.HandleFunc("/ws", serveWs)
}

func main() {
    setupRoutes()
    if err := http.ListenAndServe(":8080", nil); err != nil{
		fmt.Println(err)
	}
}