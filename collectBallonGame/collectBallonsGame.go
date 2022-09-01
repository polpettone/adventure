package collectBallonGame

import (
	"fmt"
	"math/rand"
	"os"
	"time"

	"github.com/google/uuid"
	"github.com/polpettone/adventure/engine"
	"github.com/polpettone/adventure/logging"
	"github.com/polpettone/adventure/models"
)

type CollectBallonsGame struct {
	GameMap *models.Map
	Player1 *models.Player
	Items   map[uuid.UUID]*models.Item
	Engine  engine.Engine
}

func (g *CollectBallonsGame) Init(engine engine.Engine) {

	g.Player1 = models.NewPlayer(models.PLAYER, 0, 0, "k", "j", "l", "h", "m")
	g.GameMap = models.NewMap(30, 30)
	g.Items = initializeItems(20, *g.GameMap, models.BALLON)

	elements := buildElements(g.Items, *g.Player1)
	g.GameMap.Update(elements)

	g.Engine = engine
}

func (g CollectBallonsGame) Run() {

	impulseChannel := make(chan bool, 1)
	keyChannel := make(chan string, 1)
	go impulseGenerator(impulseChannel, time.Second/10)
	go inputKeyReceiver(keyChannel)
	go inputKeyHandler(keyChannel, impulseChannel, &g)

	select {}
}

func (g CollectBallonsGame) Update(key string) error {
	updatePlayer(key, g.Player1, g)
	g.Engine.ClearScreen()
	fmt.Println(g.GameMap.Print())
	return nil
}

func updatePlayer(key string, player *models.Player, g CollectBallonsGame) {
	switch key {

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
			err := g.Engine.PlaySound("assets/gunshot.mp3")
			if err != nil {
				logging.Log.InfoLog.Println(err)
			}
		}
	}

}

func inputKeyHandler(keyChannel chan string, impulseChannel chan bool, g *CollectBallonsGame) {

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
		}

	}

}

func impulseGenerator(impulseChannel chan bool, frequence time.Duration) {
	for {
		impulseChannel <- true
		time.Sleep(frequence)
	}
}

func inputKeyReceiver(keyChannel chan string) {
	var b []byte = make([]byte, 1)
	for {
		os.Stdin.Read(b)
		i := string(b)
		keyChannel <- i
	}
}

func initializeItems(
	count int,
	gameMap models.Map,
	itemSymbol string) map[uuid.UUID]*models.Item {

	items := []*models.Item{
		models.NewItem(models.BALLON, 10, 10),
	}

	rand.Seed(time.Now().UnixNano())

	minX := 1
	maxX := gameMap.MaxX - 2

	minY := 1
	maxY := gameMap.MaxY - 2

	for n := 0; n < count; n++ {
		x := rand.Intn(maxX-minX+1) + minX
		y := rand.Intn(maxY-minY+1) + minY
		items = append(items, models.NewItem(itemSymbol, x, y))
	}

	itemsMap := map[uuid.UUID]*models.Item{}
	for _, item := range items {
		itemsMap[item.GetID()] = item
	}

	return itemsMap
}

func buildElements(
	items map[uuid.UUID]*models.Item,
	player1 models.Player) []models.Element {
	elements := []models.Element{}
	for _, i := range items {
		elements = append(elements, i)
	}
	elements = append(elements, player1)
	return elements
}
