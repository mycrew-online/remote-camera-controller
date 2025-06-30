package internal

import (
	"github.com/mycrew-online/remote-camera-controller/internal/manager"
	"github.com/mycrew-online/remote-camera-controller/internal/server"
)

type Application struct {
	SimConnectManager *manager.SimConnectManager
	Server            *server.Server
	System            *SimulatorState
}

// NewApplicationWithOptions creates an Application with a log level option.
func NewApplicationWithOptions(logLevel string) *Application {
	mgr := manager.NewSimConnectManagerWithOptions(logLevel)
	srv := server.New("website", mgr)
	app := &Application{
		SimConnectManager: mgr,
		Server:            srv,
	}
	mgr.SetOnStateChange(srv.BroadcastState)
	return app
}

// NewApplication is kept for backward compatibility, defaults to info log level.
func NewApplication() *Application {
	return NewApplicationWithOptions("info")
}
