package main

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/data/binding"
	"fyne.io/fyne/v2/widget"
)

func statsScreen(_ fyne.Window) fyne.CanvasObject {
	return container.NewVBox(
		widget.NewLabelWithStyle("Variables:", fyne.TextAlignCenter, fyne.TextStyle{Bold: true}),
		BindIntWithColumns("Gold:", &State.Save.UserData.Gold),
		BindIntWithColumns("Popularity:", &State.Save.UserData.Popularity),
		BindIntWithColumns("Karma:", &State.Save.UserData.Karma),
		BindIntWithColumns("Current Level:", &State.Save.UserData.CurrentLevel),
		widget.NewSeparator(),
		widget.NewLabelWithStyle("Statistics:", fyne.TextAlignCenter, fyne.TextStyle{Bold: true}),
		BindIntWithColumns("Potion Brewed:", &State.Save.UserData.PotionsBrewed),
		BindIntWithColumns("Legendary Substances Brewed:", &State.Save.UserData.LegendarySubstancesBrewedAmount),
		BindIntWithColumns("Clients Served:", &State.Save.UserData.ClientsServed),
		BindIntWithColumns("Profit From Haggling:", &State.Save.UserData.ProfitFromHaggling),
	)
}

func BindIntWithColumns(label string, k *int) *fyne.Container {
	data := binding.BindInt(k)
	Bindings = append(Bindings, data)
	return container.NewGridWithColumns(
		3, widget.NewLabel(label), widget.NewLabel(""), widget.NewEntryWithData(binding.IntToString(data)),
	)
}
