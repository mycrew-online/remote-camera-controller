package server

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	"github.com/mycrew-online/remote-camera-controller/internal/manager"
)

// Server wraps the Gin engine.
type Server struct {
	engine *gin.Engine
	mgr    *manager.SimConnectManager
}

// New creates a new Server instance serving static files from the given directory and manager.
func New(staticDir string, mgr *manager.SimConnectManager) *Server {
	r := gin.Default()

	// WebSocket endpoint (register first)
	r.GET("/ws", func(c *gin.Context) {
		var upgrader = websocket.Upgrader{
			CheckOrigin: func(r *http.Request) bool { return true },
		}
		conn, err := upgrader.Upgrade(c.Writer, c.Request, nil)
		if err != nil {
			log.Println("WebSocket upgrade error:", err)
			return
		}
		defer conn.Close()

		// Send current state as JSON
		type StatePayload struct {
			ConnectionState string      `json:"connectionState"`
			SimulatorState  interface{} `json:"simulatorState"`
		}
		state := StatePayload{
			ConnectionState: "Unknown",
			SimulatorState:  nil,
		}
		if mgr != nil {
			if mgr.IsOnline() {
				state.ConnectionState = "Online"
			} else {
				state.ConnectionState = "Offline"
			}
			state.SimulatorState = mgr.SimulatorState()
		}
		if data, err := json.Marshal(state); err == nil {
			conn.WriteMessage(websocket.TextMessage, data)
		}

		// Log messages from client
		for {
			_, msg, err := conn.ReadMessage()
			if err != nil {
				log.Println("WebSocket read error:", err)
				break
			}
			log.Printf("Received from client: %s\n", msg)
		}
	})

	// Serve static files (SPA) from /build, but only for files, not as a catch-all
	r.StaticFS("/_app", http.Dir(staticDir+"/build/_app"))

	// Fallback to index.html for SPA routes
	r.NoRoute(func(c *gin.Context) {
		c.File(staticDir + "/build/index.html")
	})

	return &Server{engine: r, mgr: mgr}
}

// Run starts the HTTP server on the given address.
func (s *Server) Run(addr string) error {
	return s.engine.Run(addr)
}
