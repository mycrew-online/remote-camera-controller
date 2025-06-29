package listener

import (
	"fmt"

	"github.com/mrlm-net/go-logz/pkg/logger"
	"github.com/mrlm-net/simconnect/pkg/types"
	"github.com/mycrew-online/remote-camera-controller/internal/manager"
)

// HandleCrashedEvent processes Crashed system events (event ID 103) and updates state.
func HandleCrashedEvent(log *logger.Logger, mgr *manager.SimConnectManager, event *types.SIMCONNECT_RECV_EVENT) {
	mgr.SimulatorState().SetCrashed(int(event.DwData))
	log.Debug(fmt.Sprintf("[SimConnectManager] Crashed event received: dwData=%d", event.DwData))
}
