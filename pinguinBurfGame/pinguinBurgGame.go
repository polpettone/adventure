package pinguinBurfGame

import (
	"fmt"
	"math/rand"
	"os"
	"reflect"
	"text/tabwriter"
	"time"

	"github.com/google/uuid"
	"github.com/polpettone/adventure/engine"
	"github.com/polpettone/adventure/game"
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

func (g *PinguinBurfGame) Init(engine engine.Engine) {

	var player1ControlMap game.ControlMap = game.ControlMap{
		Up:     "k",
		Down:   "j",
		Left:   "h",
		Right:  "l",
		Action: "m",
	}

	var player2ControlMap game.ControlMap = game.ControlMap{
		Up:     "w",
		Down:   "s",
		Left:   "d",
		Right:  "a",
		Action: "x",
	}
	player1 := models.NewPlayer(models.PLAYER, 0, 0, player1ControlMap)
	player2 := models.NewPlayer(models.PLAYER2, 29, 29, player2ControlMap)

	enemy := models.NewEnemy(models.PENGUIN, 3, 3)
	enemyMap := map[uuid.UUID]*models.Enemy{}
	enemyMap[enemy.ID] = enemy

	itemsMap := initItems(20)

	elements := buildElementsForUpdate(itemsMap, enemyMap, *player1, *player2)

	gameMap := models.NewMap(30, 30)

	gameMap.Update(elements)

	gameMap.SetStatusLine(
		"player1",
		fmt.Sprintf("%s %s",
			string(player1.GetSymbol()),
			"k j l h m",
		),
	)

	gameMap.SetStatusLine(
		"player2",
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

func (g PinguinBurfGame) Run() {

	impulseChannel := make(chan bool, 1)
	keyChannel := make(chan string, 1)
	go impulseGenerator(impulseChannel, time.Second/10)
	go inputKeyReceiver(keyChannel)
	go inputKeyHandler(keyChannel, impulseChannel, &g)

	select {}
}

func (g PinguinBurfGame) Update(key string) error {
	updatePlayer(key, g.Player1, g)
	updatePlayer(key, g.Player2, g)

	g.Engine.ClearScreen()

	g.GameMap.Update(g.GetElements())

	g.statusLineForPlayer(*g.Player1, "p1")
	g.statusLineForPlayer(*g.Player2, "p2")

	fmt.Println(g.GameMap.Print())

	logElementStates(g.GetElements())

	return nil
}

func (g PinguinBurfGame) UpdateEnemies() {
	for _, e := range g.Enemies {
		e.Update()
	}
	g.Engine.ClearScreen()

	g.GameMap.Update(g.GetElements())

	g.statusLineForPlayer(*g.Player1, "p1")
	g.statusLineForPlayer(*g.Player2, "p2")

	fmt.Println(g.GameMap.Print())

	logElementStates(g.GetElements())
}

func (g PinguinBurfGame) statusLineForPlayer(player models.Player, key string) {
	g.GameMap.SetStatusLine(
		key,
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

	case player.ControlMap.Action:
		if player.X == g.GameMap.MaxX-1 {
			return
		}
		enemy := models.NewEnemy(models.PENGUIN, player.X, player.Y)
		g.Enemies[enemy.ID] = enemy
		player.X += 1
		err := g.Engine.PlaySound("assets/mixkit-martial-arts-punch-2052.wav")
		if err != nil {
			logging.Log.InfoLog.Println(err)
		}
		return

	case player.ControlMap.Up:
		if player.Y == g.GameMap.MaxY-1 {
			return
		}
		player.MoveUp()

	case player.ControlMap.Right:
		if player.X == g.GameMap.MaxX-1 {
			return
		}
		player.MoveRight()

	case player.ControlMap.Left:
		if player.X == 0 {
			return
		}
		player.MoveLeft()

	case player.ControlMap.Down:
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
			err := g.Engine.PlaySound("assets/gunshot.mp3")
			if err != nil {
				logging.Log.InfoLog.Println(err)
			}
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

func initItems(count int) map[uuid.UUID]*models.Item {
	items := []*models.Item{
		models.NewItem(models.BALLON, 10, 10, true),
	}

	rand.Seed(time.Now().UnixNano())

	minX := 1
	maxX := 28

	minY := 1
	maxY := 28

	for n := 0; n < count; n++ {
		x := rand.Intn(maxX-minX+1) + minX
		y := rand.Intn(maxY-minY+1) + minY
		items = append(items, models.NewItem(models.BALLON, x, y, true))
	}

	itemsMap := map[uuid.UUID]*models.Item{}
	for _, item := range items {
		itemsMap[item.GetID()] = item
	}

	return itemsMap
}

func inputKeyReceiver(keyChannel chan string) {
	var b []byte = make([]byte, 1)
	for {
		os.Stdin.Read(b)
		i := string(b)
		keyChannel <- i
	}
}

func inputKeyHandler(keyChannel chan string, impulseChannel chan bool, g game.Game) {

	for {
		select {

		case key := <-keyChannel:

			switch key {
			case "q":
				fmt.Printf("%s", "bye bye")
				os.Exit(0)
			case "r":
				fmt.Printf("%s", "reload \n")
			default:
				g.Update(key)
			}

		case <-impulseChannel:
			g.UpdateEnemies()
		}
	}

}

func impulseGenerator(impulseChannel chan bool, frequence time.Duration) {
	for {
		impulseChannel <- true
		time.Sleep(frequence)
	}
}
