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

	// Start the web server for SPA
	go func() {
		if err := app.Server.Run(":8080"); err != nil {
			panic(err)
		}
	}()

	// Bootstrap the application, which initializes the SimConnectManager
	// and starts the main connection loop.
	app.Bootstrap()
}
