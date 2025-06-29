package listener

import (
	"fmt"

	"github.com/mrlm-net/go-logz/pkg/logger"
	"github.com/mrlm-net/simconnect/pkg/types"
)

// View event data flags (corrected order)
const (
	SIMCONNECT_VIEW_SYSTEM_EVENT_DATA_COCKPIT_VIRTUAL = 2
	SIMCONNECT_VIEW_SYSTEM_EVENT_DATA_COCKPIT_2D      = 1
	SIMCONNECT_VIEW_SYSTEM_EVENT_DATA_ORTHOGONAL      = 4
)

// HandleViewEvent processes View system events (event ID 108)
func HandleViewEvent(log *logger.Logger, event *types.SIMCONNECT_RECV_EVENT) {
	var viewType string
	switch event.DwData {
	case 0:
		viewType = "EXTERNAL"
	case SIMCONNECT_VIEW_SYSTEM_EVENT_DATA_COCKPIT_VIRTUAL:
		viewType = "COCKPIT_VIRTUAL"
	case SIMCONNECT_VIEW_SYSTEM_EVENT_DATA_COCKPIT_2D:
		viewType = "COCKPIT_2D"
	case SIMCONNECT_VIEW_SYSTEM_EVENT_DATA_ORTHOGONAL:
		viewType = "ORTHOGONAL (Map View)"
	default:
		viewType = fmt.Sprintf("Unknown (%d)", event.DwData)
	}
	log.Info(fmt.Sprintf("[SimConnectManager] View event received: dwData=%d (%s)", event.DwData, viewType))
}
