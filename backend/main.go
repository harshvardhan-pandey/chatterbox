package main

import (
    "fmt"
    "net/http"

	"github.com/harshvardhan-pandey/chatterbox/pkg/websocket"
)

func serveWs(pool *websocket.Pool, w http.ResponseWriter, r *http.Request){
	
	conn, err := websocket.Upgrade(w, r)
	if err != nil{
		fmt.Fprintf(w, "%+v\n", err)
	}

	client := &websocket.Client{
        Conn: conn,
        Pool: pool,
    }

    pool.Register <- client
    client.Read()

}

func setupRoutes() {
	pool := websocket.NewPool()
	go pool.Start()
	http.HandleFunc("/ws", func(w http.ResponseWriter, r *http.Request) {
        serveWs(pool, w, r)
    })
}

func main() {
    setupRoutes()
    if err := http.ListenAndServe(":8080", nil); err != nil{
		fmt.Println(err)
	}
}