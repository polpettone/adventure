package menu

import (
	"fmt"
	"os"
	"sort"
	"strconv"
	"time"

	"github.com/polpettone/adventure/engine"
	"github.com/polpettone/adventure/game"
	"github.com/polpettone/adventure/logging"
)

const FREQUENCE time.Duration = time.Second / 20

type MenuEntry struct {
	SelectioValue string
	Selected      bool
	Game          game.Game
}

type GameSelectionMenu struct {
	Engine  engine.Engine
	Entries map[string]*MenuEntry
	State   MenuState
}

type MenuState int

const (
	GAME_IS_RUNNING = iota
	WAIT_FOR_SELECTION
)

func NewGameSelectionMenu(engine engine.Engine, games []game.Game) *GameSelectionMenu {

	entries := map[string]*MenuEntry{}

	i := 1

	for _, g := range games {
		entryIndex := strconv.Itoa(i)
		entries[entryIndex] = &MenuEntry{
			Selected:      false,
			Game:          g,
			SelectioValue: entryIndex,
		}
		i++
	}

	m := &GameSelectionMenu{
		Engine:  engine,
		Entries: entries,
		State:   WAIT_FOR_SELECTION,
	}

	return m
}

func (g *GameSelectionMenu) Run() {

	fmt.Println(g.Print())

	keyChannel := make(chan string, 1)
	impulseChannel := make(chan bool, 1)

	go inputKeyReceiver(keyChannel)
	go inputKeyHandler(keyChannel, impulseChannel, g)
	go impulseGenerator(impulseChannel, FREQUENCE)
	select {}
}

func (v GameSelectionMenu) Print() string {
	s := fmt.Sprintln("Choose a game by pressing a number")

	keys := make([]string, len(v.Entries))

	for k := range v.Entries {
		keys = append(keys, k)
	}

	logging.Log.DebugLog.Println("-----------")
	logging.Log.DebugLog.Println(keys)
	logging.Log.DebugLog.Println(keys)
	logging.Log.DebugLog.Println("-----------")

	logging.Log.DebugLog.Println("-----------")
	logging.Log.DebugLog.Println(v.Entries)
	logging.Log.DebugLog.Println("-----------")

	sort.Strings(keys)
	for _, k := range keys {
		if _, ok := v.Entries[k]; ok {
			s += fmt.Sprintf("%s)  %s \n",
				v.Entries[k].SelectioValue,
				v.Entries[k].Game.GetName())
		}
	}
	return s
}

func inputKeyReceiver(keyChannel chan string) {
	var b []byte = make([]byte, 1)
	for {
		os.Stdin.Read(b)
		i := string(b)
		keyChannel <- i
	}
}

func inputKeyHandler(keyChannel chan string, impulseChannel chan bool, g *GameSelectionMenu) {

	for {
		if g.State == WAIT_FOR_SELECTION {
			select {

			case key := <-keyChannel:

				logging.Log.InfoLog.Printf("%s pressed", key)
				switch key {

				case "q":
					fmt.Printf("%s", "bye bye")
					os.Exit(0)
				default:
					if v, ok := g.Entries[key]; ok {
						g.State = GAME_IS_RUNNING
						logging.Log.InfoLog.Printf("Menu State %d ", g.State)
						v.Game.Run()
						// TODO:
						// Restarts games currently working due
						// a problem with the game selection
						// after fun game finished
						// temporaly quit whole program
						// after finishing a game
						os.Exit(0)
						g.State = WAIT_FOR_SELECTION
						logging.Log.InfoLog.Printf("Menu State %d ", g.State)
					} else {
						fmt.Printf("No Entry for %s", key)
					}
				}
			case <-impulseChannel:
				g.Engine.ClearScreen()
				fmt.Println(g.Print())

			default:
			}
		}
	}
}
func impulseGenerator(
	impulseChannel chan bool,
	frequence time.Duration) {
	for {
		impulseChannel <- true
		time.Sleep(frequence)
	}
}
