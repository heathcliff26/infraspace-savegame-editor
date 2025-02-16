package gui

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
)

// Wrap the objects in a box with border lines
func newBorder(content ...fyne.CanvasObject) fyne.CanvasObject {
	contentContainer := container.NewThemeOverride(container.NewPadded(content...), theme.DefaultTheme())
	border := widget.NewCard("", "", contentContainer)

	return container.NewThemeOverride(border, borderTheme{})
}

// Create NamedCheckboxes from a list of names
func createNamedCheckboxes(names []string) ([]NamedCheckbox, []fyne.CanvasObject) {
	items := make([]NamedCheckbox, len(names))
	widgets := make([]fyne.CanvasObject, len(names))
	for i := 0; i < len(names); i++ {
		items[i] = NamedCheckbox{Name: names[i]}
		items[i].Checkbox = widget.NewCheck(names[i], nil)
		widgets[i] = items[i].Checkbox
	}
	return items, widgets
}
