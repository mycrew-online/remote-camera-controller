package listener

import (
	"fmt"

	"github.com/mrlm-net/go-logz/pkg/logger"
	"github.com/mrlm-net/simconnect/pkg/types"
	"github.com/mycrew-online/remote-camera-controller/internal/manager"
)

// HandlePauseEvent processes Pause system events (event ID 100) and updates state.
func HandlePauseEvent(log *logger.Logger, mgr *manager.SimConnectManager, event *types.SIMCONNECT_RECV_EVENT) {
	mgr.SimulatorState().SetPause(int(event.DwData))
	log.Info(fmt.Sprintf("[SimConnectManager] Pause event received: dwData=%d", event.DwData))
}
