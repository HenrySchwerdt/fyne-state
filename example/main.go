package main

import (
	"fmt"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
	"github.com/HenrySchwerdt/fyne_state"
)

func main() {
	a := app.New()
	w := a.NewWindow("Zustand State Management")

	// Create the store
	fyne_state.Create(func(set func(string, interface{}), get func(string) interface{}) fyne_state.StateActions {
		return fyne_state.StateActions{
			State: map[string]interface{}{
				"count": 0,
			},
			Actions: map[string]func(){
				"inc": func() {
					count := get("count").(int)
					set("count", count+1)
				},
			},
		}
	})

	// Counter component
	countLabel := widget.NewLabel(fmt.Sprintf("Count: %d", fyne_state.Get[int]("count")))
	updateCh := fyne_state.Subscribe("count")

	go func() {
		for range updateCh {
			count := fyne_state.Get[int]("count")
			countLabel.SetText(fmt.Sprintf("Count: %d", count))
			countLabel.Refresh()
		}
	}()

	incButton := widget.NewButton("Increment", func() {
		inc := fyne_state.Get[func()]("inc")
		inc()
	})

	content := container.NewVBox(countLabel, incButton)
	w.SetContent(content)
	w.Resize(fyne.NewSize(300, 200))
	w.ShowAndRun()
}
