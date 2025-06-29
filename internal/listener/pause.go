package listener

import (
	"fmt"

	"github.com/mrlm-net/go-logz/pkg/logger"
	"github.com/mrlm-net/simconnect/pkg/types"
)

// HandlePauseEvent processes Pause system events (event ID 100)
func HandlePauseEvent(log *logger.Logger, event *types.SIMCONNECT_RECV_EVENT) {
	log.Info(fmt.Sprintf("[SimConnectManager] Pause event received: dwData=%d", event.DwData))
}
