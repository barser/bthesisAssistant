//go:build linux
// +build linux

package tray

import (
	"github.com/getlantern/systray"
	"io/ioutil"
	"log"
)

func setIcons(menuItem *systray.MenuItem) {

	if icon, err := ioutil.ReadFile("assets/trayIcon.ico"); err != nil {
		log.Println(err)
	} else {
		systray.SetIcon(icon)
		menuItem.SetIcon(icon)
	}
}
