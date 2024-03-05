package main

import (
	"fmt"

	fyne "fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/widget"
)

const minWidth, minHeight = 640, 480 

func main() {
	a := app.New()
	w := a.NewWindow("Image weight reducer")

	maxWeightLabel := widget.NewLabel("Max weight in MB (float value)")
    //header := container.NewCenter(headerLabel)

	msgLabel := widget.NewLabel("")

	

	maxWeightInput := widget.NewEntry()
	maxWeightInput.SetPlaceHolder("Max weight in MB (float value)")
	maxWeightInput.SetText("1.5")

	dirLabel := widget.NewLabel("Selected directory:")

	w.SetContent(container.NewVBox(

		maxWeightLabel,
		maxWeightInput,
		widget.NewLabel(""),

		widget.NewButton("Select Directory", func() {
			dialog.ShowFolderOpen(func(path fyne.ListableURI, err error) { //     <--- ok for folder selection
			// dialog.ShowFileOpen(func(path fyne.URIReadCloser, err error) { //  <--- ok for file selection
				if err != nil {
					fmt.Println("Error:", err)
					return
				}
				fmt.Printf("selectedDir = %s\n", path.Path()) 
				dirLabel.SetText(fmt.Sprintf("Selected directory: %s", path.Path()))
			}, w)
		}),
		dirLabel,
		widget.NewLabel(""),
		
		msgLabel,
		widget.NewLabel(""),

		widget.NewButton("Go!", func() {
			msgLabel.SetText("TODO")
		}),
	))

	// Enforce minimum size on resize
    w.SetFixedSize(false)
	w.Resize(fyne.NewSize(minWidth, minHeight)) 

	w.ShowAndRun()
}


/*
import (
	"fmt"
	"strconv"

	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/widget"
)

func main() {
	a := app.New()
	w := a.NewWindow("Form Example")

	// Directory selection
	selectedDir := ""
	dirButton := widget.NewButton("Select Directory", func() {
		dialog.ShowFileOpen(a.Driver().MainView, func(path string, err error) {
			if err != nil {
				dialog.ShowError(a.Driver().MainView, "Error", err.Error())
				return
			}
			selectedDir = path
			dirButton.SetText(fmt.Sprintf("Selected: %s", path))
		})
	})

	// Float input
	floatInput := widget.NewEntry()
	floatInput.SetPlaceHolder("Enter a float value")

	// Submit button
	submitButton := widget.NewButton("Submit", func() {
		// Access user input
		dir := selectedDir
		value, err := strconv.ParseFloat(floatInput.Text(), 64)
		if err != nil {
			dialog.ShowError(a.Driver().MainView, "Error", "Invalid float input")
			return
		}

		// Process form data (replace with your logic)
		fmt.Println("Selected directory:", dir)
		fmt.Println("Entered float value:", value)

		// Clear form (optional)
		selectedDir = ""
		dirButton.SetText("Select Directory")
		floatInput.SetText("")
	})

	// Layout the UI
	form := container.NewVBox(dirButton, floatInput, submitButton)
	w.SetContent(form)

	w.ShowAndRun()
}
*/