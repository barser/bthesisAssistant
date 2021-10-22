package main

import (
	"bthesisAssistant/internal/notification"
	"bthesisAssistant/internal/tray"
	"fmt"
	"log"
	"os"
	"time"
)

var Version = "unversioned"

func main() {
	cc := make(chan int64, 1)
	tc := make(chan time.Time, 1)
	go func(tc <-chan time.Time, cc chan<- int64) {
		notification.Inform("Привет, мир!")
		var i int64 = 0
		for {
			select {
			case <-tc:
				notification.Inform("Пришло событие от таймера")
				i++
				cc <- i
			}
		}
	}(tc, cc)
	tc <- time.Now()

	_, err := fmt.Fprintln(os.Stdout, "Version:\t", Version)
	if err != nil {
		log.Println(err)
	}

	go tray.Run()

	for {
		select {
		case v := <-cc:
			if v > 10 {
				os.Exit(0)
			} else {
				time.Sleep(2 * time.Second)
				tc <- time.Now()
			}
		}
	}
}
