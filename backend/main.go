package main

import (
    "fmt"
    "net/http"

	"github.com/harshvardhan-pandey/chatterbox/pkg/websocket"
)

func serveWs(w http.ResponseWriter, r *http.Request){
	
	conn, err := websocket.Upgrade(w, r)
	if err != nil{
		fmt.Fprintf(w, "%+V\n", err)
		return
	}

	go websocket.Writer(conn)
	websocket.Reader(conn)

}

func setupRoutes() {
	http.HandleFunc("/ws", serveWs)
}

func main() {
    setupRoutes()
    if err := http.ListenAndServe(":8080", nil); err != nil{
		fmt.Println(err)
	}
}