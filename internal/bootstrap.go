package internal

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"github.com/mrlm-net/simconnect/pkg/types"
	"github.com/mycrew-online/remote-camera-controller/internal/manager"
)

// Bootstrap initializes the SimConnectManager, starts the connection loop, launches event handler, and handles graceful shutdown.
func Bootstrap() *manager.SimConnectManager {
	mgr := manager.NewSimConnectManager()
	mgr.StartConnection()
	fmt.Println("SimConnectManager started. Waiting for shutdown signal (Ctrl+C)...")

	// Start event handler goroutine for disconnects
	go func() {
		for event := range mgr.Client().Stream() {
			if event.MessageType == types.SIMCONNECT_RECV_ID_QUIT {
				fmt.Println("[SimConnectManager] Simulator disconnected, switching to Offline.")
				mgr.SetOffline()
			}
		}
	}()

	sigs := make(chan os.Signal, 1)
	signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)
	<-sigs

	fmt.Println("Shutdown signal received. Stopping connection...")
	mgr.StopConnection()
	mgr.Disconnect()
	fmt.Println("SimConnectManager stopped and disconnected.")
	return mgr
}
