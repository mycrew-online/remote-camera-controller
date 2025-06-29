package listener

import (
	"fmt"

	"github.com/mrlm-net/go-logz/pkg/logger"
	"github.com/mrlm-net/simconnect/pkg/types"
	"github.com/mycrew-online/remote-camera-controller/internal/manager"
)

// HandleFlightPlanDeactivatedEvent processes FlightPlanDeactivated system events (event ID 105) and updates state.
func HandleFlightPlanDeactivatedEvent(log *logger.Logger, mgr *manager.SimConnectManager, event *types.SIMCONNECT_RECV_EVENT) {
	mgr.SimulatorState().SetFlightPlanDeactivated(int(event.DwData))
	log.Debug(fmt.Sprintf("[SimConnectManager] FlightPlanDeactivated event received: dwData=%d", event.DwData))
}
