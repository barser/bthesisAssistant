package tray

import (
	"github.com/getlantern/systray"
)

func Run() {
	systray.Run(onReady, onExit)
}

func onReady() {

	systray.SetTitle("Тезис Помощник 2.0")
	systray.SetTooltip("для ПК Борец")

	menuQuit := systray.AddMenuItem("Выйти", "Выход из приложения")
	setIcons(menuQuit)

	go func() {
		for {
			select {
			case <-menuQuit.ClickedCh:
				systray.Quit()
				return
			}
		}
	}()
}

func onExit() {
	// clean up here
}
