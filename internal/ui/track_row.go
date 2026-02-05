package ui

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
	"github.com/Roman77St/playsound"
)

// Функция-конструктор для отдельной дорожки
func CreateNewTrackRow(path, name string, trackList *fyne.Container) fyne.CanvasObject {
	var done chan struct{}
	var wholeRow fyne.CanvasObject

	btnPlay, btnText := createPlayButton()
	progressData, progressContainer := createCustomProgressBar()

	vol := createVolumeSlider()

	vol.OnChanged = func(f float64) {
		if done != nil {
			playsound.SetVolume(done, f)
		}
	}

	btnPlay.OnTapped = func() {
		handlePlayTap(path, &done, vol, btnText, progressData, progressContainer)
	}

	deleteBtn := createDeleteButton()

	wholeRow = assembleTrackRowLayout(name, btnPlay, progressContainer, vol, deleteBtn)

	deleteBtn.OnTapped = func() {
		if done != nil { playsound.Stop(done) }
		trackList.Remove(wholeRow)
	}

	return wholeRow
}

func assembleTrackRowLayout(name string, btn *widget.Button, progress *fyne.Container, vol *widget.Slider, del *widget.Button) fyne.CanvasObject {
	bg := canvas.NewRectangle(ColorRowBg)
	bg.CornerRadius = RowCornerRadius
	bg.StrokeColor = ColorRowStroke
	bg.StrokeWidth = RowStrokeWidth

	trackInfo := container.NewVBox(widget.NewLabel(name), progress)
	infoContainer := container.NewGridWrap(fyne.NewSize(TrackInfoWidth, TrackInfoHeight), trackInfo)

	row := container.NewHBox(
		container.NewGridWrap(fyne.NewSize(PlayBtnWidth, ControlHeight), btn),
		infoContainer,
		layout.NewSpacer(),
		widget.NewIcon(theme.VolumeUpIcon()),
		container.NewGridWrap(fyne.NewSize(VolumeSliderWidth, ControlHeight), vol),
		container.NewGridWrap(fyne.NewSize(DeleteBtnWidth, ControlHeight), del),
	)

	return container.NewStack(bg, container.NewPadded(row))
}
