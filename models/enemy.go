package models

import (
	"math/rand"
	"time"

	"github.com/google/uuid"
)

type Enemy struct {
	ID     uuid.UUID
	Symbol string
	X      int
	Y      int
}

func NewEnemy(symbol string, x, y int) *Enemy {
	return &Enemy{
		ID:     uuid.New(),
		Symbol: symbol,
		X:      x,
		Y:      y,
	}
}

func (e *Enemy) Update() {

	rand.Seed(time.Now().UnixNano())
	choice := rand.Intn(4)

	switch choice {

	case 0:
		e.MoveUp()
	case 1:
		e.MoveDown()
	case 2:
		e.MoveLeft()
	case 3:
		e.MoveRight()

	}

}

func (e *Enemy) MoveDown() {
	if e.Y > 0 {
		e.Y -= 1
	}
}

func (e *Enemy) MoveLeft() {
	if e.X > 0 {
		e.X -= 1
	}
}

func (e *Enemy) MoveUp() {
	e.Y += 1
}

func (e *Enemy) MoveRight() {
	e.X += 1
}

func (e Enemy) GetSymbol() string {
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
