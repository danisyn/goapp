package main

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	//"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
	"io/ioutil"
	"os"
	"log"
	"fmt"
)

type FilePath struct {
	name string
	path string
}

func main() {
	a := app.New()
	w := a.NewWindow("GO")

	title := widget.NewLabel("Gestor VPN")

		files := func() []string {

			home := os.Getenv("HOME")
			path := home + "/VPN/"
		
			files, err := ioutil.ReadDir(path)
			if err != nil {
				log.Fatal(err)
			}
		
			length := len(files)
		
			arr := make([] string, length)
		
			for i, file := range files {
				
				arr[i] = file.Name()
			}
		
			return arr
	}

	selecty := widget.NewSelect(files(), func(arr string) {
	})

	selecty.PlaceHolder = "Files"

	var file string

	var selection string

	selected := widget.NewLabel(selection)

	btnOK := widget.NewButton("OK", func() { 
		file = selecty.Selected
		fmt.Printf("%s Selected file index=%d\n", selecty.Selected, selecty.SelectedIndex())
		selection = "Selected VPN: " + file
		selected.SetText(selection)
	})

	

	//////////////BOXES/////////////////

	titlebox := container.New(layout.NewHBoxLayout(),title)

	hbox := container.New(layout.NewHBoxLayout() , selecty, btnOK)

	selectbox := container.New(layout.NewHBoxLayout(), selected)

	w.SetContent(container.New(layout.NewVBoxLayout(),titlebox , hbox, selectbox))

	w.Resize(fyne.NewSize(600, 600))

	w.ShowAndRun()
}
