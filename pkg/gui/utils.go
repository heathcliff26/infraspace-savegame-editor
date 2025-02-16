package gui

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

// Create a line for the border
func makeBorderStrip() fyne.CanvasObject {
	rec := canvas.NewRectangle(BORDER_COLOR)
	rec.SetMinSize(fyne.NewSize(BORDER_SIZE, BORDER_SIZE))
	return rec
}

// Wrap the objects in a box with border lines
func newBorder(content ...fyne.CanvasObject) fyne.CanvasObject {
	top := makeBorderStrip()
	left := makeBorderStrip()
	bottom := makeBorderStrip()
	right := makeBorderStrip()
	border := container.NewBorder(top, bottom, left, right, content...)
	return container.NewPadded(border)
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
