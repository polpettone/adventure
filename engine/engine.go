package engine

import (
	"fmt"

	"github.com/polpettone/adventure/models"
)

type Engine interface {
	Machine(text string)
}

type SimpleEngine struct {
	GameMap models.Map
	Player  *models.Player
	Player2 *models.Player
}

func (se SimpleEngine) Machine(key string) {
	updatePlayer(key, &se.GameMap, se.Player)
	updatePlayer(key, &se.GameMap, se.Player2)
}

func InitEngine() Engine {

	player := models.NewPlayer(models.PLAYER, 0, 0, "k", "j", "l", "h", "m")
	player2 := models.NewPlayer(models.PLAYER2, 79, 29, "w", "s", "d", "a", "x")

	gameMap := models.NewMap(80, 30)
	gameMap.Place(models.NewField(models.BOX, 30, 5))
	gameMap.Place(models.NewField(models.BOX, 10, 10))
	gameMap.Place(models.NewField(models.BOX, 40, 15))
	gameMap.Place(models.NewField(models.BOX, 55, 20))

	gameMap.Place(player)
	gameMap.Place(player2)

	clearScreen()
	fmt.Println(gameMap.Print())

	var engine Engine = SimpleEngine{
		GameMap: gameMap,
		Player:  player,
		Player2: player2,
	}
	return engine
}

func updatePlayer(key string, gameMap *models.Map, player *models.Player) {
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
		gameMap.Place(models.NewField(models.FIELD, player.X, player.Y))
		player.MoveUp()

	case player.MoveRightKey:
		if player.X == gameMap.MaxX-1 {
			return
		}
		gameMap.Place(models.NewField(models.FIELD, player.X, player.Y))
		player.MoveRight()

	case player.MoveLeftKey:
		if player.X == 0 {
			return
		}
		gameMap.Place(models.NewField(models.FIELD, player.X, player.Y))
		player.MoveLeft()

	case player.MoveDownKey:
		if player.Y == 0 {
			return
		}
		gameMap.Place(models.NewField(models.FIELD, player.X, player.Y))
		player.MoveDown()
	}

	clearScreen()
	gameMap.Place(player)
	fmt.Println(gameMap.Print())
}

func clearScreen() {
	fmt.Print("\033[H\033[2J")
}
