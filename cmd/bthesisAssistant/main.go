package main

import (
	"bthesisAssistant/internal/notification"
	"bthesisAssistant/internal/tray"
	"fmt"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
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

	breakFor := false
	for !breakFor {
		select {
		case v := <-cc:
			if v > 2 {
				//os.Exit(0)
				breakFor = true
			} else {
				time.Sleep(2 * time.Second)
				tc <- time.Now()
			}
		}
	}

	a := app.New()
	w := a.NewWindow("Hello")

	hello := widget.NewLabel("Hello Fyne!")
	w.SetContent(container.NewVBox(
		hello,
		widget.NewButton("Hi!", func() {
			hello.SetText("Welcome :)")
		}),
	))
	w.ShowAndRun()
}
