package server

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"sync"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	"github.com/mycrew-online/remote-camera-controller/internal/manager"
)

// Server wraps the Gin engine.
type Server struct {
	engine  *gin.Engine
	mgr     *manager.SimConnectManager
	clients map[*websocket.Conn]struct{}
	mu      sync.Mutex
}

// New creates a new Server instance serving static files from the given directory and manager.

func New(staticDir string, mgr *manager.SimConnectManager) *Server {
	s := &Server{
		engine:  gin.New(),
		mgr:     mgr,
		clients: make(map[*websocket.Conn]struct{}),
	}

	s.engine.Use(gin.Recovery()) // Use Gin's recovery middleware

	// Register centralized logger middleware
	s.engine.Use(GinLoggerMiddleware(mgr))

	// WebSocket endpoint (register first)
	s.engine.GET("/ws", func(c *gin.Context) {
		var upgrader = websocket.Upgrader{
			CheckOrigin: func(r *http.Request) bool { return true },
		}
		conn, err := upgrader.Upgrade(c.Writer, c.Request, nil)
		if err != nil {
			log.Println("WebSocket upgrade error:", err)
			return
		}
		s.registerClient(conn)
		s.mgr.Logger().Debug("[WebSocketClient] connected")
		defer func() {
			s.unregisterClient(conn)
			conn.Close()
			s.mgr.Logger().Debug("[WebSocketClient] disconnected")
		}()

		// Send current state as JSON
		s.sendState(conn)

		// Log messages from client
		for {
			_, msg, err := conn.ReadMessage()
			if err != nil {
				s.mgr.Logger().Debug(fmt.Sprintf("[WebSocketClient] read error: %v", err))
				break
			}
			log.Printf("Received from client: %s\n", msg)
		}
	})

	// Serve static files (SPA) from /build, but only for files, not as a catch-all
	s.engine.StaticFS("/_app", http.Dir(staticDir+"/build/_app"))

	// Fallback to index.html for SPA routes
	s.engine.NoRoute(func(c *gin.Context) {
		c.File(staticDir + "/build/index.html")
	})

	return s
}

func (s *Server) registerClient(conn *websocket.Conn) {
	s.mu.Lock()
	s.clients[conn] = struct{}{}
	s.mu.Unlock()
}

func (s *Server) unregisterClient(conn *websocket.Conn) {
	s.mu.Lock()
	delete(s.clients, conn)
	s.mu.Unlock()
}

func (s *Server) sendState(conn *websocket.Conn) {
	type StatePayload struct {
		ConnectionState string      `json:"connectionState"`
		SimulatorState  interface{} `json:"simulatorState"`
	}
	state := StatePayload{
		ConnectionState: "Unknown",
		SimulatorState:  nil,
	}
	if s.mgr != nil {
		if s.mgr.IsOnline() {
			state.ConnectionState = "Online"
		} else {
			state.ConnectionState = "Offline"
		}
		state.SimulatorState = s.mgr.SimulatorState()
	}
	if data, err := json.Marshal(state); err == nil {
		conn.WriteMessage(websocket.TextMessage, data)
	}
}

// BroadcastState sends the current state to all connected clients.
func (s *Server) BroadcastState() {
	s.mu.Lock()
	defer s.mu.Unlock()
	for conn := range s.clients {
		s.sendState(conn)
	}
}

// Run starts the HTTP server on the given address.
func (s *Server) Run(addr string) error {
	return s.engine.Run(addr)
}
