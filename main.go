package main

import (
	"github.com/Roman77St/player/internal/ui"
)

func main() {
	player := ui.NewPlayer("Audio Player")

	player.Run()
}
