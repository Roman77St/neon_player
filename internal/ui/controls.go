package ui

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/data/binding"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
)

// createPlayButton создает кнопку воспроизведения с динамическим текстом и иконкой
func createPlayButton() (*widget.Button, binding.String) {
	btnText := binding.NewString()
	btnText.Set("Play")

	btn := widget.NewButtonWithIcon("Play", theme.MediaPlayIcon(), nil)

	btnText.AddListener(binding.NewDataListener(func() {
		t, _ := btnText.Get()
		btn.SetText(t)
		if t == "Pause" {
			btn.SetIcon(theme.MediaPauseIcon())
		} else {
			btn.SetIcon(theme.MediaPlayIcon())
		}
	}))
	return btn, btnText
}

// createCustomProgressBar создает тонкую неоновую полоску прогресса
func createCustomProgressBar() (binding.Float, *fyne.Container) {
	progressData := binding.NewFloat()
	progressData.Set(0)

	bg := canvas.NewRectangle(ColorProgressBg)
	bg.SetMinSize(fyne.NewSize(TrackInfoWidth, ProgressBarHeight))

	bar := canvas.NewRectangle(ColorNeonMain)
	bar.SetMinSize(fyne.NewSize(0, ProgressBarHeight))

	progressData.AddListener(binding.NewDataListener(func() {
		val, _ := progressData.Get()
		bar.SetMinSize(fyne.NewSize(ProgressBarWidth * float32(val), ProgressBarHeight))
		bar.Refresh()
	}))

	container := container.NewStack(bg, container.NewHBox(bar))
	container.Hide()
	return progressData, container
}

// createVolumeSlider создает слайдер громкости
func createVolumeSlider() *widget.Slider {
	vol := widget.NewSlider(0, 1)
	vol.SetValue(DefaultVolume)
	vol.Step = VolumeStep
	return vol
}

// createDeleteButton создает кнопку удаления
func createDeleteButton() *widget.Button {
	return widget.NewButtonWithIcon("", theme.DeleteIcon(), nil)
}
