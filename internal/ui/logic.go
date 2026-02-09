package ui

import (
	"fmt"
	"time"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/data/binding"
	"fyne.io/fyne/v2/widget"
	"github.com/Roman77St/playsound"
)

// handlePlayTap управляет началом воспроизведения и запуском симуляции прогресса
func handlePlayTap(path string, done *chan struct{}, vol *widget.Slider, btnText binding.String, timeData binding.String, progData binding.Float, progCont *fyne.Container) {
	if *done == nil {
		d, err := playsound.PlaySound(path)
		if err != nil {
			return
		}
		*done = d
		playsound.SetVolume(*done, vol.Value)
		btnText.Set("Pause")
		progCont.Show()

		go runProgress(d, progData, timeData, btnText, func() {
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

// runProgress опрашивает движок о текущем положении и обновляет UI
func runProgress(c chan struct{}, data binding.Float, timeData binding.String, btnText binding.String, callback func()) {
	dur, errDur := playsound.GetDuration(c)
	if errDur != nil || dur <= 0 {
		dur = 1
	}

	ticker := time.NewTicker(time.Millisecond * ProgressInterval)
	defer ticker.Stop()

	for {
		select {
		case <-c:
			btnText.Set("Play")
			data.Set(0)
			// cont.Hide()
			callback()
			return
		case <-ticker.C:
			pos, err := playsound.GetPosition(c)
			if err == nil && dur > 0 {
				percentage := pos / dur
				data.Set(percentage)
				newTimeStr := fmt.Sprintf("%s / %s", formatTime(pos), formatTime(dur))
				timeData.Set(newTimeStr)
			}
		}
	}
}

// handleSeek перематывает трек на указанный процент (0.0 - 1.0)
func handleSeek(done chan struct{}, percent float64) {
	if done == nil {
		return
	}
	dur, err := playsound.GetDuration(done)
	if err != nil || dur <= 0 {
		return
	}

	// Вычисляем целевое время в секундах
	targetTime := dur * percent
	playsound.Seek(done, targetTime)
}

// Вспомогательная функция для превращения секунд в 00:00
func formatTime(seconds float64) string {
	minutes := int(seconds) / 60
	secs := int(seconds) % 60
	return fmt.Sprintf("%02d:%02d", minutes, secs)
}
