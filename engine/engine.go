package engine

import (
	"fmt"

	"github.com/polpettone/adventure/models"
)

type Engine interface {
	Machine(text string)
}

type SimpleEngine struct {
	GameMap *models.Map
	Player1 *models.Player
	Player2 *models.Player
}

func (se SimpleEngine) Machine(key string) {
	updatePlayer(key, se.Player1, se.GameMap)
	updatePlayer(key, se.Player2, se.GameMap)

	clearScreen()
	se.GameMap.UpdatePlayer1(*se.Player1)
	se.GameMap.UpdatePlayer2(*se.Player2)

	fmt.Println(se.GameMap.Print())
}

func InitEngine() Engine {

	player1 := models.NewPlayer(models.PLAYER, 0, 0, "k", "j", "l", "h", "m")
	player2 := models.NewPlayer(models.PLAYER2, 79, 29, "w", "s", "d", "a", "x")

	gameMap := models.NewMap(80, 30, *player1, *player2)
	gameMap.Place(models.NewField(models.BOX, 30, 5))
	gameMap.Place(models.NewField(models.BOX, 10, 10))
	gameMap.Place(models.NewField(models.BOX, 40, 15))
	gameMap.Place(models.NewField(models.BOX, 55, 20))

	gameMap.Place(player1)
	gameMap.Place(player2)

	clearScreen()
	fmt.Println(gameMap.Print())

	var engine Engine = SimpleEngine{
		GameMap: gameMap,
		Player1: player1,
		Player2: player2,
	}
	return engine
}

func updatePlayer(key string, player *models.Player, gameMap *models.Map) {
	switch key {

	case player.ActionKey:
		if player.X == gameMap.MaxX-1 {
			return
		}
		gameMap.Place(models.NewItem(models.PENGUIN, player.X, player.Y))
		player.X += 1

	case player.MoveUpKey:
		if player.Y == gameMap.MaxY-1 {
			return
		}
		player.MoveUp()

	case player.MoveRightKey:
		if player.X == gameMap.MaxX-1 {
			return
		}
		player.MoveRight()

	case player.MoveLeftKey:
		if player.X == 0 {
			return
		}
		player.MoveLeft()

	case player.MoveDownKey:
		if player.Y == 0 {
			return
		}
		player.MoveDown()
	}
}

func clearScreen() {
	fmt.Print("\033[H\033[2J")
}
