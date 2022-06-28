package engine

import (
	"fmt"

	"github.com/polpettone/adventure/models"
)

type Engine interface {
	Machine(text string) string
}

type SimpleEngine struct {
	GameMap models.Map
	Player  *models.Player
	Player2 *models.Player
}

func (se SimpleEngine) Machine(key string) string {
	updatePlayer(key, &se.GameMap, se.Player)
	updatePlayer(key, &se.GameMap, se.Player2)
	return "nothing happend"
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

func updatePlayer(key string, gameMap *models.Map, player *models.Player) string {
	switch key {

	case player.ActionKey:
		if player.X == gameMap.MaxX-1 {
			return "cant action, there comes the wall"
		}
		clearScreen()
		gameMap.Place(models.NewItem(models.PENGUIN, player.X, player.Y))
		player.X += 1
		gameMap.Place(player)
		fmt.Println(gameMap.Print())
		return "dit action and moved right"

	case player.MoveUpKey:
		if player.Y == gameMap.MaxY-1 {
			return "wall, cant move"
		}
		clearScreen()
		gameMap.Place(models.NewField(models.FIELD, player.X, player.Y))
		player.MoveUp()
		gameMap.Place(player)
		fmt.Println(gameMap.Print())
		return "moved up"

	case player.MoveRightKey:
		if player.X == gameMap.MaxX-1 {
			return "wall, cant move"
		}
		clearScreen()
		gameMap.Place(models.NewField(models.FIELD, player.X, player.Y))
		player.MoveRight()
		gameMap.Place(player)
		fmt.Println(gameMap.Print())
		return "moved right"

	case player.MoveLeftKey:
		if player.X == 0 {
			return "wall, cant move"
		}
		clearScreen()
		gameMap.Place(models.NewField(models.FIELD, player.X, player.Y))
		player.MoveLeft()
		gameMap.Place(player)
		fmt.Println(gameMap.Print())
		return "moved left"

	case player.MoveDownKey:
		if player.Y == 0 {
			return "wall, cant move"
		}
		clearScreen()
		gameMap.Place(models.NewField(models.FIELD, player.X, player.Y))
		player.MoveDown()
		gameMap.Place(player)
		fmt.Println(gameMap.Print())
		return "moved down"
	}
	return "nothing happend"
}

func clearScreen() {
	fmt.Print("\033[H\033[2J")
}
