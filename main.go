package main

import (
	"image/color"
	"time"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/data/binding"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
	"github.com/Roman77St/playsound"
)

func main() {
	app := app.NewWithID("player")
	app.Settings().SetTheme(&NeonTheme{})
	window := app.NewWindow("Audio Player")
	window.Resize(fyne.NewSize(950, 400))

	// Главный контейнер, куда будут добавляться треки
	trackList := container.NewVBox()

	// Кнопка для добавления нового трека
	addTrackBtn := widget.NewButton("Добавить трек", func() {
		dialog.ShowFileOpen(func(reader fyne.URIReadCloser, err error) {
			if err == nil && reader != nil {
				// Создаем панель управления для этого конкретного файла
				newRow := createNewTrackRow(reader.URI().Path(), reader.URI().Name(), trackList)
				trackList.Add(newRow)
			}
		}, window)
	})

	menuIcons := container.NewVBox(
		widget.NewIcon(theme.SettingsIcon()),
		widget.NewIcon(theme.MenuIcon()),
		widget.NewIcon(theme.HelpIcon()),
		widget.NewIcon(theme.ListIcon()),
	)

	img := canvas.NewImageFromFile("img/song1.png")
	img.FillMode = canvas.ImageFillContain
	img.SetMinSize(fyne.NewSize(350, 500))

	leftSection := container.NewHBox(
		container.NewPadded(menuIcons),
		img,
	)

	// Оборачиваем список в Scroll для прокрутки, если треков много
	scroll := container.NewVScroll(trackList)

	content := container.NewBorder(addTrackBtn, nil, leftSection, nil, scroll)

	window.SetContent(content)
	window.ShowAndRun()
}

// Функция-конструктор для отдельной дорожки
func createNewTrackRow(path, name string, trackList *fyne.Container) fyne.CanvasObject {
	var done chan struct{}
	var wholeRow fyne.CanvasObject

	btnText := binding.NewString()
	btnText.Set("Play")

	progressData := binding.NewFloat()
	progressData.Set(0)

	playIcon := theme.MediaPlayIcon()
	pauseIcon := theme.MediaPauseIcon()
	btn := widget.NewButtonWithIcon("Play", playIcon, nil)

	btnText.AddListener(binding.NewDataListener(func() {
		t, _ := btnText.Get()
		btn.SetText(t)
		if t == "Pause" {
			btn.SetIcon(pauseIcon)
			} else {
				btn.SetIcon(playIcon)
			}
		}))

		// Фон полоски (темный)
		progressBg := canvas.NewRectangle(color.NRGBA{R: 40, G: 40, B: 45, A: 255})
		progressBg.SetMinSize(fyne.NewSize(200, 2)) // Устанавливаем высоту 2 пикселя

		// Активная часть (неоновая)
		progressBar := canvas.NewRectangle(color.NRGBA{R: 0, G: 255, B: 200, A: 255})
		progressBar.SetMinSize(fyne.NewSize(0, 2)) // Начальная ширина 0

		// Контейнер, который накладывает их друг на друга
		// Оборачиваем progressBar в HBox, чтобы он рос слева направо
		progressContainer := container.NewStack(progressBg, container.NewHBox(progressBar))
		progressContainer.Hide()

	progressData.AddListener(binding.NewDataListener(func() {
		val, _ := progressData.Get()
		progressBar.SetMinSize(fyne.NewSize(200 * float32(val), 2))
		progressBar.Refresh()
	}))


	vol := widget.NewSlider(0, 1)
	vol.SetValue(1)
	vol.Step = 0.01
	vol.OnChanged = func(f float64) {
		if done != nil {
			playsound.SetVolume(done, vol.Value)
		}
	}

	btn.OnTapped = func() {
		if done == nil {
			d, err := playsound.PlaySound(path)
			if err != nil {
				return
			}
			done = d
			playsound.SetVolume(done, vol.Value)
			btnText.Set("Pause")
			progressContainer.Show()

			go func(c chan struct{}) {

// ========================================================================
//                           Это имитация!
				ticker := time.NewTicker(time.Millisecond * 500)
				defer ticker.Stop()

				var currentProgress float32 = 0

				for {
					select {
					case <-c: // Канал закрылся (песня кончилась)
						done = nil
						btnText.Set("Play")
						progressData.Set(float64(currentProgress))
						progressContainer.Hide()
						return
					case <-ticker.C:
						// Имитируем движение: добавляем по чуть-чуть
						// Если песня длинная, он дойдет до конца и замрет
						if currentProgress < 1.0 {
							currentProgress += 0.01
							// 200 - это ширина фона, заданная выше
							progressData.Set(float64(currentProgress))
						}
					}
// =========================================================================
				}
			}(done)
		} else {
			t, _ := btnText.Get()
			if t == "Pause" {
				playsound.Pause(done)
				btnText.Set("Resume")
			} else {
				playsound.PlayOn(done)
				btnText.Set("Pause")
				playsound.SetVolume(done, vol.Value)
			}
		}
	}

	deleteBtn := widget.NewButtonWithIcon("", theme.DeleteIcon(), func() {
		// Если музыка играет, останавливаем её перед удалением
		if done != nil {
			playsound.Stop(done) // Или ваша функция для полной остановки
		}
		// Удаляем этот контейнер из списка треков
		trackList.Remove(wholeRow)
	})

	bg := canvas.NewRectangle(color.NRGBA{R: 25, G: 25, B: 30, A: 255})
	bg.CornerRadius = 10
	bg.StrokeColor = color.NRGBA{R: 0, G: 255, B: 200, A: 100} // Полупрозрачный неон
	bg.StrokeWidth = 1

	btnContainer := container.NewGridWrap(fyne.NewSize(120, 40), btn)
	delContainer := container.NewGridWrap(fyne.NewSize(40, 40), deleteBtn)
	trackInfo := container.NewVBox(widget.NewLabel(name), progressContainer)
	infoContainer := container.NewGridWrap(fyne.NewSize(200, 45), trackInfo)

	// строка управления
	row := container.NewHBox(
		btnContainer,
		infoContainer,
		layout.NewSpacer(),
		widget.NewIcon(theme.VolumeUpIcon()),
		container.NewGridWrap(fyne.NewSize(150, 40), vol),
		delContainer,
	)

	wholeRow = container.NewStack(bg, container.NewPadded(row))

	return wholeRow
}

type NeonTheme struct{}

func (m NeonTheme) Color(name fyne.ThemeColorName, variant fyne.ThemeVariant) color.Color {
	neonColor := color.NRGBA{R: 0, G: 255, B: 200, A: 255}
	switch name {
	case theme.ColorNameBackground:
		return color.NRGBA{R: 5, G: 5, B: 7, A: 255} // Очень темный фон для контраста
	case theme.ColorNamePrimary:
		return neonColor
	case theme.ColorNameButton:
		return color.NRGBA{R: 25, G: 25, B: 35, A: 255} // Темные кнопки, чтобы выделялся текст
	case theme.ColorNameForeground:
		return neonColor
	case theme.ColorNameInputBackground:
		return color.NRGBA{R: 20, G: 20, B: 25, A: 255}
	}
	return theme.DefaultTheme().Color(name, theme.VariantDark)
}

func (m NeonTheme) Icon(name fyne.ThemeIconName) fyne.Resource {
	return theme.DefaultTheme().Icon(name)
}

func (m NeonTheme) Font(style fyne.TextStyle) fyne.Resource {
	return theme.DefaultTheme().Font(style)
}

func (m NeonTheme) Size(name fyne.ThemeSizeName) float32 {
	return theme.DefaultTheme().Size(name)
}
