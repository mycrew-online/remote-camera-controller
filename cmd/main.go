package main

import (
	"flag"

	"github.com/mycrew-online/remote-camera-controller/internal"
)

func main() {
	verbose := flag.Bool("verbose", false, "Enable debug logging")
	flag.Parse()

	logLevel := "info"
	if *verbose {
		logLevel = "debug"
	}

	// Create a new application instance with log level option.
	app := internal.NewApplicationWithOptions(logLevel)
	// Bootstrap the application, which initializes the SimConnectManager
	// and starts the main connection loop.
	app.Bootstrap()
}
