package websocket

import (
	"log"
	"log/slog"
	"net/http"
	"slices"
	"sync"

	"github.com/gorilla/websocket"
)

type WebSocketManager struct {
	upgrader       websocket.Upgrader
	connections    map[*websocket.Conn]bool
	mu             sync.RWMutex
	allowedOrigins []string
}

func NewWebSocketManager(allowedOrigins []string) *WebSocketManager {
	return &WebSocketManager{
		upgrader: websocket.Upgrader{
			CheckOrigin: func(r *http.Request) bool {
				origin := r.Header.Get("Origin")

				if len(allowedOrigins) == 0 {
					log.Fatal("No allowed origins specified")
				}

				return slices.Contains(allowedOrigins, origin)
			},
		},
		connections:    make(map[*websocket.Conn]bool),
		allowedOrigins: allowedOrigins,
	}
}

// HandleConnection - handles a new websocket connection
func (wm *WebSocketManager) HandleConnection(w http.ResponseWriter, r *http.Request) {
	conn, err := wm.upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Printf("Failed to upgrade connection: %v", err)
		return
	}

	wm.mu.Lock()
	wm.connections[conn] = true
	wm.mu.Unlock()

	defer func() {
		wm.mu.Lock()
		delete(wm.connections, conn)
		wm.mu.Unlock()
		conn.Close()

		slog.Error("Websocket closed")
	}()

	for {
		_, _, err := conn.ReadMessage()
		if err != nil {
			slog.Error("Websocket read error", "error", err)
			break
		}
	}
}

// Broadcast - broadcasts a message
func (wm *WebSocketManager) Broadcast(message []byte) {
	for conn := range wm.connections {
		err := conn.WriteMessage(websocket.TextMessage, message)
		if err != nil {
			slog.Error("Failed to write message to websocket", "error", err)
			conn.Close()

			wm.mu.Lock()
			delete(wm.connections, conn)
			wm.mu.Unlock()
		}
	}
}
