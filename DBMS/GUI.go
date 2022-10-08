package main

import (
	"DBMS/database"
	"DBMS/utils"
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
	"log"
)

var mainWindow fyne.Window
var newColumns int64 = 1
var rowCreate *fyne.Container
var layoutCreate *fyne.Container
var rows []*fyne.Container

func startGUI() {
	myApp := app.New()
	mainWindow = myApp.NewWindow("Database Management System")
	mainWindow.Resize(fyne.Size{
		Width:  1000,
		Height: 420,
	})
	mainWindow.SetContent(newBoxLayout())

	mainWindow.ShowAndRun()
}

func newBoxLayout() *fyne.Container {
	butt := widget.NewLabel("HUY")
	top := container.New(layout.NewHBoxLayout(), butt)

	c := container.New(layout.NewBorderLayout(top, top, nil, nil),
		widget.NewButton("Amogus", func() {}))

	return c
}

func newRowCreate() *fyne.Container {
	entry := widget.NewEntry()
	entry.SetPlaceHolder("Column name")
	entry.Wrapping = fyne.TextWrapOff

	dropDown := widget.NewSelect([]string{
		database.TypeInteger{}.TypeName(),
		database.TypeReal{}.TypeName(),
		database.TypeChar{}.TypeName(),
		database.TypeString{}.TypeName(),
		database.TypeHTML{}.TypeName(),
		database.TypeStringRange{}.TypeName(),
	}, func(s string) {})

	return container.New(layout.NewGridLayoutWithColumns(2),
		entry,
		dropDown,
	)
}

func setCurrentLayout(layout *fyne.Container) {
	mainWindow.SetContent(layout)
}

func newCreateTableLayout() *container.Scroll {
	//button := widget.NewButton("dd", func() {})
	//box := container.NewBorder(button, button, button, button, widget.NewButton("Create table", func() {}))

	createTableButton := widget.NewButton("Create table", func() {

	})

	createRowButton := widget.NewButtonWithIcon("", theme.ContentAddIcon(), func() {
		layoutCreate.Objects = append(layoutCreate.Objects, newRowCreate())
		layoutCreate.Objects[len(layoutCreate.Objects)-1], layoutCreate.Objects[len(layoutCreate.Objects)-2] =
			layoutCreate.Objects[len(layoutCreate.Objects)-2], layoutCreate.Objects[len(layoutCreate.Objects)-1]
	})

	deleteButton := widget.NewButtonWithIcon("", theme.DeleteIcon(), func() {
		if len(layoutCreate.Objects) > 3 {
			layoutCreate.Objects = utils.RemoveIndex(layoutCreate.Objects, len(layoutCreate.Objects)-2)
		}
	})

	layoutCreate = container.New(layout.NewVBoxLayout(),
		createTableButton,
		newRowCreate(),
		container.New(layout.NewGridLayoutWithColumns(2),
			createRowButton,
			deleteButton,
		),
	)

	scrollContainer := container.NewVScroll(layoutCreate)
	return scrollContainer
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
