package internal

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"github.com/mycrew-online/remote-camera-controller/internal/listener"
	"github.com/mycrew-online/remote-camera-controller/internal/manager"
)

// Bootstrap initializes the SimConnectManager, starts the connection loop, launches event handler, and handles graceful shutdown.
func (app *Application) Bootstrap() *manager.SimConnectManager {
	mgr := app.SimConnectManager
	log := mgr.Logger()

	log.Info("[SimConnectManager] App started. Waiting for shutdown signal (Ctrl+C)...")
	mgr.StartConnection()

	// Start event handler goroutine
	listener.StartSimConnectEventListener(mgr, log)

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	sigs := make(chan os.Signal, 1)
	signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)
	go func() {
		sig := <-sigs
		log.Info(fmt.Sprintf("[SimConnectManager] Received signal: %v\n", sig))
		cancel()
	}()

	<-ctx.Done()

	log.Info("[SimConnectManager] Shutdown signal received. Stopping connection...")
	mgr.StopConnection()
	log.Info("[SimConnectManager] Stopped and disconnected.")
	return mgr
}
