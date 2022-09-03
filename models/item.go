package models

import (
	"github.com/google/uuid"
)

type Item struct {
	Symbol      string
	X           int
	Y           int
	ID          uuid.UUID
	Collectable bool
}

func NewItem(symbol string, x, y int, collectable bool) *Item {
	return &Item{
		ID:          uuid.New(),
		Symbol:      symbol,
		X:           x,
		Y:           y,
		Collectable: collectable,
	}
}

func (i Item) GetSymbol() string {
	return i.Symbol
}

func (i Item) GetX() int {
	return i.X
}

func (i Item) GetY() int {
	return i.Y
}

func (i Item) GetID() uuid.UUID {
	return i.ID
}
