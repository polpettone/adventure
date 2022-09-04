package collectBallonGame

import (
	"github.com/polpettone/adventure/game"
	"github.com/polpettone/adventure/models"
)

var gameConfig game.GameConfig = game.GameConfig{

	MapSize:       game.NewCoord(30, 30),
	ItemCount:     1,
	InitPlayerPos: game.NewCoord(0, 0),

	PlayerControlMap: game.ControlMap{
		Up:     "k",
		Down:   "j",
		Left:   "h",
		Right:  "l",
		Action: "m",
	},

	ItemSymbol: models.BALLON,
}
