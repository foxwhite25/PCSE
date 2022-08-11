package pages

import (
	"PotionCraftSaveEditor/res"
	"bytes"
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
	"net/url"
)

func welcomeScreen(_ fyne.Window) fyne.CanvasObject {
	logo := canvas.NewImageFromReader(bytes.NewReader(res.Logo), "logo.png")
	logo.FillMode = canvas.ImageFillContain
	if fyne.CurrentDevice().IsMobile() {
		logo.SetMinSize(fyne.NewSize(192, 192))
	} else {
		logo.SetMinSize(fyne.NewSize(256, 256))
	}
	return container.NewCenter(container.NewVBox(
		logo,
		widget.NewLabelWithStyle("Welcome to the Potion Craft Save editor", fyne.TextAlignCenter, fyne.TextStyle{Bold: true}),
		widget.NewLabelWithStyle("Be sure to backup your save before editing", fyne.TextAlignCenter, fyne.TextStyle{}),
		widget.NewHyperlinkWithStyle("Github", parseURL("https://github.com/foxwhite25/PCSE"), fyne.TextAlignCenter, fyne.TextStyle{}),
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
