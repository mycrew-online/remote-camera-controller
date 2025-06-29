package listener

import (
	"github.com/mrlm-net/go-logz/pkg/logger"
	"github.com/mycrew-online/remote-camera-controller/internal/manager"
)

// HandleOpen processes SIMCONNECT_RECV_ID_OPEN events.
func HandleOpen(log *logger.Logger, mgr *manager.SimConnectManager) {
	log.Info("[SimConnectManager] Simulator loaded successfully.")
	mgr.Client().RequestSystemState(1, "AircraftLoaded")
	mgr.Client().RequestSystemState(2, "DialogMode")
	mgr.Client().RequestSystemState(3, "FlightLoaded")
	mgr.Client().RequestSystemState(4, "FlightPlan")
	mgr.Client().RequestSystemState(5, "Sim")

	// Subscribe to system events
	mgr.Client().SubscribeToSystemEvent(100, "Pause")
	mgr.Client().SubscribeToSystemEvent(101, "AircraftLoaded")
	mgr.Client().SubscribeToSystemEvent(102, "FlightLoaded")
	mgr.Client().SubscribeToSystemEvent(103, "Crashed")
	mgr.Client().SubscribeToSystemEvent(104, "FlightPlanActivated")
	mgr.Client().SubscribeToSystemEvent(105, "FlightPlanDeactivated")
	mgr.Client().SubscribeToSystemEvent(106, "PositionChanged")
	mgr.Client().SubscribeToSystemEvent(107, "Sim")
	mgr.Client().SubscribeToSystemEvent(108, "View")
}
