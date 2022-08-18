package main

import (
	"github.com/polpettone/adventure/game"
)

func main() {
	startPinguinBurgGame()
}

func startPinguinBurgGame() {

	game := &game.PinguinBurfGame{}
	game.Init()
	game.Run()

}
