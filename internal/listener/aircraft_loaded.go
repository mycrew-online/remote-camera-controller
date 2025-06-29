package listener

import (
	"fmt"

	"github.com/mrlm-net/go-logz/pkg/logger"
	"github.com/mrlm-net/simconnect/pkg/types"
)

// HandleAircraftLoadedEvent processes AircraftLoaded system events (event ID 101)
func HandleAircraftLoadedEvent(log *logger.Logger, event *types.SIMCONNECT_RECV_EVENT) {
	log.Info(fmt.Sprintf("[SimConnectManager] AircraftLoaded event received: dwData=%d", event.DwData))
}
