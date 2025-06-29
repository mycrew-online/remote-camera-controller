package listener

import (
	"fmt"

	"github.com/mrlm-net/go-logz/pkg/logger"
	"github.com/mrlm-net/simconnect/pkg/types"
	"github.com/mycrew-online/remote-camera-controller/internal/manager"
)

// HandleFlightPlanActivatedEvent processes FlightPlanActivated system events (event ID 104) and updates state.
func HandleFlightPlanActivatedEvent(log *logger.Logger, mgr *manager.SimConnectManager, event *types.SIMCONNECT_RECV_EVENT) {
	mgr.SimulatorState().SetFlightPlanActivated(int(event.DwData))
	log.Info(fmt.Sprintf("[SimConnectManager] FlightPlanActivated event received: dwData=%d", event.DwData))
}
