package manager

import "sync"

// ConnectionState represents the state of the SimConnect connection.
type ConnectionState int

const (
	Offline ConnectionState = iota
	Connecting
	Online
)

type SimulatorState struct {
	mutex sync.RWMutex

	loaded bool
	paused bool
}
