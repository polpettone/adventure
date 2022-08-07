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
	Enemies []*models.Enemy
}

func (se SimpleEngine) Machine(key string) {
	updatePlayer(key, se.Player1, se.GameMap)
	updatePlayer(key, se.Player2, se.GameMap)

	clearScreen()

	se.GameMap.Update()
	fmt.Println(se.GameMap.Print())
}

func InitEngine() Engine {

	player1 := models.NewPlayer(models.PLAYER, 0, 0, "k", "j", "l", "h", "m")
	player2 := models.NewPlayer(models.PLAYER2, 79, 29, "w", "s", "d", "a", "x")

	enemy := models.NewEnemy(models.PENGUIN, 3, 3)
	enemies := []*models.Enemy{enemy}

	mapElements := []models.MapElement{
		models.NewItem(models.BOX, 30, 5),
		models.NewItem(models.BOX, 10, 10),
		models.NewItem(models.BOX, 40, 15),
		models.NewItem(models.BOX, 55, 20),
		player1,
		player2,
		enemy,
	}

	gameMap := models.NewMap(80, 30, mapElements)

	clearScreen()
	fmt.Println(gameMap.Print())

	var engine Engine = SimpleEngine{
		GameMap: gameMap,
		Player1: player1,
		Player2: player2,
		Enemies: enemies,
	}
	return engine
}

func updatePlayer(key string, player *models.Player, gameMap *models.Map) {
	switch key {

	case player.ActionKey:
		if player.X == gameMap.MaxX-1 {
			return
		}
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

	mapElement := gameMap.GetElementFromPos(player.X, player.Y)

	t := fmt.Sprintf("%T", mapElement)

	gameMap.SetStatusLine(0, string(mapElement.GetSymbol())+" "+t)

}

func clearScreen() {
	fmt.Print("\033[H\033[2J")
}
