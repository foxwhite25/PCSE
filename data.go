package main

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/theme"
)

type Page struct {
	Title string
	View  func(w fyne.Window) fyne.CanvasObject
	Icon  fyne.Resource
}

var Pages []Page
var Bindings []interface {
	Reload() error
}

func FillPages() {
	Pages = []Page{
		{
			Title: "Intro",
			View:  welcomeScreen,
			Icon:  theme.HomeIcon(),
		},
		{
			Title: "Player Stats",
			View:  statsScreen,
			Icon:  theme.AccountIcon(),
		},
		{
			Title: "Inventory",
			View:  inventoryScreen,
			Icon:  theme.FolderOpenIcon(),
		},
	}
}
