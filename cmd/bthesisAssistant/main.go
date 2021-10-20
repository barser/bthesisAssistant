package main

import (
	"bthesisAssistant/internal/notification"
	"bthesisAssistant/internal/tray"
	"fmt"
	"log"
	"os"
)

var Version = "unversioned"

func main() {
	go func() {
		notification.Inform("Привет, мир!")
	}()

	_, err := fmt.Fprintln(os.Stdout, "Version:\t", Version)
	if err != nil {
		log.Println(err)
	}

	tray.Run()
}
