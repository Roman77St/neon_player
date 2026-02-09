package ui

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
)

type PlayerApp struct {
	App    fyne.App
	Window fyne.Window
}

// NewPlayer создает и настраивает экземпляр приложения
func NewPlayer(title string) *PlayerApp {
	myApp := app.NewWithID("player")
	myApp.Settings().SetTheme(&NeonTheme{})

	window := myApp.NewWindow(title)
	window.Resize(fyne.NewSize(MainWindowWidth, MainWindowHeight))
	p := &PlayerApp{
		App:    myApp,
		Window: window,
	}

	p.setupContent()

	return p
}

func (p *PlayerApp) setupContent() {
	mainContent := NewMainScreen(p.Window)
	p.Window.SetContent(mainContent)
}

func (p *PlayerApp) Run() {
	p.Window.ShowAndRun()
}

func NewMainScreen(window fyne.Window) fyne.CanvasObject {
	trackList := container.NewVBox()

	// Кнопка добавления трека
	addTrackBtn := widget.NewButton("Добавить трек", func() {
		dialog.ShowFileOpen(func(reader fyne.URIReadCloser, err error) {
			if err == nil && reader != nil {
				newRow := CreateNewTrackRow(reader.URI().Path(), reader.URI().Name(), trackList)
				trackList.Add(newRow)
			}
		}, window)
	})

	// Боковая панель иконок
	menuIcons := container.NewVBox(
		widget.NewIcon(theme.SettingsIcon()),
		widget.NewIcon(theme.MenuIcon()),
		widget.NewIcon(theme.HelpIcon()),
		widget.NewIcon(theme.ListIcon()),
	)

	// Изображение обложки
	img := canvas.NewImageFromFile(MainImagePath)
	img.FillMode = canvas.ImageFillContain
	img.SetMinSize(fyne.NewSize(MainImageWidth, MainImageHeight))

	leftSection := container.NewHBox(
		container.NewPadded(menuIcons),
		img,
	)

	// Скролл для списка треков
	scroll := container.NewVScroll(trackList)

	// Итоговая сборка через Border
	return container.NewBorder(addTrackBtn, nil, leftSection, nil, scroll)
}
