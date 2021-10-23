package main

import (
	"bthesisAssistant/internal/notification"
	"bthesisAssistant/internal/tray"
	"fmt"
	"github.com/gen2brain/dlgs"
	"log"
	"os"
	"time"
)

// Version contains version which could be modified while building via -ldflags parameter
var Version = "unversioned"

func main() {
	cc := make(chan int64, 1)
	tc := make(chan time.Time, 1)

	item, _, err := dlgs.List("List", "Select item from list:", []string{"Bug", "New Feature", "Improvement"})
	if err != nil {
		panic(err)
	}

	go func(tc <-chan time.Time, cc chan<- int64) {
		notification.Inform("Привет, мир!")
		var i int64 = 0
		for {
			select {
			case <-tc:
				notification.Inform(item)
				i++
				cc <- i
			}
		}
	}(tc, cc)
	tc <- time.Now()

	_, err = fmt.Fprintln(os.Stdout, "Version:\t", Version)
	if err != nil {
		log.Println(err)
	}

	go tray.Run()

	for {
		select {
		case v := <-cc:
			if v > 2 {
				os.Exit(0)
			} else {
				time.Sleep(2 * time.Second)
				tc <- time.Now()
			}
		}
	}
}
