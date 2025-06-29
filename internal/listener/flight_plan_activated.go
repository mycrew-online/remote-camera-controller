package listener

import (
	"fmt"

	"github.com/mrlm-net/go-logz/pkg/logger"
	"github.com/mrlm-net/simconnect/pkg/types"
)

// HandleFlightPlanActivatedEvent processes FlightPlanActivated system events (event ID 104)
func HandleFlightPlanActivatedEvent(log *logger.Logger, event *types.SIMCONNECT_RECV_EVENT) {
	log.Info(fmt.Sprintf("[SimConnectManager] FlightPlanActivated event received: dwData=%d", event.DwData))
}
