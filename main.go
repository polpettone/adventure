package main

import (
	"github.com/polpettone/adventure/game"
)

func main() {
	startPinguinBurgGame()
}

func startPinguinBurgGame() {

	for {
		game := &game.PinguinBurfGame{}
		game.Init()
		quit := game.Run()

		if quit {
			return
		}

	}
}
