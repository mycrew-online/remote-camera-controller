package internal

import "sync"

type SimulatorState struct {
	mutex sync.RWMutex

	loaded bool
	paused bool
}
