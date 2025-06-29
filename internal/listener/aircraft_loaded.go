package listener

import (
	"fmt"

	"github.com/mrlm-net/go-logz/pkg/logger"
	"github.com/mrlm-net/simconnect/pkg/types"
	"github.com/mycrew-online/remote-camera-controller/internal/manager"
)

// HandleAircraftLoadedEvent processes AircraftLoaded system events (event ID 101) and updates state.
func HandleAircraftLoadedEvent(log *logger.Logger, mgr *manager.SimConnectManager, event *types.SIMCONNECT_RECV_EVENT) {
	mgr.SimulatorState().SetAircraftLoaded(fmt.Sprintf("%d", event.DwData))
	log.Info(fmt.Sprintf("[SimConnectManager] AircraftLoaded event received: dwData=%d", event.DwData))
}
