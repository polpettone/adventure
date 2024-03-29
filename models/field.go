package models

import (
	"github.com/google/uuid"
)

type Field struct {
	ID     uuid.UUID
	Symbol string
	X      int
	Y      int
}

func NewField(symbol string, x, y int) Field {
	return Field{
		ID:     uuid.New(),
		Symbol: symbol,
		X:      x,
		Y:      y,
	}
}

func (f Field) GetSymbol() string {
	return f.Symbol
}

func (f Field) GetX() int {
	return f.X
}

func (f Field) GetY() int {
	return f.Y
}

func (f Field) GetID() uuid.UUID {
	return f.ID
}
