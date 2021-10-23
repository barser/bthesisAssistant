package notification

import (
	"github.com/gen2brain/beeep"
	"log"
)

const title string = "Тезис-Помощник"

// Inform shows notification with specified text in it in the corner of desktop
func Inform(text string) {

	err := beeep.Notify(title, text, "assets/borets.png")
	if err != nil {
		log.Println(err)
	}
}
