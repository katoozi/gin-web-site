package website

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/gorilla/websocket"
)

func TestWebsocket(t *testing.T) {
	// Create test server with the serveWs handler.
	s := httptest.NewServer(http.HandlerFunc(serveWs))
	defer s.Close()

	// Convert http://127.0.0.1 to ws://127.0.0.
	u := "ws" + strings.TrimPrefix(s.URL, "http")

	// Connect to the server
	ws, _, err := websocket.DefaultDialer.Dial(u, nil)
	if err != nil {
		t.Fatalf("%v", err)
	}
	defer ws.Close()

	// Send message to server, read response and check to see if it's what we expect.
	data, _ := json.Marshal(Packet{
		ActionType:  "download",
		Text:        "Download",
		MessageType: websocket.TextMessage,
	})
	if err := ws.WriteMessage(websocket.TextMessage, data); err != nil {
		t.Fatalf("%v", err)
	}
	_, p, err := ws.ReadMessage()
	if err != nil {
		t.Fatalf("%v", err)
	}
	if string(p) != string(data) {
		t.Fatalf("bad message")
	}
}

func BenchmarkWebsocket(b *testing.B) {
	// Create test server with the serveWs handler.
	s := httptest.NewServer(http.HandlerFunc(serveWs))
	defer s.Close()

	// Convert http://127.0.0.1 to ws://127.0.0.
	u := "ws" + strings.TrimPrefix(s.URL, "http")

	// Connect to the server
	ws, _, err := websocket.DefaultDialer.Dial(u, nil)
	if err != nil {
		b.Fatalf("%v", err)
	}
	defer ws.Close()

	for i := 0; i < b.N; i++ {
		// Send message to server, read response and check to see if it's what we expect.
		data, _ := json.Marshal(Packet{
			ActionType:  "download",
			Text:        "Download",
			MessageType: websocket.TextMessage,
		})
		if err := ws.WriteMessage(websocket.TextMessage, data); err != nil {
			b.Fatalf("%v", err)
		}
		_, p, err := ws.ReadMessage()
		if err != nil {
			b.Fatalf("%v", err)
		}
		if string(p) != string(data) {
			b.Fatalf("bad message")
		}
	}
}
