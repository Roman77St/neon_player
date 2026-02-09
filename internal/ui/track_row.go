package ui

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
	"github.com/Roman77St/playsound"
)

// Функция-конструктор для отдельной дорожки
func CreateNewTrackRow(path, name string, trackList *fyne.Container) fyne.CanvasObject {
	var done chan struct{}
	var wholeRow fyne.CanvasObject

	btnPlay, btnText := createPlayButton()
	progressData, timeData, progressContainer := createCustomProgressBar()
	timeLabel := widget.NewLabelWithData(timeData)
	timeLabel.Alignment = fyne.TextAlignTrailing

	// Создаем "невидимый" слой для кликов поверх прогресс-бара
	seekOverlay := &clickableBar{
		onSeek: func(p float64) {
			if done != nil {
				handleSeek(done, p)
				progressData.Set(p) // Сразу визуально двигаем бар
			}
		},
	}
	seekOverlay.ExtendBaseWidget(seekOverlay)

	// Накладываем кликабельный слой поверх полоски прогресса
	clickableProgress := container.NewStack(progressContainer, seekOverlay)

	vol := createVolumeSlider()

	vol.OnChanged = func(f float64) {
		if done != nil {
			playsound.SetVolume(done, f)
		}
	}

	btnPlay.OnTapped = func() {
		handlePlayTap(path, &done, vol, btnText, timeData, progressData, clickableProgress)
	}

	deleteBtn := createDeleteButton()

	wholeRow = assembleTrackRowLayout(name, btnPlay, clickableProgress, timeLabel, vol, deleteBtn)

	deleteBtn.OnTapped = func() {
		if done != nil {
			playsound.Stop(done)
		}
		trackList.Remove(wholeRow)
	}

	return wholeRow
}

func assembleTrackRowLayout(name string, btn *widget.Button, progress *fyne.Container, timeLabel fyne.CanvasObject, vol *widget.Slider, del *widget.Button) fyne.CanvasObject {
	// 1. Левый блок: Название сверху, Бар снизу
	title := widget.NewLabelWithStyle(name, fyne.TextAlignLeading, fyne.TextStyle{Bold: true})
	title.Truncation = fyne.TextTruncateEllipsis
	fixedProgress := container.NewGridWrap(fyne.NewSize(ProgressBarWidth, ProgressBarHeight), progress)
	leftStack := container.NewVBox(title, fixedProgress)

	// 2. Правый блок: Громкость сверху, Время снизу
	if l, ok := timeLabel.(*widget.Label); ok {
		l.Alignment = fyne.TextAlignCenter
	}
	rightStack := container.NewVBox(vol, timeLabel)

	// 3. Центральная часть (Инфо + Управление)
	centerContent := container.NewBorder(nil, nil, nil, rightStack, leftStack)

	fixedPlayBtn := container.NewGridWrap(fyne.NewSize(PlayBtnWidth, PlayBtnHeight), btn)

	// 4. Итоговая сборка всей строки
	content := container.NewBorder(
		nil,
		nil,
		container.NewPadded(fixedPlayBtn),  // Слева (Play)
		del,                                // Справа (Удалить)
		container.NewPadded(centerContent), // Центр (Название, Бар, Громкость, Время)
	)

	// 5. Фоновая подложка
	bg := canvas.NewRectangle(ColorRowBg)
	bg.CornerRadius = 10

	mainStack := container.NewStack(bg, content)

	return container.NewPadded(mainStack)
}
