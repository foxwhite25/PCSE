package main

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/data/binding"
	"fyne.io/fyne/v2/widget"
)

var InventoryMap []interface{}

func inventoryScreen(_ fyne.Window) fyne.CanvasObject {
	field := container.NewVBox(
		container.NewGridWithColumns(
			3, widget.NewLabel("Count"), widget.NewLabel(""), widget.NewEntry(),
		),
	)
	list := widget.NewListWithData(binding.BindUntypedList(&InventoryMap),
		func() fyne.CanvasObject {
			return widget.NewLabel("item x.y")
		},
		func(item binding.DataItem, obj fyne.CanvasObject) {
			f, _ := item.(binding.Untyped).Get()
			inventory := f.(*Inventory)
			text := obj.(*widget.Label)
			text.Bind(binding.BindString(&inventory.Name))
		})
	return container.NewHSplit(list, field)
}
