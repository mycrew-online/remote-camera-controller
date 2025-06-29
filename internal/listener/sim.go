package listener

import (
	"fmt"

	"github.com/mrlm-net/go-logz/pkg/logger"
	"github.com/mrlm-net/simconnect/pkg/types"
	"github.com/mycrew-online/remote-camera-controller/internal/manager"
)

// HandleSimEvent processes Sim system events (event ID 107) and updates state.
func HandleSimEvent(log *logger.Logger, mgr *manager.SimConnectManager, event *types.SIMCONNECT_RECV_EVENT) {
	mgr.SimulatorState().SetSim(int(event.DwData))
	log.Debug(fmt.Sprintf("[SimConnectManager] Sim event received: dwData=%d", event.DwData))
}
