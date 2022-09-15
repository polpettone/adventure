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

	go inputKeyReceiver(keyChannel)
	go inputKeyHandler(keyChannel, g)
	select {}
}

func inputKeyReceiver(keyChannel chan string) {
	var b []byte = make([]byte, 1)
	for {
		os.Stdin.Read(b)
		i := string(b)
		logging.Log.InfoLog.Printf("%s inputKeyReceiver", i)
		keyChannel <- i
	}
}

func inputKeyHandler(keyChannel chan string, g *GameSelectionMenu) {

	//pprof.Lookup("goroutine").WriteTo(logging.Log.DebugLog.Writer(), 1)
	for {
		select {

		case key := <-keyChannel:

			logging.Log.InfoLog.Printf("%s pressed", key)

			if g.State == WAIT_FOR_SELECTION {
				switch key {

				case "q":
					fmt.Printf("%s", "bye bye")
					os.Exit(0)
				default:
					if v, ok := g.Entries[key]; ok {
						g.State = GAME_IS_RUNNING
						v.Game.Run()
						g.State = WAIT_FOR_SELECTION
						time.Sleep(3 * time.Second)
						logging.Log.InfoLog.Println("Print Selection ")
						g.Engine.ClearScreen()
						fmt.Println(g.Print())
					} else {
						fmt.Printf("No Entry for %s\n", key)
					}
				}
			}
		default:
		}
	}
}

func (v GameSelectionMenu) Print() string {
	s := fmt.Sprintln("Choose a game by pressing a number")

	keys := make([]string, len(v.Entries))

	for k := range v.Entries {
		keys = append(keys, k)
	}

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
