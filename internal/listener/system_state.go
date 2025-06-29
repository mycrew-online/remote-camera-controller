package listener

import (
	"fmt"

	"github.com/mrlm-net/go-logz/pkg/logger"
	"github.com/mrlm-net/simconnect/pkg/types"
	"github.com/mycrew-online/remote-camera-controller/internal/manager"
)

// HandleSystemState processes SIMCONNECT_RECV_ID_SYSTEM_STATE events and updates the state store.
func HandleSystemState(log *logger.Logger, mgr *manager.SimConnectManager, event *types.SIMCONNECT_RECV_SYSTEM_STATE) {
	szString := ""
	for i, b := range event.SzString {
		if b == 0 {
			szString = string(event.SzString[:i])
			break
		}
	}
	// removed unused variables and duplicate switch
	switch event.DwRequestID {
	case 1:
		mgr.SimulatorState().SetAircraftLoaded(szString)
		log.Debug(fmt.Sprintf("[SimConnectManager] SystemState Response | AircraftLoaded (RequestID: %d): String: %q", event.DwRequestID, szString))
	case 2:
		mgr.SimulatorState().SetDialogMode(int(event.DwInteger))
		log.Debug(fmt.Sprintf("[SimConnectManager] SystemState Response | DialogMode (RequestID: %d): Integer: %d", event.DwRequestID, event.DwInteger))
	case 3:
		mgr.SimulatorState().SetFlightLoaded(szString)
		log.Debug(fmt.Sprintf("[SimConnectManager] SystemState Response | FlightLoaded (RequestID: %d): String: %q", event.DwRequestID, szString))
	case 4:
		mgr.SimulatorState().SetFlightPlan(szString)
		log.Debug(fmt.Sprintf("[SimConnectManager] SystemState Response | FlightPlan (RequestID: %d): String: %q", event.DwRequestID, szString))
	case 5:
		mgr.SimulatorState().SetSim(int(event.DwInteger))
		log.Debug(fmt.Sprintf("[SimConnectManager] SystemState Response | Sim (RequestID: %d): Integer: %d", event.DwRequestID, event.DwInteger))
	default:
		log.Debug(fmt.Sprintf("[SimConnectManager] SystemState Response | RequestID: %d | Integer: %d | Float (raw uint32): %d | String: %q", event.DwRequestID, event.DwInteger, event.DwFloat, szString))
	}
}
