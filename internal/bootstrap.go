package internal

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"github.com/mrlm-net/simconnect/pkg/types"
	"github.com/mycrew-online/remote-camera-controller/internal/manager"
)

// Bootstrap initializes the SimConnectManager, starts the connection loop, launches event handler, and handles graceful shutdown.
func (app *Application) Bootstrap() *manager.SimConnectManager {
	mgr := app.SimConnectManager
	log := mgr.Logger()

	log.Info("[SimConnectManager] App started. Waiting for shutdown signal (Ctrl+C)...")
	mgr.StartConnection()

	// Start event handler goroutine for disconnects
	go func() {
		for event := range mgr.Stream() {
			log.Info(fmt.Sprintf("[SimConnectManager] Received event: %v", event.MessageType))
			// we should not process messages while there is no connection
			if !mgr.IsOnline() {
				return
			}

			if event.MessageType == types.SIMCONNECT_RECV_ID_OPEN {
				log.Info("[SimConnectManager] Simulator loaded successfully.")
				// Request system states after connection is established
				mgr.Client().RequestSystemState(1, "AircraftLoaded")
				mgr.Client().RequestSystemState(2, "DialogMode")
				mgr.Client().RequestSystemState(3, "FlightLoaded")
				mgr.Client().RequestSystemState(4, "FlightPlan")
				mgr.Client().RequestSystemState(5, "Sim")
			}

			// Log all system state messages for inspection
			if event.MessageType == types.SIMCONNECT_RECV_ID_SYSTEM_STATE {
				if sysState, ok := event.Data.(*types.SIMCONNECT_RECV_SYSTEM_STATE); ok {
					// Convert SzString (fixed array) to Go string, trim at null terminator
					szString := ""
					for i, b := range sysState.SzString {
						if b == 0 {
							szString = string(sysState.SzString[:i])
							break
						}
					}
					// Map request IDs to their meaning and type
					var stateName string
					var isStringState, isIntegerState bool
					switch sysState.DwRequestID {
					case 1:
						stateName = "AircraftLoaded"
						isStringState = true
					case 2:
						stateName = "DialogMode"
						isIntegerState = true
					case 3:
						stateName = "FlightLoaded"
						isStringState = true
					case 4:
						stateName = "FlightPlan"
						isStringState = true
					case 5:
						stateName = "Sim"
						isIntegerState = true
					default:
						stateName = "Unknown"
					}

					if isStringState {
						log.Info(fmt.Sprintf(
							"[SimConnectManager] SystemState Response | %s (RequestID: %d): String: %q",
							stateName, sysState.DwRequestID, szString,
						))
					} else if isIntegerState {
						log.Info(fmt.Sprintf(
							"[SimConnectManager] SystemState Response | %s (RequestID: %d): Integer: %d",
							stateName, sysState.DwRequestID, sysState.DwInteger,
						))
					} else {
						log.Info(fmt.Sprintf(
							"[SimConnectManager] SystemState Response | RequestID: %d | Integer: %d | Float (raw uint32): %d | String: %q",
							sysState.DwRequestID,
							sysState.DwInteger,
							sysState.DwFloat,
							szString,
						))
					}
				} else {
					log.Info(fmt.Sprintf("[SimConnectManager] Received system state (raw): %+v", event))
				}
			}
			if event.MessageType == types.SIMCONNECT_RECV_ID_QUIT {
				log.Info("[SimConnectManager] Simulator disconnected, switching to Offline.")
				mgr.SetOffline()
			}
		}
	}()

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	sigs := make(chan os.Signal, 1)
	signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)
	go func() {
		sig := <-sigs
		log.Info(fmt.Sprintf("[SimConnectManager] Received signal: %v\n", sig))
		cancel()
	}()

	<-ctx.Done()

	log.Info("[SimConnectManager] Shutdown signal received. Stopping connection...")
	mgr.StopConnection()
	log.Info("[SimConnectManager] Stopped and disconnected.")
	return mgr
}
