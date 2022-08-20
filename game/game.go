package game

import "github.com/polpettone/adventure/engine"

type Game interface {
	Init(engine engine.Engine)
	Update(key string) error
	Run() bool
}
