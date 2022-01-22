package domain

import (
	"os"
	"os/signal"
	"time"
)

// Application encapsulates a running application
type Application interface {
	// Start starts the app
	Start()
	// Should clean up and stop
	Stop()
}

// RunApplication starts the application and waits for a stop signal to stop the application
func RunApplication(application Application, gracePeriod time.Duration, stopSignals ...os.Signal) {
	go application.Start()
	c := make(chan os.Signal, 1)
	signal.Notify(c, stopSignals...)
	// Wait for a stop signal from the user or the system
	<-c
	go application.Stop()
	// Allow the application to end gracefully
	<-time.NewTimer(gracePeriod).C
}
