package main

import "github.com/mycrew-online/remote-camera-controller/internal"

func main() {
	// Create a new application instance.
	app := internal.NewApplication()
	// Bootstrap the application, which initializes the SimConnectManager
	// and starts the main connection loop.
	app.Bootstrap()
}
