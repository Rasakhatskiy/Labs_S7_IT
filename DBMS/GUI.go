package main

import (
	"DBMS/database"
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
	"log"
)

var mainWindow fyne.Window

func startGUI() {
	myApp := app.New()
	mainWindow = myApp.NewWindow("Database Management System")
	mainWindow.Resize(fyne.Size{
		Width:  640,
		Height: 0,
	})

	mainWindow.SetContent(newCreateTableLayout())

	mainWindow.ShowAndRun()
}

func setCurrentLayout(layout *fyne.Container) {
	mainWindow.SetContent(layout)
}

func newCreateTableLayout() *fyne.Container {
	c := container.New(layout.NewVBoxLayout(),
		container.New(layout.NewGridLayoutWithColumns(3),
			container.New(layout.NewHBoxLayout(), widget.NewButtonWithIcon("", theme.DeleteIcon(), func() {})),
			widget.NewEntry(),
			widget.NewSelect([]string{"a", "b", "c"}, func(s string) {}),
		),
		widget.NewButtonWithIcon("", theme.ContentAddIcon(), func() {}),
	)

	return c
}

func newSelectDatabaseCreationLayout() *fyne.Container {
	return container.New(
		layout.NewVBoxLayout(),
		widget.NewButtonWithIcon("Load from file", theme.UploadIcon(), func() {}),
		widget.NewButtonWithIcon("Create new database", theme.ContentAddIcon(), func() {
			setCurrentLayout(newCreateDatabaseLayout())
		}),
	)
}

func newCreateDatabaseLayout() *fyne.Container {
	entry := widget.NewEntry()
	entry.SetPlaceHolder("Your database name")

	form := &widget.Form{
		Items: []*widget.FormItem{{
			Widget:   entry,
			HintText: "нижній текст",
		}},
		OnSubmit: func() {
			log.Println("Form submitted:", entry.Text)
			gCurrentDB = database.CreateDatabase(entry.Text)
		},
	}

	label := widget.NewLabel("Create database")

	return container.New(layout.NewVBoxLayout(), label, form)
}
