package listener

import (
	"github.com/mrlm-net/go-logz/pkg/logger"
	"github.com/mycrew-online/remote-camera-controller/internal/manager"
)

// HandleQuit processes SIMCONNECT_RECV_ID_QUIT events.
func HandleQuit(log *logger.Logger, mgr *manager.SimConnectManager) {
	log.Info("[SimConnectManager] Simulator disconnected, switching to Offline.")
	mgr.SetOffline()
}
