package pages

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
	"net/url"
)

func welcomeScreen(_ fyne.Window) fyne.CanvasObject {
	logo := canvas.NewImageFromURI()
	return container.NewCenter(container.NewVBox(
		widget.NewLabelWithStyle("Welcome to the Potion Craft Save editor", fyne.TextAlignCenter, fyne.TextStyle{Bold: true}),
		widget.NewLabelWithStyle("Be sure to backup your save before editing", fyne.TextAlignCenter, fyne.TextStyle{}),
		container.NewHBox(
			widget.NewHyperlink("github", parseURL("https://github.com/foxwhite25/PCSE")),
		),
		widget.NewLabel(""),
	))
}

func parseURL(urlStr string) *url.URL {
	link, err := url.Parse(urlStr)
	if err != nil {
		fyne.LogError("Could not parse URL", err)
	}

	return link
}
