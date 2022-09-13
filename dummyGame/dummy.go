package dummyGame

import (
	"fmt"

	"github.com/polpettone/adventure/engine"
)

type DummyGame struct {
	Engine engine.Engine
}

func (g *DummyGame) GetName() string {
	return "dummy game"
}

func (g *DummyGame) Init(engine engine.Engine) {
	g.Engine = engine
}

func (g *DummyGame) Run() {

	g.Engine.ClearScreen()
	fmt.Println("Play me")

}
