package manager

import (
	"fmt"
	"time"
)

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
				m.logger.Debug("[SimConnectManager] Connection loop stopped.")
				return
			default:
				m.stateMu.Lock()
				if m.state == Offline {
					m.logger.Debug("[SimConnectManager] State is Offline, will try to connect.")
					m.state = Connecting
					m.stateMu.Unlock()
					m.connect() // Use real connection logic
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
		m.disconnect() // Ensure we disconnect when stopping
	}
}

// Connect attempts to establish a connection to the simulator.
func (m *SimConnectManager) connect() {
	m.logger.Info("[SimConnectManager] Attempting to connect...")
	err := m.client.Connect()
	m.stateMu.Lock()
	defer m.stateMu.Unlock()
	if err != nil {
		m.logger.Debug(fmt.Sprintf("[SimConnectManager] Connection failed: %v", err))
		m.state = Offline
		return
	}
	m.logger.Info("[SimConnectManager] Connected successfully.")
	m.state = Online
}

// Disconnect closes the connection to the simulator.
func (m *SimConnectManager) disconnect() {
	m.logger.Debug("[SimConnectManager] Disconnecting...")
	if m.client != nil {
		_ = m.client.Disconnect() // ignore error for now
	}
	m.stateMu.Lock()
	m.state = Offline
	m.stateMu.Unlock()
	m.logger.Debug("[SimConnectManager] Disconnected.")
}
