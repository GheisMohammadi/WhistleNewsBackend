package ws

// hub maintains the set of active clients and broadcasts messages to the
// clients.
type Hub struct {
	// Registered clients.
	clients map[*Client]bool

	// Inbound messages from the clients.
	broadcast chan []byte

	// Emit message for particular client
	Emit chan *Emitter

	// Register requests from the clients.
	register chan *Client

	// Unregister requests from clients.
	unregister chan *Client
}

type Emitter struct {
	Ids     []string
	Message []byte
}

func NewHub() *Hub {
	return &Hub{
		broadcast:  make(chan []byte),
		Emit:       make(chan *Emitter),
		register:   make(chan *Client),
		unregister: make(chan *Client),
		clients:    make(map[*Client]bool),
	}
}

func (h *Hub) Run() {
	for {
		select {
		case client := <-h.register:
			h.clients[client] = true
		case client := <-h.unregister:
			if _, ok := h.clients[client]; ok {
				delete(h.clients, client)
				close(client.send)
			}
		case message := <-h.broadcast:
			for client := range h.clients {
				select {
				case client.send <- message:
				default:
					close(client.send)
					delete(h.clients, client)
				}
			}
		case emitter := <-h.Emit:
			for client := range h.clients {
				for _, id := range emitter.Ids {
					if client.Id == id {
						select {
						case client.send <- emitter.Message:
						default:
							close(client.send)
							delete(h.clients, client)
						}
					}
				}
			}
		}
	}
}

func (h *Hub) GetClientCount() int {
	return len(h.clients)
}
