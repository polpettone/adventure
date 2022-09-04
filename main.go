package main

import (
	"github.com/polpettone/adventure/collectBallonGame"
	"github.com/polpettone/adventure/engine"
	"github.com/polpettone/adventure/game"
	"github.com/polpettone/adventure/menu"
	"github.com/polpettone/adventure/pinguinBurfGame"
)

func main() {
	gameSelection()
}

func gameSelection() {

	engine := &engine.EngineOne{}
	engine.Setup()

	collectBallonGame := &collectBallonGame.CollectBallonsGame{}
	collectBallonGame.Init(engine)

	pinguinBarfGame := &pinguinBurfGame.PinguinBurfGame{}
	pinguinBarfGame.Init(engine)

	games := []game.Game{
		collectBallonGame,
		pinguinBarfGame,
	}

	menu := menu.NewGameSelectionMenu(engine, games)
	menu.Run()

}
