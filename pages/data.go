package pages

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

func FillPages() {
	Pages = []Page{
		{
			Title: "Intro",
			View:  welcomeScreen,
			Icon:  theme.HomeIcon(),
		},
	}
}
