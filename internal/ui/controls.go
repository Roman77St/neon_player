package ui

import (
	"image/color"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/data/binding"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
)

// Тип для кликабельной полоски
type clickableBar struct {
	widget.BaseWidget
	onSeek func(float64) // callback, который вернет процент клика
}

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
func createCustomProgressBar() (binding.Float, binding.String, *fyne.Container) {
	progressData := binding.NewFloat()
	progressData.Set(0)

	timeData := binding.NewString()
	timeData.Set("00:00 / 00:00")

	bg := canvas.NewRectangle(ColorProgressBg)
	bg.SetMinSize(fyne.NewSize(TrackInfoWidth, ProgressBarHeight))

	bar := canvas.NewRectangle(ColorNeonMain)
	bar.SetMinSize(fyne.NewSize(0, ProgressBarHeight))

	progressData.AddListener(binding.NewDataListener(func() {
		val, _ := progressData.Get()
		bar.SetMinSize(fyne.NewSize(ProgressBarWidth*float32(val), ProgressBarHeight))
		bar.Refresh()
	}))

	container := container.NewStack(bg, container.NewHBox(bar))
	return progressData, timeData, container
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

// Клик по полоске бара
func (b *clickableBar) Tapped(e *fyne.PointEvent) {
	// e.Position.X — координата клика
	// b.Size().Width — общая ширина полоски
	w := b.Size().Width
	if w <= 0 {
		return // Защита от деления на ноль
	}
	percent := float64(e.Position.X / b.Size().Width)
	if percent < 0 {
		percent = 0
	}
	if percent > 0.99 {
		percent = 0.99
	}
	b.onSeek(percent)
}

// Нужно для реализации интерфейса Tappable
func (b *clickableBar) CreateRenderer() fyne.WidgetRenderer {
	rect := canvas.NewRectangle(color.Transparent)
	return widget.NewSimpleRenderer(rect)
}
