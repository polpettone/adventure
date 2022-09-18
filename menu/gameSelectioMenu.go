package menu

import (
	"fmt"
	"os"
	"sort"
	"strconv"

	"github.com/polpettone/adventure/engine"
	"github.com/polpettone/adventure/game"
	"github.com/polpettone/adventure/logging"
)

type MenuEntry struct {
	SelectioValue string
	Selected      bool
	Game          game.Game
}

type MenuState int

const (
	GAME_IS_RUNNING = iota
	WAIT_FOR_SELECTION
)

type GameSelectionMenu struct {
	Engine  engine.Engine
	Entries map[string]*MenuEntry
	State   MenuState

	StopChannel chan struct{}
	KeyChannel  chan string
}

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

		StopChannel: make(chan struct{}),
		KeyChannel:  make(chan string, 1),
	}

	return m
}

func (g *GameSelectionMenu) Run() {

	fmt.Println(g.Print())

	go g.inputKeyReceiver()
	go g.inputKeyHandler()

	select {

	case _, ok := <-g.StopChannel:
		if !ok {
			logging.Log.InfoLog.Println("Finished Game Menu")
			return
		}

	}
}

func (g *GameSelectionMenu) inputKeyReceiver() {
	var b []byte = make([]byte, 1)
	for {

		select {
		case _, ok := <-g.StopChannel:
			if !ok {
				logging.Log.InfoLog.Println("Input KeyReceiver stopped")
				return
			}

		default:

			os.Stdin.Read(b)
			i := string(b)
			g.KeyChannel <- i
		}
	}
}

func (g *GameSelectionMenu) inputKeyHandler() {

	//pprof.Lookup("goroutine").WriteTo(logging.Log.DebugLog.Writer(), 1)
	for {
		select {

		case key := <-g.KeyChannel:

			logging.Log.InfoLog.Printf("%s pressed", key)

			if g.State == WAIT_FOR_SELECTION {
				switch key {

				case "q":
					fmt.Printf("%s", "bye bye")
					close(g.StopChannel)
					close(g.KeyChannel)
					os.Exit(0)
				default:
					if v, ok := g.Entries[key]; ok {
						g.State = GAME_IS_RUNNING

						logging.Log.InfoLog.Println("Run Game")
						v.Game.Init(g.Engine)
						v.Game.Run()
						logging.Log.InfoLog.Println("Game finished")

						g.State = WAIT_FOR_SELECTION

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

func (g GameSelectionMenu) Print() string {
	s := fmt.Sprintln("Choose a game by pressing a number")

	keys := make([]string, len(g.Entries))

	for k := range g.Entries {
		keys = append(keys, k)
	}

	sort.Strings(keys)
	for _, k := range keys {
		if _, ok := g.Entries[k]; ok {
			s += fmt.Sprintf("%s)  %s \n",
				g.Entries[k].SelectioValue,
				g.Entries[k].Game.GetName())
		}
	}
	return s
}
