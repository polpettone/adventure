package models

import (
	"github.com/google/uuid"
)

type Player struct {
	ID     uuid.UUID
	Symbol string
	X      int
	Y      int

	LifeCount int

	MoveUpKey    string
	MoveDownKey  string
	MoveRightKey string
	MoveLeftKey  string
	ActionKey    string

	Items []Item
}

func NewPlayer(symbol string, x, y int, u, d, r, l, a string) *Player {
	return &Player{
		ID:           uuid.New(),
		Symbol:       symbol,
		X:            x,
		Y:            y,
		MoveUpKey:    u,
		MoveDownKey:  d,
		MoveRightKey: r,
		MoveLeftKey:  l,
		ActionKey:    a,
		LifeCount:    3,
	}
}

func (p *Player) AddItem(item Item) {
	p.Items = append(p.Items, item)
}

func (p *Player) MoveUp() {
	p.Y += 1
}
func (p *Player) MoveDown() {
	p.Y -= 1
}
func (p *Player) MoveLeft() {
	p.X -= 1
}
func (p *Player) MoveRight() {
	p.X += 1
}

func (b Player) GetSymbol() string {
	return b.Symbol
}

func (b Player) GetX() int {
	return b.X
}

func (b Player) GetY() int {
	return b.Y
}

func (p Player) GetID() uuid.UUID {
	return p.ID
}
