package main

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/cmd/fyne_settings/settings"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/data/binding"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/driver/desktop"
	"fyne.io/fyne/v2/widget"
	"log"
)

func logLifecycle(a fyne.App) {
	a.Lifecycle().SetOnStarted(func() {
		log.Println("Lifecycle: Started")
	})
	a.Lifecycle().SetOnStopped(func() {
		log.Println("Lifecycle: Stopped")
	})
	a.Lifecycle().SetOnEnteredForeground(func() {
		log.Println("Lifecycle: Entered Foreground")
	})
	a.Lifecycle().SetOnExitedForeground(func() {
		log.Println("Lifecycle: Exited Foreground")
	})
}

func makeMenu(a fyne.App, w fyne.Window) *fyne.MainMenu {
	openSettings := func() {
		w := a.NewWindow("Display Settings")
		w.SetContent(settings.NewSettings().LoadAppearanceScreen(w))
		w.Resize(fyne.NewSize(480, 480))
		w.Show()
	}
	settingsItem := fyne.NewMenuItem("Settings", openSettings)
	settingsShortcut := &desktop.CustomShortcut{KeyName: fyne.KeyComma, Modifier: fyne.KeyModifierShortcutDefault}
	settingsItem.Shortcut = settingsShortcut
	w.Canvas().AddShortcut(settingsShortcut, func(shortcut fyne.Shortcut) {
		openSettings()
	})

	// a quit item will be appended to our first (File) menu
	file := fyne.NewMenu("File", fyne.NewMenuItem("Load", func() {
		dialog.ShowConfirm("Confirm overwrite", "Loading a New save will overwrite current edits!", func(b bool) {
			if b {
				LoadDialog(w)
			}
		}, w)
	}), fyne.NewMenuItem("Save", func() {
		SaveDialog(w)
	}))
	device := fyne.CurrentDevice()
	if !device.IsMobile() && !device.IsBrowser() {
		file.Items = append(file.Items, fyne.NewMenuItemSeparator(), settingsItem)
	}
	return fyne.NewMainMenu(
		file,
	)
}

func BindIntWithColumns(label string, k *int) *fyne.Container {
	println("create" + label)
	data := binding.BindInt(k)
	Bindings = append(Bindings, data)
	println("Bind " + label)
	e := widget.NewEntryWithData(binding.IntToString(data))
	println("entry " + label)
	newGridWithColumns := container.NewGridWithColumns(
		2, widget.NewLabel(label), e,
	)
	println("return " + label)
	return newGridWithColumns
}

func BindStringWithSelection(label string, k *string, selection []string) *fyne.Container {
	println("create " + label)
	data := binding.BindString(k)
	Bindings = append(Bindings, data)
	println("Bind " + label)
	s := widget.NewSelectEntry(selection)
	s.Bind(data)
	println("entry " + label)
	newGridWithColumns := container.NewGridWithColumns(
		2, widget.NewLabel(label), s,
	)
	println("return " + label)
	return newGridWithColumns
}

func BindStringWithMultiLine(label string, k *string) *fyne.Container {
	println("create " + label)
	data := binding.BindString(k)
	Bindings = append(Bindings, data)
	println("Bind " + label)
	m := widget.NewMultiLineEntry()
	m.Bind(data)
	println("entry " + label)
	n := container.NewGridWithColumns(
		2, widget.NewLabel(label), m,
	)
	println("return " + label)
	return n
}
