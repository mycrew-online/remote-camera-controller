package listener

import (
	"fmt"

	"github.com/mrlm-net/go-logz/pkg/logger"
	"github.com/mrlm-net/simconnect/pkg/types"
)

// HandleSystemState processes SIMCONNECT_RECV_ID_SYSTEM_STATE events.
func HandleSystemState(log *logger.Logger, event *types.SIMCONNECT_RECV_SYSTEM_STATE) {
	szString := ""
	for i, b := range event.SzString {
		if b == 0 {
			szString = string(event.SzString[:i])
			break
		}
	}
	var stateName string
	var isStringState, isIntegerState bool
	switch event.DwRequestID {
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
		log.Info(fmt.Sprintf("[SimConnectManager] SystemState Response | %s (RequestID: %d): String: %q", stateName, event.DwRequestID, szString))
	} else if isIntegerState {
		log.Info(fmt.Sprintf("[SimConnectManager] SystemState Response | %s (RequestID: %d): Integer: %d", stateName, event.DwRequestID, event.DwInteger))
	} else {
		log.Info(fmt.Sprintf("[SimConnectManager] SystemState Response | RequestID: %d | Integer: %d | Float (raw uint32): %d | String: %q", event.DwRequestID, event.DwInteger, event.DwFloat, szString))
	}
}
