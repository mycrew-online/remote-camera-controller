package internal

import "github.com/mycrew-online/remote-camera-controller/internal/manager"

type Application struct {
	// SimConnectManager handles the connection to the simulator.
	SimConnectManager *manager.SimConnectManager
	System            *SimulatorState
}
