package models

import (
	"github.com/google/uuid"
)

type Element interface {
	GetSymbol() rune
	GetX() int
	GetY() int
	GetID() uuid.UUID
	IsDisplayed() bool
	DisplayOn()
	DisplayOff()
}
