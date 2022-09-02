package collectBallonGame

import (
	"fmt"
	"math/rand"
	"os"
	"reflect"
	"text/tabwriter"
	"time"

	"github.com/google/uuid"
	"github.com/polpettone/adventure/engine"
	"github.com/polpettone/adventure/logging"
	"github.com/polpettone/adventure/models"
)

const GAME_FREQUENCE time.Duration = time.Second / 10

type GameState int

const (
	RUNNING GameState = iota
	GAMEOVER
)

type CollectBallonsGame struct {
	GameMap   *models.Map
	Player1   *models.Player
	Items     map[uuid.UUID]*models.Item
	Engine    engine.Engine
	Clock     time.Duration
	GameState GameState
}

func (g *CollectBallonsGame) Init(engine engine.Engine) {

	g.Player1 = models.NewPlayer(models.PLAYER, 0, 0, "k", "j", "l", "h", "m")
	g.GameMap = models.NewMap(30, 30)
	g.Items = initializeItems(1, *g.GameMap, models.BALLON)

	elements := buildElements(g.Items, *g.Player1)
	g.GameMap.Update(elements)

	g.Clock = 0 * time.Minute
	g.Engine = engine
	g.GameState = RUNNING
}

func (g *CollectBallonsGame) checkGameOverCriteria() {

	if g.GameState == RUNNING && len(g.Items) == 0 {

		finishTime := g.Clock
		g.GameState = GAMEOVER
		g.GameMap.SetStatusLine(
			"Gameover",
			fmt.Sprintf("All Ballons collected in %v. GameOver", finishTime))
	}
}

func (g CollectBallonsGame) Run() {

	impulseChannel := make(chan bool, 1)
	keyChannel := make(chan string, 1)
	go impulseGenerator(impulseChannel, GAME_FREQUENCE)
	go inputKeyReceiver(keyChannel)
	go inputKeyHandler(keyChannel, impulseChannel, &g)

	select {}
}

func (g *CollectBallonsGame) Update(key string) error {
	updatePlayer(key, g.Player1, g)
	g.Engine.ClearScreen()
	g.statusLineForPlayer(*g.Player1, "p1")
	g.GameMap.Update(g.GetElements())
	fmt.Println(g.GameMap.Print())
	logElementStates(g.GetElements())
	g.checkGameOverCriteria()
	return nil
}

func (g CollectBallonsGame) GetElements() []models.Element {
	return buildElements(g.Items, *g.Player1)
}

func (g CollectBallonsGame) statusLineForPlayer(player models.Player, key string) {
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
func updatePlayer(key string, player *models.Player, g *CollectBallonsGame) {
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
		case <-impulseChannel:
			g.Clock += GAME_FREQUENCE
			g.GameMap.SetStatusLine("Clock", fmt.Sprintf("%v", g.Clock))
			g.Engine.ClearScreen()
			g.GameMap.Update(g.GetElements())
			fmt.Println(g.GameMap.Print())
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
