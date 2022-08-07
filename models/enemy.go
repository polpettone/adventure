package models

import (
	"github.com/google/uuid"
)

type Enemy struct {
	ID        uuid.UUID
	Symbol    rune
	X         int
	Y         int
	Displayed bool
}

func NewEnemy(symbol rune, x, y int) *Enemy {
	return &Enemy{
		ID:        uuid.New(),
		Symbol:    symbol,
		X:         x,
		Y:         y,
		Displayed: true,
	}
}

func (e *Enemy) MoveUp() {
	e.Y += 1
}
func (e *Enemy) MoveDown() {
	e.Y -= 1
}
func (e *Enemy) MoveLeft() {
	e.X -= 1
}
func (e *Enemy) MoveRight() {
	e.X += 1
}

func (e Enemy) GetSymbol() rune {
	return e.Symbol
}

func (e Enemy) GetX() int {
	return e.X
}

func (e Enemy) GetY() int {
	return e.Y
}

func (e Enemy) GetID() uuid.UUID {
	return e.ID
}

func (i Enemy) IsDisplayed() bool {
	return i.Displayed
}

func (i Enemy) DisplayOn() {
	i.Displayed = true
}

func (i Enemy) DisplayOff() {
	i.Displayed = false
}
