package ui

import (
	"time"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/data/binding"
	"fyne.io/fyne/v2/widget"
	"github.com/Roman77St/playsound"
)

// handlePlayTap управляет началом воспроизведения и запуском симуляции прогресса
func handlePlayTap(path string, done *chan struct{}, vol *widget.Slider, btnText binding.String, progData binding.Float, progCont *fyne.Container) {
	if *done == nil {
		d, err := playsound.PlaySound(path)
		if err != nil {
			return
		}
		*done = d
		playsound.SetVolume(*done, vol.Value)
		btnText.Set("Pause")
		progCont.Show()

		go runProgressSimulation(d, progData, progCont, btnText, func() {
			*done = nil
		})
	} else {
		handlePlaybackToggle(*done, btnText, vol.Value)
	}
}

// handlePlaybackToggle отвечает за переключение между паузой и возобновлением
func handlePlaybackToggle(done chan struct{}, btnText binding.String, currentVol float64) {
	t, _ := btnText.Get()
	if t == "Pause" {
		playsound.Pause(done)
		btnText.Set("Resume")
	} else {
		playsound.PlayOn(done)
		btnText.Set("Pause")
		playsound.SetVolume(done, currentVol)
	}
}

// runProgressSimulation имитирует движение полоски прогресса
func runProgressSimulation(c chan struct{}, data binding.Float, cont *fyne.Container, btnText binding.String, collBack func()) {
	ticker := time.NewTicker(time.Millisecond * ProgressInterval)
	defer ticker.Stop()
	var current float32 = 0

	for {
		select {
		case <-c:
			btnText.Set("Play")
			data.Set(0)
			cont.Hide()
			collBack()
			return
		case <-ticker.C:
			if current < 1.0 {
				current += ProgressStep
				data.Set(float64(current))
			}
		}
	}
}