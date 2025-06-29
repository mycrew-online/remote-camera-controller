package listener

import (
	"fmt"

	"github.com/mrlm-net/go-logz/pkg/logger"
	"github.com/mrlm-net/simconnect/pkg/types"
	"github.com/mycrew-online/remote-camera-controller/internal/manager"
)

// StartSimConnectEventListener starts the event handler goroutine for SimConnect events.
func StartSimConnectEventListener(mgr *manager.SimConnectManager, log *logger.Logger) {
	go func() {
		for event := range mgr.Stream() {
			log.Debug(fmt.Sprintf("[SimConnectManager] Received event: %v, data type: %T", event.MessageType, event.Data))
			if !mgr.IsOnline() {
				return
			}

			switch event.MessageType {
			case types.SIMCONNECT_RECV_ID_OPEN:
				HandleOpen(log, mgr)

			case types.SIMCONNECT_RECV_ID_SYSTEM_STATE:
				if ev, ok := event.Data.(*types.SIMCONNECT_RECV_SYSTEM_STATE); ok {
					HandleSystemState(log, mgr, ev)
				}
			case types.SIMCONNECT_RECV_ID_EVENT:
				if ev, ok := event.Data.(*types.SIMCONNECT_RECV_EVENT); ok {
					switch ev.UEventID {
					case 100:
						HandlePauseEvent(log, mgr, ev)
					case 101:
						HandleAircraftLoadedEvent(log, mgr, ev)
					case 102:
						HandleFlightLoadedEvent(log, mgr, ev)
					case 103:
						HandleCrashedEvent(log, mgr, ev)
					case 104:
						HandleFlightPlanActivatedEvent(log, mgr, ev)
					case 105:
						HandleFlightPlanDeactivatedEvent(log, mgr, ev)
					case 106:
						HandlePositionChangedEvent(log, mgr, ev)
					case 107:
						HandleSimEvent(log, mgr, ev)
					case 108:
						HandleViewEvent(log, mgr, ev)
					}
				}
			case types.SIMCONNECT_RECV_ID_EVENT_FILENAME:
				if ev, ok := event.Data.(*types.SIMCONNECT_RECV_EVENT_FILENAME); ok {
					switch ev.UEventID {
					case 101:
						HandleAircraftLoadedFilenameEvent(log, ev)
					case 102:
						HandleFlightLoadedFilenameEvent(log, ev)
					}
				}
			case types.SIMCONNECT_RECV_ID_QUIT:
				HandleQuit(log, mgr)
			default:
				log.Info(fmt.Sprintf("[SimConnectManager] Unhandled event type: %v, data: %+v", event.MessageType, event.Data))
			}
		}
	}()
}
