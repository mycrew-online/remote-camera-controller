package manager

import (
	"fmt"
	"sync"

	"github.com/mrlm-net/go-logz/pkg/logger"
)

// SimulatorState holds the current tracked state from the simulator.
type SimulatorState struct {
	mu             sync.RWMutex
	AircraftLoaded string
	FlightLoaded   string
	Sim            int
	Pause          int
	Crashed        int
	View           int
	logger         *logger.Logger

	onStateChange func()
}

func NewSimulatorState() *SimulatorState {
	return &SimulatorState{}
}

func NewSimulatorStateWithLogger(log *logger.Logger) *SimulatorState {
	return &SimulatorState{logger: log}
}

// Getters and setters for each field, all thread-safe.
func (s *SimulatorState) SetAircraftLoaded(val string) {
	s.mu.Lock()
	s.AircraftLoaded = val
	if s.logger != nil {
		s.logger.Debug(fmt.Sprintf("[SimulatorState] AircraftLoaded set to %q", val))
	}
	s.mu.Unlock()
	if s.onStateChange != nil {
		s.onStateChange()
	}
}
func (s *SimulatorState) GetAircraftLoaded() string {
	s.mu.RLock()
	defer s.mu.RUnlock()
	return s.AircraftLoaded
}

func (s *SimulatorState) SetFlightLoaded(val string) {
	s.mu.Lock()
	s.FlightLoaded = val
	if s.logger != nil {
		s.logger.Debug(fmt.Sprintf("[SimulatorState] FlightLoaded set to %q", val))
	}
	s.mu.Unlock()
	if s.onStateChange != nil {
		s.onStateChange()
	}
}
func (s *SimulatorState) GetFlightLoaded() string {
	s.mu.RLock()
	defer s.mu.RUnlock()
	return s.FlightLoaded
}

func (s *SimulatorState) SetSim(val int) {
	s.mu.Lock()
	s.Sim = val
	if s.logger != nil {
		s.logger.Debug(fmt.Sprintf("[SimulatorState] Sim set to %d", val))
	}
	s.mu.Unlock()
	if s.onStateChange != nil {
		s.onStateChange()
	}
}
func (s *SimulatorState) GetSim() int {
	s.mu.RLock()
	defer s.mu.RUnlock()
	return s.Sim
}

func (s *SimulatorState) SetPause(val int) {
	s.mu.Lock()
	s.Pause = val
	if s.logger != nil {
		s.logger.Debug(fmt.Sprintf("[SimulatorState] Pause set to %d", val))
	}
	s.mu.Unlock()
	if s.onStateChange != nil {
		s.onStateChange()
	}
}
func (s *SimulatorState) GetPause() int {
	s.mu.RLock()
	defer s.mu.RUnlock()
	return s.Pause
}

func (s *SimulatorState) SetCrashed(val int) {
	s.mu.Lock()
	s.Crashed = val
	if s.logger != nil {
		s.logger.Debug(fmt.Sprintf("[SimulatorState] Crashed set to %d", val))
	}
	s.mu.Unlock()
	if s.onStateChange != nil {
		s.onStateChange()
	}
}
func (s *SimulatorState) GetCrashed() int {
	s.mu.RLock()
	defer s.mu.RUnlock()
	return s.Crashed
}

func (s *SimulatorState) SetView(val int) {
	s.mu.Lock()
	s.View = val
	if s.logger != nil {
		s.logger.Debug(fmt.Sprintf("[SimulatorState] View set to %d", val))
	}
	s.mu.Unlock()
	if s.onStateChange != nil {
		s.onStateChange()
	}
}
func (s *SimulatorState) GetView() int {
	s.mu.RLock()
	defer s.mu.RUnlock()
	return s.View
}

func (s *SimulatorState) SetOnStateChange(cb func()) {
	s.onStateChange = cb
}
