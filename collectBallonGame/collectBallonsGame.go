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

const GAME_FREQUENCE time.Duration = time.Second / 20

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

func (g *CollectBallonsGame) GetName() string {
	return "Collect Ballons"
}

func (g *CollectBallonsGame) Init(engine engine.Engine) {

	g.Player1 = models.NewPlayer(models.PLAYER, gameConfig.InitPlayerPos.X, gameConfig.InitPlayerPos.Y, gameConfig.PlayerControlMap)
	g.GameMap = models.NewMap(gameConfig.MapSize.X, gameConfig.MapSize.Y)
	g.Items = initializeItems(gameConfig.ItemCount, *g.GameMap, gameConfig.ItemSymbol)

	g.Clock = 0 * time.Minute
	g.Engine = engine
	g.GameState = RUNNING
}

func countCollectableItems(items map[uuid.UUID]*models.Item) int {
	c := 0
	for _, i := range items {
		if i.Collectable {
			c++
		}
	}
	return c
}

func (g *CollectBallonsGame) checkGameOverCriteria() {

	if g.GameState == RUNNING && countCollectableItems(g.Items) == 0 {

		finishTime := g.Clock
		g.GameState = GAMEOVER
		g.GameMap.SetStatusLine(
			"Gameover",
			fmt.Sprintf("All Ballons collected in %v. GameOver", finishTime))
	}
}

func (g CollectBallonsGame) Run() {

	gameControlChannel := make(chan bool, 1)
	impulseChannel := make(chan bool, 1)
	keyChannel := make(chan string, 1)
	go impulseGenerator(impulseChannel, GAME_FREQUENCE, gameControlChannel)
	go inputKeyReceiver(keyChannel)
	go inputKeyHandler(gameControlChannel, keyChannel, impulseChannel, &g)
	go gameOverHandler(&g, gameControlChannel)

	select {
	case <-gameControlChannel:
		logging.Log.InfoLog.Printf("%s Gameover", g.GetName())
		return
	}
}

func (g *CollectBallonsGame) Update() error {
	g.Engine.ClearScreen()
	g.statusLineForPlayer(*g.Player1, "p1")
	g.GameMap.Update(g.GetElements())
	fmt.Println(g.GameMap.Print())
	logElementStates(g.GetElements())
	g.checkGameOverCriteria()
	g.Clock += GAME_FREQUENCE
	g.GameMap.SetStatusLine("Clock", fmt.Sprintf("%v", g.Clock))
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
		if item != nil && item.Collectable {
			delete(g.Items, item.GetID())
			kaputtItem := models.NewItem(models.KAPUTT, player.X, player.Y, false)
			g.Items[kaputtItem.ID] = kaputtItem
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

func gameOverHandler(g *CollectBallonsGame, gameControlChannel chan bool) {

	for {
		time.Sleep(time.Second)
		if g.GameState == GAMEOVER {
			g.Engine.ClearScreen()
			gameControlChannel <- true
		}
	}

}

func inputKeyHandler(
	gameControlChannel chan bool,
	keyChannel chan string,
	impulseChannel chan bool,
	g *CollectBallonsGame) {

	for {
		select {

		case key := <-keyChannel:

			switch key {
			case "q":
				gameControlChannel <- true
			default:
				updatePlayer(key, g.Player1, g)
			}
		case <-impulseChannel:
			g.Update()
		default:
		}

	}

}

func impulseGenerator(
	impulseChannel chan bool,
	frequence time.Duration,
	gameControlChannel chan bool) {
	for {
		impulseChannel <- true
		time.Sleep(frequence)

		select {
		case <-gameControlChannel:

			logging.Log.InfoLog.Println("Stop Impulse Generator")
			return
		default:
		}
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
		models.NewItem(models.BALLON, 10, 10, true),
	}

	rand.Seed(time.Now().UnixNano())

	minX := 1
	maxX := gameMap.MaxX - 2

	minY := 1
	maxY := gameMap.MaxY - 2

	for n := 0; n < count; n++ {
		x := rand.Intn(maxX-minX+1) + minX
		y := rand.Intn(maxY-minY+1) + minY
		items = append(items, models.NewItem(itemSymbol, x, y, true))
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
