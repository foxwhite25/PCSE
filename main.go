package main

import (
	"PotionCraftSaveEditor/pages"
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/theme"
)

var state State

func main() {
	state = State{}
	a := app.NewWithID("xyz.foxwhite.pcse")
	pages.FillPages()
	a.SetIcon(theme.SettingsIcon())
	logLifecycle(a)
	state.Windows = a.NewWindow("Potion Craft Save Editor")
	state.Windows.SetMainMenu(makeMenu(a, state.Windows))
	state.Windows.SetMaster()
	state.Windows.Resize(fyne.NewSize(1080, 720))

	tabs := container.NewAppTabs()

	for _, page := range pages.Pages {
		tabs.Append(container.NewTabItemWithIcon(page.Title, page.Icon, page.View(state.Windows)))
	}

	tabs.SetTabLocation(container.TabLocationLeading)

	state.Windows.SetContent(tabs)
	SaveDialog(state.Windows)
	state.Windows.ShowAndRun()
}
