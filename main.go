package main

import (
	"github.com/polpettone/adventure/collectBallonGame"
	"github.com/polpettone/adventure/engine"
	"github.com/polpettone/adventure/pinguinBurfGame"
)

func main() {
	startPinguinBurgGame()
}

func startPinguinBurgGame() {

	engine := &engine.EngineOne{}
	engine.Setup()

	game := &pinguinBurfGame.PinguinBurfGame{}
	game.Init(engine)
	game.Run()

}

func startCollectBallonsGame() {

	engine := &engine.EngineOne{}
	engine.Setup()

	game := &collectBallonGame.CollectBallonsGame{}
	game.Init(engine)
	game.Run()

}
