package manager

import (
	"fmt"
	"sync"

	"github.com/mrlm-net/go-logz/pkg/logger"
)

// SimulatorState holds the current tracked state from the simulator.
type SimulatorState struct {
	mu                    sync.RWMutex
	AircraftLoaded        string
	DialogMode            int
	FlightLoaded          string
	FlightPlan            string
	Sim                   int
	Pause                 int
	Crashed               int
	FlightPlanActivated   int
	FlightPlanDeactivated int
	PositionChanged       int
	View                  int
	logger                *logger.Logger
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
}
func (s *SimulatorState) GetAircraftLoaded() string {
	s.mu.RLock()
	defer s.mu.RUnlock()
	return s.AircraftLoaded
}

func (s *SimulatorState) SetDialogMode(val int) {
	s.mu.Lock()
	s.DialogMode = val
	if s.logger != nil {
		s.logger.Debug(fmt.Sprintf("[SimulatorState] DialogMode set to %d", val))
	}
	s.mu.Unlock()
}
func (s *SimulatorState) GetDialogMode() int {
	s.mu.RLock()
	defer s.mu.RUnlock()
	return s.DialogMode
}

func (s *SimulatorState) SetFlightLoaded(val string) {
	s.mu.Lock()
	s.FlightLoaded = val
	if s.logger != nil {
		s.logger.Debug(fmt.Sprintf("[SimulatorState] FlightLoaded set to %q", val))
	}
	s.mu.Unlock()
}
func (s *SimulatorState) GetFlightLoaded() string {
	s.mu.RLock()
	defer s.mu.RUnlock()
	return s.FlightLoaded
}

func (s *SimulatorState) SetFlightPlan(val string) {
	s.mu.Lock()
	s.FlightPlan = val
	if s.logger != nil {
		s.logger.Debug(fmt.Sprintf("[SimulatorState] FlightPlan set to %q", val))
	}
	s.mu.Unlock()
}
func (s *SimulatorState) GetFlightPlan() string {
	s.mu.RLock()
	defer s.mu.RUnlock()
	return s.FlightPlan
}

func (s *SimulatorState) SetSim(val int) {
	s.mu.Lock()
	s.Sim = val
	if s.logger != nil {
		s.logger.Debug(fmt.Sprintf("[SimulatorState] Sim set to %d", val))
	}
	s.mu.Unlock()
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
}
func (s *SimulatorState) GetCrashed() int {
	s.mu.RLock()
	defer s.mu.RUnlock()
	return s.Crashed
}

func (s *SimulatorState) SetFlightPlanActivated(val int) {
	s.mu.Lock()
	s.FlightPlanActivated = val
	if s.logger != nil {
		s.logger.Debug(fmt.Sprintf("[SimulatorState] FlightPlanActivated set to %d", val))
	}
	s.mu.Unlock()
}
func (s *SimulatorState) GetFlightPlanActivated() int {
	s.mu.RLock()
	defer s.mu.RUnlock()
	return s.FlightPlanActivated
}

func (s *SimulatorState) SetFlightPlanDeactivated(val int) {
	s.mu.Lock()
	s.FlightPlanDeactivated = val
	if s.logger != nil {
		s.logger.Debug(fmt.Sprintf("[SimulatorState] FlightPlanDeactivated set to %d", val))
	}
	s.mu.Unlock()
}
func (s *SimulatorState) GetFlightPlanDeactivated() int {
	s.mu.RLock()
	defer s.mu.RUnlock()
	return s.FlightPlanDeactivated
}

func (s *SimulatorState) SetPositionChanged(val int) {
	s.mu.Lock()
	s.PositionChanged = val
	if s.logger != nil {
		s.logger.Debug(fmt.Sprintf("[SimulatorState] PositionChanged set to %d", val))
	}
	s.mu.Unlock()
}
func (s *SimulatorState) GetPositionChanged() int {
	s.mu.RLock()
	defer s.mu.RUnlock()
	return s.PositionChanged
}

func (s *SimulatorState) SetView(val int) {
	s.mu.Lock()
	s.View = val
	if s.logger != nil {
		s.logger.Debug(fmt.Sprintf("[SimulatorState] View set to %d", val))
	}
	s.mu.Unlock()
}
func (s *SimulatorState) GetView() int {
	s.mu.RLock()
	defer s.mu.RUnlock()
	return s.View
}
