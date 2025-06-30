package listener

import (
	"fmt"

	"github.com/mrlm-net/go-logz/pkg/logger"
	"github.com/mrlm-net/simconnect/pkg/types"
	"github.com/mycrew-online/remote-camera-controller/internal/manager"
)

// HandleAircraftLoadedFilenameEvent processes AircraftLoaded events received as SIMCONNECT_RECV_EVENT_FILENAME (event ID 101)
func HandleAircraftLoadedFilenameEvent(log *logger.Logger, mgr *manager.SimConnectManager, event *types.SIMCONNECT_RECV_EVENT_FILENAME) {
	filename := ""
	for i, b := range event.SzFileName {
		if b == 0 {
			filename = string(event.SzFileName[:i])
			break
		}
	}
	mgr.SimulatorState().SetAircraftLoaded(filename)
	log.Info(fmt.Sprintf("[SimConnectManager] AircraftLoaded (filename) event received: filename=%q", filename))
}
