package pinguinBurfGame

import (
	"fmt"
	"math/rand"
	"os"
	"reflect"
	"sync"
	"text/tabwriter"
	"time"

	"github.com/google/uuid"
	"github.com/polpettone/adventure/engine"
	"github.com/polpettone/adventure/game"
	"github.com/polpettone/adventure/logging"
	"github.com/polpettone/adventure/models"
)

const GAME_FREQUENCE time.Duration = time.Second / 20

type GameState int

const (
	RUNNING GameState = iota
	GAMEOVER
)

type PinguinBurfGame struct {
	GameMap *models.Map
	Player1 *models.Player
	Player2 *models.Player

	GameState GameState

	Items   map[uuid.UUID]*models.Item
	Enemies map[uuid.UUID]*models.Enemy

	Engine engine.Engine

	ImpulseChannel chan struct{}
	DoneChannel    chan struct{}
	KeyChannel     chan string

	Frequence time.Duration
}

func (g PinguinBurfGame) checkPlayerLifes() (*models.Player, bool) {

	if g.Player1.LifeCount == 0 {
		return g.Player1, true
	}

	if g.Player2.LifeCount == 0 {
		return g.Player2, true
	}

	return nil, false
}

func (g *PinguinBurfGame) checkGameOverCriteria() {

	player, dead := g.checkPlayerLifes()

	if g.GameState == RUNNING && dead {

		g.GameState = GAMEOVER
		g.GameMap.SetStatusLine(
			"Gameover",
			fmt.Sprintf("GameOver \n %s is dead. \n Press q to go to main menu", player.Symbol))
	}
}

func (g *PinguinBurfGame) GetName() string {
	return "Pinguin Barf"
}

func (g *PinguinBurfGame) Init(engine engine.Engine) {

	g.GameState = RUNNING

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
		Left:   "a",
		Right:  "d",
		Action: "x",
	}
	player1 := models.NewPlayer(models.PLAYER, 0, 0, player1ControlMap)
	player2 := models.NewPlayer(models.PLAYER2, 29, 29, player2ControlMap)

	enemy := models.NewEnemy(models.PENGUIN, 3, 3)
	enemyMap := map[uuid.UUID]*models.Enemy{}
	enemyMap[enemy.ID] = enemy

	itemsMap := initItems(20)

	gameMap := models.NewMap(30, 30)

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

	g.ImpulseChannel = make(chan struct{})
	g.DoneChannel = make(chan struct{})
	g.KeyChannel = make(chan string, 1)
	g.Frequence = GAME_FREQUENCE

}

func (g *PinguinBurfGame) Run() {

	wg := new(sync.WaitGroup)
	wg.Add(3)

	go g.impulseGenerator(wg)
	go g.inputKeyReceiver(wg)
	go g.gameHandler(wg)

	wg.Wait()
	close(g.KeyChannel)
	logging.Log.InfoLog.Printf("%s Gameover", g.GetName())

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

	g.checkGameOverCriteria()

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

func (g *PinguinBurfGame) inputKeyReceiver(wg *sync.WaitGroup) {
	defer wg.Done()
	logging.Log.InfoLog.Println("Start InputKeyReceiver")
	var b []byte = make([]byte, 1)
	for {

		select {
		case _, ok := <-g.DoneChannel:
			if !ok {
				logging.Log.InfoLog.Println("Input KeyReceiver stopped")
				return
			}
		default:
			//take care, this will block until key pressed
			os.Stdin.Read(b)
			i := string(b)
			g.KeyChannel <- i
		}
	}
}

func (g *PinguinBurfGame) gameHandler(wg *sync.WaitGroup) {
	defer wg.Done()

	for {
		select {

		case key, ok := <-g.KeyChannel:
			if !ok {
				logging.Log.InfoLog.Println("KeyChannel Closed, leave gameHandler")
				return
			}
			switch key {
			case "q":
				close(g.DoneChannel)
				close(g.ImpulseChannel)
				return
			default:
				g.Update(key)
			}

		case _, ok := <-g.ImpulseChannel:
			if !ok {
				logging.Log.InfoLog.Println("ImpulseChannel Closed, leave gameHandler")
				return
			}
			g.UpdateEnemies()
		default:

		}
	}

}

func (g *PinguinBurfGame) impulseGenerator(wg *sync.WaitGroup) {
	defer wg.Done()
	for {
		select {
		case _, ok := <-g.DoneChannel:
			if !ok {
				logging.Log.InfoLog.Println("Stop Impulse Generator")
				return
			}
		default:
			g.ImpulseChannel <- struct{}{}
			time.Sleep(g.Frequence)
		}
	}
}
