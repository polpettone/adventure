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

	Items   map[uuid.UUID]*models.Item
	Enemies map[uuid.UUID]*models.Enemy
}

func (se SimpleEngine) GetElements() []models.Element {
	return buildElementsForUpdate(se.Items, se.Enemies, *se.Player1, *se.Player2)

}

func InitEngine() Engine {

	player1 := models.NewPlayer(models.PLAYER, 0, 0, "k", "j", "l", "h", "m")
	player2 := models.NewPlayer(models.PLAYER2, 79, 29, "w", "s", "d", "a", "x")

	enemy := models.NewEnemy(models.PENGUIN, 3, 3)
	enemyMap := map[uuid.UUID]*models.Enemy{}
	enemyMap[enemy.ID] = enemy

	items := []*models.Item{
		models.NewItem(models.BOX, 30, 5),
		models.NewItem(models.BOX, 10, 10),
		models.NewItem(models.BOX, 40, 15),
		models.NewItem(models.BOX, 55, 20),
	}

	itemsMap := map[uuid.UUID]*models.Item{}
	for _, item := range items {
		itemsMap[item.GetID()] = item
	}

	elements := buildElementsForUpdate(itemsMap, enemyMap, *player1, *player2)

	gameMap := models.NewMap(80, 30, elements)
	gameMap.Update(elements)

	clearScreen()
	fmt.Println(gameMap.Print())

	var engine Engine = SimpleEngine{
		GameMap: gameMap,
		Player1: player1,
		Player2: player2,
		Items:   itemsMap,
	}
	return engine
}

func (se SimpleEngine) Machine(key string) {
	updatePlayer(key, se.Player1, se)
	updatePlayer(key, se.Player2, se)

	clearScreen()

	se.GameMap.Update(se.GetElements())
	fmt.Println(se.GameMap.Print())

	logElementStates(se.GetElements())
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

	element := se.GameMap.GetElementFromPos(player.X, player.Y)

	switch element.(type) {

	case *models.Item:
		logging.Log.DebugLog.Println("Item found")
		item := se.Items[element.GetID()]
		if item != nil {
			item.DisplayOff()
		}
	}

	logElement(element)
}

func clearScreen() {
	fmt.Print("\033[H\033[2J")
}

func logElementStates(elements []models.Element) {
	for _, elem := range elements {
		logElement(elem)
	}
}

func logElement(elem models.Element) {
	tw := tabwriter.NewWriter(logging.Log.DebugLog.Writer(), 1, 4, 1, '\t', 1)
	fmt.Fprint(tw, fmt.Sprintf(
		"%s \t%s \t %t \n",
		reflect.TypeOf(elem).String(),
		elem.GetID().String(),
		elem.IsDisplayed()),
	)
	tw.Flush()
}

func buildElementsForUpdate(
	items map[uuid.UUID]*models.Item,
	enemies map[uuid.UUID]*models.Enemy,
	player1 models.Player,
	player2 models.Player,
) []models.Element {

	elements := []models.Element{}

	for _, i := range items {
		elements = append(elements, i)
	}

	for _, e := range enemies {
		elements = append(elements, e)
	}

	elements = append(elements, player1)
	elements = append(elements, player2)

	return elements
}
