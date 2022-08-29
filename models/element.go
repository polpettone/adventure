package models

import (
	"github.com/google/uuid"
)

type Element interface {
	GetSymbol() string
	GetX() int
	GetY() int
	GetID() uuid.UUID
	IsDisplayed() bool
	DisplayOn()
	DisplayOff()
}
