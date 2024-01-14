package gui

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
	"github.com/heathcliff26/infraspace-savegame-editor/pkg/version"
)

// Create the content for the version dialog
func getVersionContent() fyne.CanvasObject {
	v := version.Version()

	r1 := make([]fyne.CanvasObject, 4)
	r2 := make([]fyne.CanvasObject, 4)
	r1[0] = canvas.NewText("Version:", TEXT_COLOR)
	r2[0] = canvas.NewText(v.Version, TEXT_COLOR)
	r1[1] = canvas.NewText("Commit:", TEXT_COLOR)
	r2[1] = canvas.NewText(v.Commit, TEXT_COLOR)
	r1[2] = canvas.NewText("Go:", TEXT_COLOR)
	r2[2] = canvas.NewText(v.Go, TEXT_COLOR)
	r1[3] = canvas.NewText("InfraSpace:", TEXT_COLOR)
	r2[3] = canvas.NewText(v.GameVersion, TEXT_COLOR)

	row1 := container.NewVBox(r1...)
	row2 := container.NewVBox(r2...)

	return container.NewPadded(container.NewHBox(row1, row2))
}

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
