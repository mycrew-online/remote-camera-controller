package manager

import (
	"sync"

	"github.com/mrlm-net/go-logz/pkg/logger"
	"github.com/mrlm-net/simconnect/pkg/client"
)

const APP_NAME = "[MyCrew.online] - Remote Camera Controller"
const DLL_DEFAULT_PATH = "C:/MSFS 2024 SDK/SimConnect SDK/lib/SimConnect.dll"

// SimConnectManager handles the connection to the simulator.
type SimConnectManager struct {
	state    ConnectionState
	stateMu  sync.RWMutex
	stopCh   chan struct{}
	stopped  sync.WaitGroup
	logger   *logger.Logger
	client   *client.Engine
	simState *SimulatorState

	onStateChange func()
}

// NewSimConnectManagerWithOptions creates a new SimConnectManager in Offline state.
func NewSimConnectManagerWithOptions(logLevel string) *SimConnectManager {
	level := parseLogLevel(logLevel)
	log := logger.NewLogger(logger.LogOptions{
		Level: level,
	})
	mgr := &SimConnectManager{
		client: client.NewWithDLL(APP_NAME, DLL_DEFAULT_PATH),
		logger: log,
		state:  Offline,
	}
	mgr.simState = NewSimulatorStateWithLogger(log)
	// Wire simulator state's onStateChange to manager's onStateChange
	mgr.simState.SetOnStateChange(func() {
		if mgr.onStateChange != nil {
			mgr.onStateChange()
		}
	})
	return mgr
}

// SimulatorState returns the current simulator state store.
func (m *SimConnectManager) SimulatorState() *SimulatorState {
	return m.simState
}

// NewSimConnectManager is kept for backward compatibility, defaults to Info level.
func NewSimConnectManager() *SimConnectManager {
	return NewSimConnectManagerWithOptions("info")
}

// parseLogLevel converts a string to logger.LogLevel, defaults to Info.
func parseLogLevel(level string) logger.LogLevel {
	switch level {
	case "debug":
		return logger.Debug
	case "info":
		fallthrough
	default:
		return logger.Info
	}
}

// IsOnline returns true if the connection is established.
func (m *SimConnectManager) IsOnline() bool {
	m.stateMu.RLock()
	defer m.stateMu.RUnlock()
	return m.state == Online
}

// Client returns the underlying SimConnect client.
func (m *SimConnectManager) Client() *client.Engine {
	return m.client
}

func (m *SimConnectManager) Stream() <-chan client.ParsedMessage {
	return m.client.Stream()
}

// Logger returns the logger used by the SimConnectManager.
func (m *SimConnectManager) Logger() *logger.Logger {
	return m.logger
}

// SetOffline sets the manager state to Offline in a thread-safe way.
func (m *SimConnectManager) SetOffline() {
	m.stateMu.Lock()
	m.state = Offline
	m.stateMu.Unlock()
	if m.onStateChange != nil {
		m.onStateChange()
	}
}

// SetOnline sets the manager state to Online in a thread-safe way.
func (m *SimConnectManager) SetOnline() {
	m.stateMu.Lock()
	m.state = Online
	m.stateMu.Unlock()
	if m.onStateChange != nil {
		m.onStateChange()
	}
}

// SetOnStateChange sets the callback for state changes.
func (m *SimConnectManager) SetOnStateChange(cb func()) {
	m.onStateChange = cb
}
