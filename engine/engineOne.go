package engine

import (
	"fmt"
)

type SimpleEngine struct {
	Game Game
}

func (e SimpleEngine) Machine(key string) {
	e.Game.Update(key)
}

func clearScreen() {
	fmt.Print("\033[H\033[2J")
}
