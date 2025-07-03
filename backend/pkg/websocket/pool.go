package websocket

import (
	"fmt"
	"log"
)

type Pool struct {
    Register chan *Client
    Unregister chan *Client
    Clients map[*Client]bool
    Broadcast chan Message
}

func NewPool() *Pool {
    return &Pool{
        Register:   make(chan *Client),
        Unregister: make(chan *Client),
        Clients:    make(map[*Client]bool),
        Broadcast:  make(chan Message),
    }
}

func (pool *Pool) Start(){

	for{
		select{

		case client := <-pool.Register:
			pool.Clients[client] = true
			fmt.Printf("Size of Connection = %d\n", len(pool.Clients))
			for client := range pool.Clients{
				client.Conn.WriteJSON(Message{Type:1, Body: "New User Joined"})
			}

		case client := <- pool.Unregister:
			delete(pool.Clients, client)
			fmt.Printf("Size of Connection = %d\n", len(pool.Clients))
			for client := range pool.Clients{
				client.Conn.WriteJSON(Message{Type:1, Body: "User Disconnected"})
			}

		case message := <- pool.Broadcast:
			fmt.Printf("Sending messaged to all clients")
			for client := range pool.Clients{
				if err:= client.Conn.WriteJSON(message); err != nil{
					log.Println(err)
					return
				}
			}

		}
	}

}