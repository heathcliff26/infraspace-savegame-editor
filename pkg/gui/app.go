package gui

import (
	"fmt"
	"image/color"
	"os"
	"path/filepath"
	"slices"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/data/binding"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/storage"
	"fyne.io/fyne/v2/widget"
	"github.com/heathcliff26/infraspace-savegame-editor/pkg/save"
)

var (
	TEXT_COLOR   = color.White
	BORDER_COLOR = color.White
)

const (
	ENTRY_WIDTH = 120
	BORDER_SIZE = 1
)

type GUI struct {
	App                 fyne.App
	Main                fyne.Window
	Version             Version
	Menu                *fyne.MainMenu
	Save                *save.Savegame
	Backup              bool
	Resources           []Resource
	Research            []NamedCheckbox
	UnlockAllResearch   *widget.Check
	UnlockResearchQueue *widget.Button
	SpaceshipParts      []NamedCheckbox
	OtherOptions        OtherOptions
	ActionButtons       []*widget.Button
}

func New() *GUI {
	a := app.New()
	version := NewVersionFromApp(a)
	main := a.NewWindow(version.Name)
	g := &GUI{
		App:     a,
		Main:    main,
		Version: version,
		Backup:  true,
	}

	g.Main.SetMainMenu(g.makeMenu())

	resourcesBox := g.makeResourcesBox()
	researchBox := g.makeResearchBox()
	spaceshipBox := g.makeSpaceshipBox()
	optionsBox := g.makeOptionsBox()
	actionButtons := g.makeActionButtons()
	g.Main.SetContent(container.NewVBox(resourcesBox, researchBox, spaceshipBox, optionsBox, actionButtons))
	g.Main.SetFixedSize(true)

	g.Main.Show()

	return g
}

func (g *GUI) Run() {
	g.App.Run()
}

func (g *GUI) DisplayError(err error) {
	dialog.ShowError(err, g.Main)
}

func (g *GUI) loadSavegame(uri fyne.URIReadCloser, err error) {
	if err != nil {
		g.DisplayError(err)
		return
	}
	if uri == nil {
		return
	}

	path := uri.URI().Path()
	g.Save, err = save.LoadSavegame(path)
	if err != nil {
		g.DisplayError(err)
		return
	}

	g.ReloadFromSave()
	fmt.Println("Successfully loaded save file: " + path)

	newTitle := g.Version.Name + " - " + filepath.Base(path)
	fmt.Println("Setting title to: " + newTitle) // Leaving this here for debug, since it keeps panicking here
	g.Main.SetTitle(newTitle)
	for _, b := range g.ActionButtons {
		b.Enable()
	}
}

func (g *GUI) writeSavegame() {
	abortDialog := func(err error) {
		msg := fmt.Sprintf("Encountered an error while saving: %v. No changes have been written to the save-file", err)
		dialog.NewInformation("Error", msg, g.Main).Show()
	}

	for _, resource := range g.Resources {
		value, err := resource.Value.Get()
		if err != nil {
			abortDialog(err)
			return
		}
		err = g.Save.SetResource(resource.Name, value)
		if err != nil {
			abortDialog(err)
			return
		}
	}

	if g.UnlockAllResearch.Checked {
		g.Save.UnlockAllResearch()
	} else {
		for _, research := range g.Research {
			if research.Checkbox.Checked {
				g.Save.UnlockResearch(research.Name)
			} else {
				g.Save.LockResearch(research.Name)
			}
		}
	}

	for _, part := range g.SpaceshipParts {
		if part.Checkbox.Checked {
			g.Save.RepairSpaceshipPart(part.Name)
		}
	}

	starterWorkerCount, err := g.OtherOptions.StarterWorker.Value.Get()
	if err != nil {
		abortDialog(err)
		return
	}
	if starterWorkerCount > g.Save.GetStarterWorkerCount() {
		g.Save.AddStarterWorkers(starterWorkerCount)
	}

	buildingOptions := save.EditBuildingsOptions{
		HabitatWorkers: g.OtherOptions.HabitatWorkers.Checked,
		HabitatStorage: g.OtherOptions.HabitatStorage.Checked,
		FactoryStorage: g.OtherOptions.FactoryStorage.Checked,
		UpgradesOnly:   g.OtherOptions.UpgradesOnly.Checked,
	}
	g.Save.EditBuildings(buildingOptions)

	if g.Save.Changed {
		if g.Backup {
			path, err := g.Save.Backup()
			if err != nil {
				abortDialog(err)
				return
			}
			dialog.NewInformation("Created Backup", "Created backup of save at "+path, g.Main).Show()
		}

		err = g.Save.Save()
		if err != nil {
			abortDialog(err)
			return
		}
	} else {
		dialog.NewInformation("Info", "Please make some changes first.", g.Main).Show()
	}
}

func (g *GUI) ReloadFromSave() {
	for _, resource := range g.Resources {
		value, ok := g.Save.GetResource(resource.Name)
		if !ok {
			g.DisplayError(fmt.Errorf("unkown resource name: %s", resource.Name))
		}
		err := resource.Value.Set(value)
		if err != nil {
			g.DisplayError(err)
			return
		}
		resource.Entry.Refresh()
	}

	unlockedResearch := g.Save.GetUnlockedResearch()
	for _, item := range g.Research {
		item.Checkbox.Enable()
		item.Checkbox.Checked = slices.Contains(unlockedResearch, item.Name)
		item.Checkbox.Refresh()
	}
	g.UnlockAllResearch.Checked = false
	g.UnlockAllResearch.Refresh()
	g.UnlockResearchQueue.Enable()

	repairedSpaceshipParts := g.Save.GetRepairedSpaceshipParts()
	for _, item := range g.SpaceshipParts {
		item.Checkbox.Checked = slices.Contains(repairedSpaceshipParts, item.Name)
		if item.Checkbox.Checked {
			item.Checkbox.Disable()
		} else {
			item.Checkbox.Enable()
		}
		item.Checkbox.Refresh()
	}

	err := g.OtherOptions.StarterWorker.Value.Set(g.Save.GetStarterWorkerCount())
	if err != nil {
		g.DisplayError(err)
		return
	}
	g.OtherOptions.StarterWorker.Entry.Refresh()
}

func (g *GUI) makeMenu() *fyne.MainMenu {
	loadSavegame := func() {
		dir, err := save.DefaultSaveLocation()
		if !RELEASE {
			// When developing, you likely have a copy of the save in the current directory
			dir, err = os.Getwd()
		}
		if err != nil {
			g.DisplayError(err)
			return
		}

		uri, err := storage.ParseURI("file://" + dir)
		if err != nil {
			g.DisplayError(fmt.Errorf("failed to create URI from Path \"%s\": %v", dir, err))
			return
		}

		listURI, err := storage.ListerForURI(uri)
		if err != nil {
			g.DisplayError(err)
			return
		}

		dialog := dialog.NewFileOpen(g.loadSavegame, g.Main)
		dialog.SetLocation(listURI)
		dialog.SetFilter(storage.NewExtensionFileFilter([]string{".sav"}))
		dialog.Resize(fyne.NewSize(800, 600))
		dialog.Show()
	}
	openSave := fyne.NewMenuItem("Load Save", loadSavegame)

	backup := fyne.NewMenuItem("Backup", nil)
	backup.Checked = g.Backup
	backup.Action = func() {
		backup.Checked = !backup.Checked
		g.Backup = backup.Checked
		g.Menu.Refresh()
	}

	fileMenu := fyne.NewMenu("File", openSave, fyne.NewMenuItemSeparator(), backup)

	about := fyne.NewMenuItem("About", nil)
	about.Action = func() {
		vInfo := dialog.NewCustom(g.Version.Name, "close", g.Version.CreateContent(), g.Main)
		vInfo.Show()
	}

	helpMenu := fyne.NewMenu("Help", about)

	g.Menu = fyne.NewMainMenu(fileMenu, helpMenu)
	return g.Menu
}

type Resource struct {
	Name  string
	Value binding.Int
	Entry *widget.Entry
}

func (g *GUI) makeResourcesBox() fyne.CanvasObject {
	resources := []Resource{{Name: "concrete"}, {Name: "steel"}, {Name: "car"}, {Name: "adamantine"}}

	content := make([]fyne.CanvasObject, len(resources))
	for i := 0; i < len(resources); i++ {
		label := canvas.NewText(resources[i].Name+": ", TEXT_COLOR)
		resources[i].Value = binding.NewInt()
		resources[i].Entry = widget.NewEntryWithData(binding.IntToString(resources[i].Value))
		r := resources[i]
		button := widget.NewButton("1 mio.", func() {
			err := r.Value.Set(1000000)
			if err != nil {
				dialog.NewError(err, g.Main).Show()
				return
			}
			r.Entry.Refresh()
		})
		size := resources[i].Entry.MinSize()
		size.Width = ENTRY_WIDTH
		wrappedEntry := container.NewGridWrap(size, resources[i].Entry)
		content[i] = container.NewHBox(label, wrappedEntry, button)
	}

	g.Resources = resources
	return newBorder(container.NewGridWithColumns(len(content), content...))
}

type NamedCheckbox struct {
	Name     string
	Checkbox *widget.Check
}

func (g *GUI) makeResearchBox() fyne.CanvasObject {
	items, widgets := createNamedCheckboxes(save.ResearchNames())
	g.Research = items

	rows := make([]fyne.CanvasObject, 0, (len(widgets)/10)+1)
	for i := 0; i < len(widgets); {
		row := make([]fyne.CanvasObject, 0, 10)
		for x := 0; x < 10; x++ {
			if i < len(widgets) {
				row = append(row, widgets[i])
				i++
			} else {
				row = append(row, layout.NewSpacer())
				break
			}
		}
		rows = append(rows, container.NewVBox(row...))
	}
	researchGrid := newBorder(container.NewHBox(rows...))

	g.UnlockAllResearch = widget.NewCheck("Unlock all Research", func(checked bool) {
		for _, item := range g.Research {
			if checked {
				item.Checkbox.Disable()
			} else {
				item.Checkbox.Enable()
			}
		}
	})

	g.UnlockResearchQueue = widget.NewButton("Unlock Queue", func() {
		queue := g.Save.GetResearchQueue()
		for _, research := range g.Research {
			if slices.Contains(queue, research.Name) {
				research.Checkbox.Checked = true
				research.Checkbox.Refresh()
			}
		}
	})
	g.UnlockResearchQueue.Disable()

	actions := container.NewHBox(g.UnlockAllResearch, g.UnlockResearchQueue)

	return newBorder(container.NewVBox(actions, researchGrid))
}

func (g *GUI) makeSpaceshipBox() fyne.CanvasObject {
	items, widgets := createNamedCheckboxes(save.SpaceshipParts())
	g.SpaceshipParts = items

	return newBorder(container.NewGridWithColumns(7, widgets...))
}

type OtherOptions struct {
	StarterWorker struct {
		Value binding.Int
		Entry *widget.Entry
	}
	HabitatWorkers *widget.Check
	HabitatStorage *widget.Check
	FactoryStorage *widget.Check
	UpgradesOnly   *widget.Check
}

func (g *GUI) makeOptionsBox() fyne.CanvasObject {
	g.OtherOptions = OtherOptions{
		HabitatWorkers: widget.NewCheck("Fill all habitats with workers", nil),
		HabitatStorage: widget.NewCheck("Fill the storage of all habitats", nil),
		FactoryStorage: widget.NewCheck("Fill the storage of all factories", nil),
		UpgradesOnly:   widget.NewCheck("Fill only upgrades factories", nil),
	}

	g.OtherOptions.FactoryStorage.OnChanged = func(b bool) {
		if b {
			g.OtherOptions.UpgradesOnly.Checked = b
			g.OtherOptions.UpgradesOnly.Disable()
		} else {
			g.OtherOptions.UpgradesOnly.Enable()
		}
		g.OtherOptions.UpgradesOnly.Refresh()
	}

	g.OtherOptions.StarterWorker.Value = binding.NewInt()
	g.OtherOptions.StarterWorker.Entry = widget.NewEntryWithData(binding.IntToString(g.OtherOptions.StarterWorker.Value))
	size := g.OtherOptions.StarterWorker.Entry.MinSize()
	size.Width = ENTRY_WIDTH
	wrappedStarterWorkerEntry := container.NewGridWrap(size, g.OtherOptions.StarterWorker.Entry)
	starterWorkerLabel := canvas.NewText("Increase starter worker count: ", TEXT_COLOR)
	starterWorkerBox := container.NewHBox(starterWorkerLabel, wrappedStarterWorkerEntry)

	checkboxes := container.NewGridWithColumns(4, g.OtherOptions.HabitatWorkers, g.OtherOptions.HabitatStorage, g.OtherOptions.FactoryStorage, g.OtherOptions.UpgradesOnly)

	return newBorder(container.NewVBox(starterWorkerBox, checkboxes))
}

func (g *GUI) makeActionButtons() fyne.CanvasObject {
	resetWarning := func() {
		dialog.NewConfirm("Warning", "This will reload all values from the save", func(b bool) {
			if b {
				g.ReloadFromSave()
			}
		}, g.Main).Show()
	}
	reset := widget.NewButton("Reset", resetWarning)
	reset.Disable()
	saveFile := widget.NewButton("Save", g.writeSavegame)
	saveFile.Disable()

	g.ActionButtons = []*widget.Button{reset, saveFile}

	return container.NewPadded(container.NewCenter(container.NewHBox(reset, saveFile)))
}
