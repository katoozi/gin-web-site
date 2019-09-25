package website

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/gorilla/websocket"
)

// Packet represents application level data.
type Packet struct {
	ActionType  string `json:"action_type"`
	Text        string `json:"text"`
	MessageType int
}

// Channel wraps user connection.
type Channel struct {
	conn *websocket.Conn // WebSocket connection.
	send chan Packet     // Outgoing packets queue.
}

func (ch *Channel) reader() {
	for {
		messageType, p, err := ch.conn.ReadMessage()
		if err != nil {
			return
		}
		pkt := readPacket(p, messageType)
		ch.send <- *pkt
	}
}

func (ch *Channel) writer() {
	for pkt := range ch.send {
		data, err := json.Marshal(pkt)
		if err != nil {
			log.Printf("Error while marshaling json: %v", err)
		}
		if err := ch.conn.WriteMessage(pkt.MessageType, data); err != nil {
			return
		}
	}
}

func (ch *Channel) receiveNotification() {
	log.Println("receive notification called!!!")

}

// NewChannel is the Channel factorty function
func NewChannel(conn *websocket.Conn) *Channel {
	c := &Channel{
		conn: conn,
		send: make(chan Packet, 10),
	}

	go c.reader()
	go c.writer()

	q, err := RabbitMQCon.QueueDeclare(
		"hello",
		false,
		false,
		false,
		false,
		nil,
	)
	if err != nil {
		log.Printf("Failed to open queue, %v\n", err)
	}
	err = RabbitMQCon.Qos(
		1,     // prefetch count
		0,     // prefetch size
		false, // global
	)
	if err != nil {
		log.Println("Failed to set Qos")
	}

	msgs, err := RabbitMQCon.Consume(
		q.Name,
		"",
		false,
		false,
		false,
		false,
		nil,
	)
	if err != nil {
		log.Printf("Failed to register consumer: %v\n", err)
	}

	go func() {
		for d := range msgs {
			log.Printf("Received a message: %s\n", d.Body)
			pkt := readPacket(d.Body, 1)
			c.send <- *pkt
			d.Ack(false)
		}
	}()

	return c
}

func readPacket(data []byte, messageType int) *Packet {
	packetData := &Packet{}
	err := json.Unmarshal(data, packetData)
	if err != nil {
		log.Printf("Error while unmarshal json: %v", err)
	}
	packetData.MessageType = messageType
	return packetData
}

// We'll need to define an Upgrader
// this will require a Read and Write buffer size
var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,

	// We'll need to check the origin of our connection
	// this will allow us to make requests from our React
	// development server to here.
	// For now, we'll do no checking and just allow any connection
	// in production must be nil
	// in develop must be `func(r *http.Request) bool { return true }`
	CheckOrigin: func(r *http.Request) bool { return true },
}

// define our WebSocket endpoint
func serveWs(w http.ResponseWriter, r *http.Request) {
	// upgrade this connection to a WebSocket
	// connection
	ws, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Printf("Error while upgrade connection: %v", err)
		return
	}
	// listen indefinitely for new messages coming
	// through on our WebSocket connection
	// ch := NewChannel(ws)
	NewChannel(ws)
}
