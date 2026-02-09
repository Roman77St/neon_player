package ui

import (
	"image/color"
)

// Настройки интерфейса
const (
	// Основной экран
	MainWindowWidth  = 950
	MainWindowHeight = 400
	MainImagePath    = "img/song1.png"
	MainImageSize    = 350

	// Размеры элементов в TrackRow
	TrackRowHeight     = 60
	TrackInfoWidth     = 200
	TrackInfoHeight    = 80
	PlayBtnWidth       = 100
	PlayBtnHeight      = TrackInfoHeight
	VolumeSliderWidth  = 150
	DeleteBtnWidth     = 40
	ControlHeight      = 40

	// Параметры прогресс-бара
	ProgressBarHeight  = 10
	ProgressBarWidth   = 200
	ProgressStep       = 0.01 // Шаг обновления (1%)
	ProgressInterval   = 100  // Мс между обновлениями

	// Параметры звука
	VolumeStep         = 0.01
	DefaultVolume      = 1.0
)

// Цветовая палитра Neon Dark
var (
	ColorNeonMain       = color.NRGBA{R: 0, G: 255, B: 200, A: 255}
	ColorBackground     = color.NRGBA{R: 5, G: 5, B: 7, A: 255}
	ColorRowBg          = color.NRGBA{R: 25, G: 25, B: 30, A: 255}
	ColorProgressBg     = color.NRGBA{R: 40, G: 40, B: 45, A: 255}
	ColorButtonBg       = color.NRGBA{R: 25, G: 25, B: 35, A: 255}

	// Параметры рамки
	RowStrokeWidth      = float32(1)
	RowCornerRadius     = float32(10)
	ColorRowStroke      = color.NRGBA{R: 0, G: 255, B: 200, A: 100}
)
