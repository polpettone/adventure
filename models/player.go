package models

import (
	"github.com/google/uuid"
)

type Player struct {
	ID     uuid.UUID
	Symbol rune
	X      int
	Y      int

	Displayed bool

	MoveUpKey    string
	MoveDownKey  string
	MoveRightKey string
	MoveLeftKey  string
	ActionKey    string

	Items []Item
}

func NewPlayer(symbol rune, x, y int, u, d, r, l, a string) *Player {
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
		Displayed:    true,
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

func (b Player) GetSymbol() rune {
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

func (i Player) IsDisplayed() bool {
	return i.Displayed
}

func (i Player) DisplayOn() {
	i.Displayed = true
}

func (i Player) DisplayOff() {
	i.Displayed = false
}
