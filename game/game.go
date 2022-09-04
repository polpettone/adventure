package game

import "github.com/polpettone/adventure/engine"

type Game interface {
	Init(engine engine.Engine)
	Run()
	GetName() string
}

type Coord struct {
	X int
	Y int
}

func NewCoord(x, y int) Coord {
	return Coord{X: x, Y: y}
}

type ControlMap struct {
	Up     string
	Down   string
	Left   string
	Right  string
	Action string
}

type GameConfig struct {
	MapSize          Coord
	ItemCount        int
	InitPlayerPos    Coord
	PlayerControlMap ControlMap
	ItemSymbol       string
}
