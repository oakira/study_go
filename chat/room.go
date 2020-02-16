package main

type room struct {
	forward chan []byte
	join    chan *client
	leave   chan *client
	clients map[*client]bool
}

func (r *room) run() {
	for {
		select {
		case client := <-r.join:
			//join
			r.clients[client] = true
		case client := <-r.leave:
			//leave
			delete(r.clients, client)
			close(client.send)
		case msg := <-r.forward:
			//all clients send message
			for client := range r.clients {
				select {
				case client.send <- msg:
					//send message
				default:
					//feal send message
					delete(r.clients, client)
					close(client.send)
				}
			}
		}
	}
}
