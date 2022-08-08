package models

import (
	"github.com/google/uuid"
	"github.com/polpettone/adventure/logging"
)

type Item struct {
	Symbol    rune
	X         int
	Y         int
	ID        uuid.UUID
	Displayed bool
}

func NewItem(symbol rune, x, y int) *Item {
	return &Item{
		ID:        uuid.New(),
		Symbol:    symbol,
		X:         x,
		Y:         y,
		Displayed: true,
	}
}

func (i Item) GetSymbol() rune {
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

func (i Item) IsDisplayed() bool {
	return i.Displayed
}

func (i Item) DisplayOn() {
	i.Displayed = true
}

func (i Item) DisplayOff() {
	i.Displayed = false
	logging.Log.DebugLog.Printf("Item %s display off", i.GetID().String())
}
