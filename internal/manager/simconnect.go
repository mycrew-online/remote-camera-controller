package manager

import (
	"fmt"
	"sync"
	"time"

	"github.com/mrlm-net/simconnect/pkg/client"
)

const DLL_DEFAULT_PATH = "C:/MSFS 2024 SDK/SimConnect SDK/lib/SimConnect.dll"

// ConnectionState represents the state of the SimConnect connection.
type ConnectionState int

const (
	Offline ConnectionState = iota
	Connecting
	Online
)

// SimConnectManager handles the connection to the simulator.
type SimConnectManager struct {
	state   ConnectionState
	stateMu sync.RWMutex
	stopCh  chan struct{}
	stopped sync.WaitGroup
	client  *client.Engine
}

// NewSimConnectManager creates a new SimConnectManager in Offline state.
func NewSimConnectManager() *SimConnectManager {
	return &SimConnectManager{
		client: client.NewWithDLL("NAME", DLL_DEFAULT_PATH),
		state:  Offline,
	}
}

// Connect attempts to establish a connection to the simulator.
func (m *SimConnectManager) Connect() {
	fmt.Println("[SimConnectManager] Attempting to connect...")
	err := m.client.Connect()
	m.stateMu.Lock()
	defer m.stateMu.Unlock()
	if err != nil {
		fmt.Printf("[SimConnectManager] Connection failed: %v\n", err)
		m.state = Offline
		return
	}
	fmt.Println("[SimConnectManager] Connected successfully.")
	m.state = Online
}

// Disconnect closes the connection to the simulator.
func (m *SimConnectManager) Disconnect() {
	fmt.Println("[SimConnectManager] Disconnecting...")
	if m.client != nil {
		_ = m.client.Disconnect() // ignore error for now
	}
	m.stateMu.Lock()
	m.state = Offline
	m.stateMu.Unlock()
	fmt.Println("[SimConnectManager] Disconnected.")
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

// SetOffline sets the manager state to Offline in a thread-safe way.
func (m *SimConnectManager) SetOffline() {
	m.stateMu.Lock()
	m.state = Offline
	m.stateMu.Unlock()
}

// StartConnection starts a goroutine to monitor and maintain the connection.
func (m *SimConnectManager) StartConnection() {
	m.stopCh = make(chan struct{})
	m.stopped.Add(1)
	go func() {
		defer m.stopped.Done()
		retryInterval := 3 * time.Second
		for {
			select {
			case <-m.stopCh:
				fmt.Println("[SimConnectManager] Connection loop stopped.")
				return
			default:
				m.stateMu.Lock()
				if m.state == Offline {
					fmt.Println("[SimConnectManager] State is Offline, will try to connect.")
					m.state = Connecting
					m.stateMu.Unlock()
					m.Connect() // Use real connection logic
				} else {
					m.stateMu.Unlock()
				}
				time.Sleep(retryInterval)
			}
		}
	}()
}

// StopConnection signals the monitoring goroutine to stop and waits for it to finish.
func (m *SimConnectManager) StopConnection() {
	if m.stopCh != nil {
		close(m.stopCh)
		m.stopped.Wait()
	}
}
