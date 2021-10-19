//go:build windows
// +build windows

package tray

import (
	"github.com/getlantern/systray"
	"github.com/getlantern/systray/example/icon"
)

func setIcons(menuItem *systray.MenuItem) {
	systray.SetIcon(icon.Data)
	menuItem.SetIcon(icon.Data)
}
