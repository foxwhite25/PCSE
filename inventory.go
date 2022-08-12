package main

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/data/validation"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
	"golang.org/x/exp/slices"
	"strconv"
)

func inventoryScreen(_ fyne.Window) fyne.CanvasObject {
	tabs := container.NewDocTabs()
	for _, inventory := range State.Save.UserData.PlayerInventory {
		println(inventory.Name)
		t := newItemTab(inventory)
		println("append tab")
		tabs.Append(t)
	}

	tabs.SetTabLocation(container.TabLocationTop)
	tabs.OnClosed = func(item *container.TabItem) {
		i := slices.IndexFunc(State.Save.UserData.PlayerInventory, func(i *Inventory) bool { return i.Name == item.Text })
		State.Save.UserData.PlayerInventory = append(State.Save.UserData.PlayerInventory[:i], State.Save.UserData.PlayerInventory[i+1:]...)
		println(State.Save.UserData.PlayerInventory)
	}
	tabs.CreateTab = func() *container.TabItem {
		newItemForm(tabs)
		return nil
	}
	return tabs
}

func newItemForm(tabs *container.DocTabs) {
	name := widget.NewEntry()
	count := widget.NewEntry()
	classFullName := widget.NewEntry()
	classFullName.SetPlaceHolder("PotionCraft.ScriptableObjects.Ingredient.Ingredient")
	data := widget.NewMultiLineEntry()
	count.Validator = validation.NewRegexp(`^[+-]?[\d]+$`, "Count can only be number")
	items := []*widget.FormItem{
		{
			Text:     "Name",
			Widget:   name,
			HintText: "Internal Name for item",
		}, {
			Text:     "Count",
			Widget:   count,
			HintText: "Number of item",
		}, {
			Text:     "Class Full Name",
			Widget:   classFullName,
			HintText: "Internal Name for class, default for ingredients",
		}, {
			Text:     "Extra Data",
			Widget:   data,
			HintText: "Extra Data for item, usually in potions",
		},
	}

	form := dialog.NewForm("Create New Item", "Confirm", "Cancel", items, func(b bool) {
		if !b {
			return
		}
		i, _ := strconv.Atoi(count.Text)
		inventory := &Inventory{
			Name:          name.Text,
			Count:         i,
			ClassFullName: classFullName.Text,
			Data:          data.Text,
		}
		State.Save.UserData.PlayerInventory = append(State.Save.UserData.PlayerInventory, inventory)
		tabs.Append(newItemTab(inventory))
	}, State.Windows)
	form.Resize(fyne.Size{
		Width:  720,
		Height: 405,
	})
	form.Show()
}

func newItemTab(i *Inventory) *container.TabItem {
	println("Create tab")
	c := container.NewVBox(
		BindIntWithColumns("Count", &i.Count),
		BindStringWithSelection("Class Full Name", &i.ClassFullName, []string{"PotionCraft.ScriptableObjects.Ingredient.Ingredient", "PotionCraft.ScriptableObjects.Potion"}),
		BindStringWithMultiLine("Data", &i.Data),
	)
	println("create container")
	return container.NewTabItemWithIcon(i.Name, theme.MenuIcon(), c)

}
