package main

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/theme"
)

var State StateType

func main() {
	State = StateType{}
	a := app.NewWithID("xyz.foxwhite.pcse")
	FillPages()
	a.SetIcon(theme.SettingsIcon())
	logLifecycle(a)
	State.Windows = a.NewWindow("Potion Craft Save Editor")
	State.Windows.SetMainMenu(makeMenu(a, State.Windows))
	State.Windows.SetMaster()
	State.Windows.Resize(fyne.NewSize(1080, 720))
	LoadDialog(State.Windows)
	State.Windows.ShowAndRun()
}
