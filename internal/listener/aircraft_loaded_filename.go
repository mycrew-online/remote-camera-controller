package listener

import (
	"fmt"

	"github.com/mrlm-net/go-logz/pkg/logger"
	"github.com/mrlm-net/simconnect/pkg/types"
)

// HandleAircraftLoadedFilenameEvent processes AircraftLoaded events received as SIMCONNECT_RECV_EVENT_FILENAME (event ID 101)
func HandleAircraftLoadedFilenameEvent(log *logger.Logger, event *types.SIMCONNECT_RECV_EVENT_FILENAME) {
	filename := ""
	for i, b := range event.SzFileName {
		if b == 0 {
			filename = string(event.SzFileName[:i])
			break
		}
	}
	log.Info(fmt.Sprintf("[SimConnectManager] AircraftLoaded (filename) event received: filename=%q", filename))
}
