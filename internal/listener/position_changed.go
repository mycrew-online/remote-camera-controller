package listener

import (
	"fmt"

	"github.com/mrlm-net/go-logz/pkg/logger"
	"github.com/mrlm-net/simconnect/pkg/types"
)

// HandlePositionChangedEvent processes PositionChanged system events (event ID 106)
func HandlePositionChangedEvent(log *logger.Logger, event *types.SIMCONNECT_RECV_EVENT) {
	log.Info(fmt.Sprintf("[SimConnectManager] PositionChanged event received: dwData=%d", event.DwData))
}
