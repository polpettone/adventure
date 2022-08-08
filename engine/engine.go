package engine

import (
	"fmt"
	"reflect"
	"text/tabwriter"

	"github.com/google/uuid"
	"github.com/polpettone/adventure/logging"
	"github.com/polpettone/adventure/models"
)

type Engine interface {
	Machine(text string)
}

type SimpleEngine struct {
	GameMap *models.Map
	Player1 *models.Player
	Player2 *models.Player

	Elements map[uuid.UUID]models.Element
}

func InitEngine() Engine {

	player1 := models.NewPlayer(models.PLAYER, 0, 0, "k", "j", "l", "h", "m")
	player2 := models.NewPlayer(models.PLAYER2, 79, 29, "w", "s", "d", "a", "x")

	enemy := models.NewEnemy(models.PENGUIN, 3, 3)

	mapElements := []models.Element{
		models.NewItem(models.BOX, 30, 5),
		models.NewItem(models.BOX, 10, 10),
		models.NewItem(models.BOX, 40, 15),
		models.NewItem(models.BOX, 55, 20),
		player1,
		player2,
		enemy,
	}

	elements := map[uuid.UUID]models.Element{}

	for _, elem := range mapElements {
		elements[elem.GetID()] = elem
	}

	gameMap := models.NewMap(80, 30, elements)
	gameMap.Update(elements)

	clearScreen()
	fmt.Println(gameMap.Print())

	var engine Engine = SimpleEngine{
		GameMap:  gameMap,
		Player1:  player1,
		Player2:  player2,
		Elements: elements,
	}
	return engine
}

func (se SimpleEngine) Machine(key string) {
	updatePlayer(key, se.Player1, se)
	updatePlayer(key, se.Player2, se)

	clearScreen()

	se.GameMap.Update(se.Elements)
	fmt.Println(se.GameMap.Print())

	logElementStates(se.Elements)
}

func logElementStates(elements map[uuid.UUID]models.Element) {

	tw := tabwriter.NewWriter(logging.Log.DebugLog.Writer(), 1, 4, 1, '\t', 1)

	for _, elem := range elements {
		fmt.Fprint(tw, fmt.Sprintf(
			"%s \t%s \t %t \n",
			reflect.TypeOf(elem).String(),
			elem.GetID().String(),
			elem.IsDisplayed()),
		)
	}
	tw.Flush()
}

func updatePlayer(key string, player *models.Player, se SimpleEngine) {
	switch key {

	case player.ActionKey:
		if player.X == se.GameMap.MaxX-1 {
			return
		}
		player.X += 1

	case player.MoveUpKey:
		if player.Y == se.GameMap.MaxY-1 {
			return
		}
		player.MoveUp()

	case player.MoveRightKey:
		if player.X == se.GameMap.MaxX-1 {
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

	mapElement := se.GameMap.GetElementFromPos(player.X, player.Y)

	mapElement.DisplayOff()

	t := fmt.Sprintf("%T", mapElement)

	se.GameMap.SetStatusLine(0, string(mapElement.GetSymbol())+" "+t)

}

func clearScreen() {
	fmt.Print("\033[H\033[2J")
}
