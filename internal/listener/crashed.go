package listener

import (
	"fmt"

	"github.com/mrlm-net/go-logz/pkg/logger"
	"github.com/mrlm-net/simconnect/pkg/types"
)

// HandleCrashedEvent processes Crashed system events (event ID 103)
func HandleCrashedEvent(log *logger.Logger, event *types.SIMCONNECT_RECV_EVENT) {
	log.Info(fmt.Sprintf("[SimConnectManager] Crashed event received: dwData=%d", event.DwData))
}
