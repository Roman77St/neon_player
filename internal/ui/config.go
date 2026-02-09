package ui

import (
	"image/color"
)

// Настройки интерфейса
const (
	// Основной экран
	MainWindowWidth  = 950
	MainWindowHeight = 500
	MainImagePath    = "assets/song1.png"
	MainImageWidth   = 350
	MainImageHeight  = 500

	// Размеры элементов в TrackRow
	TrackInfoWidth  = 200
	TrackInfoHeight = 80
	PlayBtnWidth    = 100
	PlayBtnHeight   = TrackInfoHeight

	// Параметры прогресс-бара
	ProgressBarHeight = 10
	ProgressBarWidth  = 200
	ProgressInterval  = 100 // Мс между обновлениями

	// Параметры звука
	VolumeStep    = 0.01
	DefaultVolume = 1.0
)

// Цветовая палитра Neon Dark
var (
	ColorNeonMain   = color.NRGBA{R: 0, G: 255, B: 200, A: 255}
	ColorBackground = color.NRGBA{R: 5, G: 5, B: 7, A: 255}
	ColorRowBg      = color.NRGBA{R: 25, G: 25, B: 30, A: 255}
	ColorProgressBg = color.NRGBA{R: 40, G: 40, B: 45, A: 255}
	ColorButtonBg   = color.NRGBA{R: 25, G: 25, B: 35, A: 255}
)
