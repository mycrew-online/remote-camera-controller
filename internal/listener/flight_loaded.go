package listener

import (
	"fmt"

	"github.com/mrlm-net/go-logz/pkg/logger"
	"github.com/mrlm-net/simconnect/pkg/types"
	"github.com/mycrew-online/remote-camera-controller/internal/manager"
)

// HandleFlightLoadedEvent processes FlightLoaded system events (event ID 102) and updates state.
func HandleFlightLoadedEvent(log *logger.Logger, mgr *manager.SimConnectManager, event *types.SIMCONNECT_RECV_EVENT) {
	mgr.SimulatorState().SetFlightLoaded(fmt.Sprintf("%d", event.DwData))
	log.Info(fmt.Sprintf("[SimConnectManager] FlightLoaded event received: dwData=%d", event.DwData))
}
