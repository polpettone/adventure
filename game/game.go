package game

import "github.com/polpettone/adventure/engine"

type Game interface {
	Init(engine engine.Engine)
	Update(key string) error
	UpdateEnemies()
	Run()
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
