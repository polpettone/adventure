package menu

import (
	"fmt"
	"os"
	"strconv"

	"github.com/polpettone/adventure/engine"
	"github.com/polpettone/adventure/game"
)

type MenuEntry struct {
	SelectioValue string
	Selected      bool
	Game          game.Game
}

type GameSelectionMenu struct {
	Engine  engine.Engine
	Entries map[string]*MenuEntry
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
	}

	return m
}

func (v *GameSelectionMenu) Run() {

	fmt.Println(v.Print())

	keyChannel := make(chan string, 1)
	go inputKeyReceiver(keyChannel)
	go inputKeyHandler(keyChannel, v.Entries)
	select {}
}

func (v GameSelectionMenu) Print() string {
	s := fmt.Sprintln("Choose a game by pressing a number")
	for _, e := range v.Entries {
		s += fmt.Sprintf("%s)  %s \n", e.SelectioValue, e.Game.GetName())
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

func inputKeyHandler(keyChannel chan string, entries map[string]*MenuEntry) {

	for {
		select {

		case key := <-keyChannel:

			switch key {
			case "q":
				fmt.Printf("%s", "bye bye")
				os.Exit(0)
			default:
				if v, ok := entries[key]; ok {
					v.Game.Run()
				} else {
					fmt.Printf("No Entry for %s", key)
				}

			}
		}
	}
}
