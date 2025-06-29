package listener

import (
	"fmt"

	"github.com/mrlm-net/go-logz/pkg/logger"
	"github.com/mrlm-net/simconnect/pkg/types"
	"github.com/mycrew-online/remote-camera-controller/internal/manager"
)

// HandlePositionChangedEvent processes PositionChanged system events (event ID 106) and updates state.
func HandlePositionChangedEvent(log *logger.Logger, mgr *manager.SimConnectManager, event *types.SIMCONNECT_RECV_EVENT) {
	mgr.SimulatorState().SetPositionChanged(int(event.DwData))
	log.Info(fmt.Sprintf("[SimConnectManager] PositionChanged event received: dwData=%d", event.DwData))
}
