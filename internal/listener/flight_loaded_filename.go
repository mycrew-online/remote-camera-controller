package listener

import (
	"fmt"

	"github.com/mrlm-net/go-logz/pkg/logger"
	"github.com/mrlm-net/simconnect/pkg/types"
)

// HandleFlightLoadedFilenameEvent processes FlightLoaded events received as SIMCONNECT_RECV_EVENT_FILENAME (event ID 102)
func HandleFlightLoadedFilenameEvent(log *logger.Logger, event *types.SIMCONNECT_RECV_EVENT_FILENAME) {
	filename := ""
	for i, b := range event.SzFileName {
		if b == 0 {
			filename = string(event.SzFileName[:i])
			break
		}
	}
	log.Info(fmt.Sprintf("[SimConnectManager] FlightLoaded (filename) event received: filename=%q", filename))
}
