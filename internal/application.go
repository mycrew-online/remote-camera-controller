package internal

import "github.com/mycrew-online/remote-camera-controller/internal/manager"

type Application struct {
	// SimConnectManager handles the connection to the simulator.
	SimConnectManager *manager.SimConnectManager
	System            *SimulatorState
}

// NewApplicationWithOptions creates an Application with a log level option.
func NewApplicationWithOptions(logLevel string) *Application {
	return &Application{
		SimConnectManager: manager.NewSimConnectManagerWithOptions(logLevel),
	}
}

// NewApplication is kept for backward compatibility, defaults to info log level.
func NewApplication() *Application {
	return NewApplicationWithOptions("info")
}
