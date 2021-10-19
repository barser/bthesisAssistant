package main

import (
	"bthesisAssistant/internal/notification"
	"bthesisAssistant/internal/tray"
)

func main() {
	go func() {
		notification.Inform("Привет, мир!")
	}()
	tray.Run()
}
