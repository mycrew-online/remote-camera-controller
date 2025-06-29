package listener

import (
	"fmt"

	"github.com/mrlm-net/go-logz/pkg/logger"
	"github.com/mrlm-net/simconnect/pkg/types"
)

// HandleFlightPlanDeactivatedEvent processes FlightPlanDeactivated system events (event ID 105)
func HandleFlightPlanDeactivatedEvent(log *logger.Logger, event *types.SIMCONNECT_RECV_EVENT) {
	log.Info(fmt.Sprintf("[SimConnectManager] FlightPlanDeactivated event received: dwData=%d", event.DwData))
}
