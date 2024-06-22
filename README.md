# Fyne-State

## Example

```go
package main

func main() {
	a := app.New()
	w := a.NewWindow("FyneState State Management")

	// Create the store
	fyne_state.Create(func(set func(string, interface{}), get func(string) interface{}) fyne_state.StateActions {
		return StateActions{
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
	countLabel := widget.NewLabel(fmt.Sprintf("Count: %d", Get[int]("count")))
	updateCh := Subscribe("count")

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

```