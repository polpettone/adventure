package game

import (
	"fmt"
	"math/rand"
	"os"
	"reflect"
	"strings"
	"text/tabwriter"
	"time"

	"github.com/google/uuid"
	"github.com/polpettone/adventure/engine"
	"github.com/polpettone/adventure/logging"
	"github.com/polpettone/adventure/models"
)

type PinguinBurfGame struct {
	GameMap *models.Map
	Player1 *models.Player
	Player2 *models.Player

	Items   map[uuid.UUID]*models.Item
	Enemies map[uuid.UUID]*models.Enemy

	Engine engine.Engine
}

func initItems(count int) map[uuid.UUID]*models.Item {
	items := []*models.Item{
		models.NewItem(models.BALLON, 10, 10),
	}

	rand.Seed(time.Now().UnixNano())

	minX := 1
	maxX := 79

	minY := 1
	maxY := 28

	for n := 0; n < count; n++ {
		x := rand.Intn(maxX-minX+1) + minX
		y := rand.Intn(maxY-minY+1) + minY
		items = append(items, models.NewItem(models.BALLON, x, y))
	}

	itemsMap := map[uuid.UUID]*models.Item{}
	for _, item := range items {
		itemsMap[item.GetID()] = item
	}

	return itemsMap
}

func (g *PinguinBurfGame) Init(engine engine.Engine) {

	player1 := models.NewPlayer(models.PLAYER, 0, 0, "k", "j", "l", "h", "m")
	player2 := models.NewPlayer(models.PLAYER2, 79, 29, "w", "s", "d", "a", "x")

	enemy := models.NewEnemy(models.PENGUIN, 3, 3)
	enemyMap := map[uuid.UUID]*models.Enemy{}
	enemyMap[enemy.ID] = enemy

	itemsMap := initItems(500)

	elements := buildElementsForUpdate(itemsMap, enemyMap, *player1, *player2)

	gameMap := models.NewMap(80, 30, elements)
	gameMap.Update(elements)

	gameMap.SetStatusLine(
		2,
		fmt.Sprintf("%s %s",
			string(player1.GetSymbol()),
			"k j l h m",
		),
	)

	gameMap.SetStatusLine(
		3,
		fmt.Sprintf("%s %s",
			string(player2.GetSymbol()),
			"w s d a x",
		),
	)

	g.GameMap = gameMap
	g.Enemies = enemyMap
	g.Items = itemsMap
	g.Enemies = enemyMap
	g.Player1 = player1
	g.Player2 = player2

	g.Engine = engine
	g.Engine.ClearScreen()

	fmt.Println(gameMap.Print())
	logElementStates(elements)
}

func (g PinguinBurfGame) Run() bool {
	var b []byte = make([]byte, 1)
	for {
		os.Stdin.Read(b)
		text := string(b)
		g.Update(text)

		if strings.Compare(text, "q") == 0 {
			return true
		}

		if strings.Compare(text, "r") == 0 {
			return false
		}

	}
}

func (g PinguinBurfGame) Update(key string) error {
	updatePlayer(key, g.Player1, g)
	updatePlayer(key, g.Player2, g)

	g.Engine.ClearScreen()

	g.GameMap.Update(g.GetElements())

	g.statusLineForPlayer(*g.Player1, 0)
	g.statusLineForPlayer(*g.Player2, 1)

	fmt.Println(g.GameMap.Print())

	logElementStates(g.GetElements())

	return nil

}

func (g PinguinBurfGame) statusLineForPlayer(player models.Player, statusLineIndex int) {
	g.GameMap.SetStatusLine(
		statusLineIndex,
		fmt.Sprintf("%s %d %s %d %s",
			string(player.GetSymbol()),
			player.LifeCount,
			models.HEART,
			len(player.Items),
			string(models.BALLON),
		),
	)
}

func (g PinguinBurfGame) GetElements() []models.Element {
	return buildElementsForUpdate(g.Items, g.Enemies, *g.Player1, *g.Player2)
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

func logElementStates(elements []models.Element) {
	for _, elem := range elements {
		logElements(elem)
	}
	logging.Log.DebugLog.Println("-------------------------------------")
}

func logElements(elem models.Element) {
	tw := tabwriter.NewWriter(logging.Log.DebugLog.Writer(), 1, 4, 1, '\t', 1)
	fmt.Fprint(tw, fmt.Sprintf(
		"%s \t%s \t %d %d \n",
		reflect.TypeOf(elem).String(),
		elem.GetID().String(),
		elem.GetX(),
		elem.GetY()),
	)
	tw.Flush()
}

func updatePlayer(key string, player *models.Player, g PinguinBurfGame) {
	switch key {

	case player.ActionKey:
		if player.X == g.GameMap.MaxX-1 {
			return
		}
		enemy := models.NewEnemy(models.PENGUIN, player.X, player.Y)
		g.Enemies[enemy.ID] = enemy
		player.X += 1
		g.Engine.PlaySound()
		return

	case player.MoveUpKey:
		if player.Y == g.GameMap.MaxY-1 {
			return
		}
		player.MoveUp()

	case player.MoveRightKey:
		if player.X == g.GameMap.MaxX-1 {
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

	element := g.GameMap.GetElementFromPos(player.X, player.Y)

	switch element.(type) {

	case *models.Item:
		item := g.Items[element.GetID()]
		if item != nil {
			delete(g.Items, item.GetID())
			logging.Log.DebugLog.Printf("Item %s deleted", item.GetID())
			player.AddItem(*item)
			logging.Log.DebugLog.Println(player.Items)
			g.Engine.PlaySound()
		}

	case *models.Enemy:
		enemy := g.Enemies[element.GetID()]
		if enemy != nil {
			player.LifeCount = player.LifeCount - 1

			delete(g.Enemies, enemy.GetID())
			logging.Log.DebugLog.Printf("Enemy %s deleted", enemy.GetID())
		}

	}

	logElements(element)

}
