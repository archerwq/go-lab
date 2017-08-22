package syncutil

import (
	"fmt"
	"os"
	"os/signal"
)

// CtrC accept graceful shutdowns when quit via Ctrl+C
func CtrC() {
	c := make(chan os.Signal, 1)
	// We'll accept graceful shutdowns when quit via SIGINT (Ctrl+C)
	// SIGKILL, SIGQUIT or SIGTERM (Ctrl+/) will not be caught.
	signal.Notify(c, os.Interrupt)

	// Block until we receive our signal.
	<-c

	fmt.Println("Shuting down...")
}
