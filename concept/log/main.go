package main

/*
	Package log implements a simple logging package. It defines a type, Logger, with methods for formatting output.
	It also has a predefined 'standard' Logger accessible through helper functions Print[f|ln], Fatal[f|ln],
	and Panic[f|ln], which are easier to use than creating a Logger manually. That logger writes to standard error
	and prints the date and time of each logged message. Every log message is output on a separate line:
	if the message being printed does not end in a newline, the logger will add one. The Fatal functions call os.Exit(1)
	after writing the log message. The Panic functions call panic after writing the log message.
*/
import (
	"bytes"
	"fmt"
	"log"
)

func main() {
	log.Println("file created")

	var buf bytes.Buffer
	logger := log.New(&buf, "logger: ", log.Ldate|log.Ltime|log.Lshortfile)
	logger.Println("file created")
	logger.Println("file updated")
	// logger.Panicln("panic")
	fmt.Print(&buf)
}
