package main

import (
	"github.com/polpettone/adventure/engine"
	"github.com/polpettone/adventure/game"
)

func main() {
	startPinguinBurgGame()
}

func startPinguinBurgGame() {

	engine := &engine.EngineOne{}
	engine.Setup()

	game := &game.PinguinBurfGame{}
	game.Init(engine)
	game.Run()

}
