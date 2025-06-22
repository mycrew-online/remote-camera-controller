package internal

import (
	"context"
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
			fmt.Println("[SimConnectManager] Received event:", event)
			// we should not process messages while there is no connection
			if !mgr.IsOnline() {
				return
			}

			if event.MessageType == types.SIMCONNECT_RECV_ID_OPEN {
				fmt.Println("[SimConnectManager] Simulator loaded successfully.")
			}

			if event.MessageType == types.SIMCONNECT_RECV_ID_QUIT {
				fmt.Println("[SimConnectManager] Simulator disconnected, switching to Offline.")
				mgr.SetOffline()
			}
		}
	}()

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	sigs := make(chan os.Signal, 1)
	signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)
	go func() {
		sig := <-sigs
		fmt.Printf("Received signal: %v\n", sig)
		cancel()
	}()

	<-ctx.Done()

	fmt.Println("Shutdown signal received. Stopping connection...")
	mgr.StopConnection()
	mgr.Disconnect()
	fmt.Println("SimConnectManager stopped and disconnected.")
	return mgr
}
