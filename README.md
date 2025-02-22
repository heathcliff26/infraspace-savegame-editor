[![CI](https://github.com/heathcliff26/infraspace-savegame-editor/actions/workflows/ci.yaml/badge.svg?event=push)](https://github.com/heathcliff26/infraspace-savegame-editor/actions/workflows/ci.yaml)
[![Coverage Status](https://coveralls.io/repos/github/heathcliff26/infraspace-savegame-editor/badge.svg)](https://coveralls.io/github/heathcliff26/infraspace-savegame-editor)
[![Editorconfig Check](https://github.com/heathcliff26/infraspace-savegame-editor/actions/workflows/editorconfig-check.yaml/badge.svg?event=push)](https://github.com/heathcliff26/infraspace-savegame-editor/actions/workflows/editorconfig-check.yaml)
[![Generate go test cover report](https://github.com/heathcliff26/infraspace-savegame-editor/actions/workflows/go-testcover-report.yaml/badge.svg)](https://github.com/heathcliff26/infraspace-savegame-editor/actions/workflows/go-testcover-report.yaml)
[![Renovate](https://github.com/heathcliff26/infraspace-savegame-editor/actions/workflows/renovate.yaml/badge.svg)](https://github.com/heathcliff26/infraspace-savegame-editor/actions/workflows/renovate.yaml)

# InfraSpace Savegame Editor

This is a savegame editor for InfraSpace. It is written in golang with the App framework from fyne.io.

## Usage

1. Select `File -> Load Save` to load a savegame
2. Edit the save how you want to
3. Save the savegame by pressing `Save`

Notes:
- You can reset all changes by pressing reset. This will cause the GUI to reset all selections made
- If you do not want to create a backup every time you save, unselect `File -> Backup`

## Images

### Main Window

![](images/dark/MainWindow.png#gh-dark-mode-only)
![](images/light/MainWindow.png#gh-light-mode-only)
![](images/dark/SaveEditing.png#gh-dark-mode-only)
![](images/light/SaveEditing.png#gh-light-mode-only)

### File Menu

![](images/dark/FileMenu.png#gh-dark-mode-only)
![](images/light/FileMenu.png#gh-light-mode-only)
![](images/dark/FileDialog.png#gh-dark-mode-only)
![](images/light/FileDialog.png#gh-light-mode-only)
