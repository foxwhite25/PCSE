package main

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/storage"
	"io/ioutil"
	"os"
)

func LoadDialog(win fyne.Window) {
	d := dialog.NewFileOpen(func(f fyne.URIReadCloser, err error) {
		if err != nil {
			dialog.ShowError(err, win)
			return
		}
		if f == nil {
			if State.Save.MetaData.SavePool == 0 {
				LoadDialog(win)
				dialog.ShowInformation("Cannot Cancel", "You need to select a save", win)
			}
			return
		}

		b, err := ioutil.ReadFile(f.URI().Path())
		if err != nil {
			dialog.ShowError(err, win)
			return
		}
		decode, err := Decode(b)
		if err != nil {
			dialog.ShowError(err, win)
			return
		}
		State.Save = decode
		for _, binding := range Bindings {
			err = binding.Reload()
			if err != nil {
				dialog.ShowError(err, win)
			}
		}
		if decode.MetaData.SavePool == AutoSave {
			dialog.ShowInformation("Auto Save Decoded", "Save Time: "+decode.MetaData.SaveTime, win)
		} else {
			dialog.ShowInformation("Save Decoded", "Save Name: "+decode.MetaData.CustomName.CustomName, win)
		}
	}, win)
	home, _ := os.UserHomeDir()
	uri := storage.NewFileURI(home + "\\AppData\\LocalLow\\niceplay games\\Potion Craft\\Saves")
	lister, err := storage.ListerForURI(uri)
	if err != nil {
		dialog.ShowError(err, win)
		return
	}
	d.SetLocation(lister)
	d.SetFilter(storage.NewExtensionFileFilter([]string{".pcsave"}))
	d.Resize(fyne.NewSize(1080, 720))
	d.Show()
}

func SaveDialog(win fyne.Window) {
	d := dialog.NewFileSave(func(f fyne.URIWriteCloser, err error) {
		if err != nil {
			dialog.ShowError(err, win)
			return
		}
		if f == nil {
			return
		}
		b, err := State.Save.Encode()
		if err != nil {
			dialog.ShowError(err, win)
			return
		}
		err = ioutil.WriteFile(f.URI().Path(), b, 777)
		if err != nil {
			dialog.ShowError(err, win)
			return
		}

		dialog.ShowInformation("Save Encoded", "Path: "+f.URI().Path(), win)
	}, win)
	home, _ := os.UserHomeDir()
	uri := storage.NewFileURI(home + "\\AppData\\LocalLow\\niceplay games\\Potion Craft\\Saves")
	lister, err := storage.ListerForURI(uri)
	if err != nil {
		dialog.ShowError(err, win)
		return
	}
	d.SetLocation(lister)
	d.SetFilter(storage.NewExtensionFileFilter([]string{".pcsave"}))
	d.Resize(fyne.NewSize(1080, 720))
	d.Show()
}
