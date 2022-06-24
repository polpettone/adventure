package main

import (
	"fmt"
)

type MapElement interface {
	GetSymbol() rune
	GetX() int
	GetY() int
}

type Map struct {
	Fields [][]MapElement
	MaxX   int
	MaxY   int
}

func NewMap(maxX, maxY int) Map {

	fields := make([][]MapElement, maxX)
	for n := range fields {
		fields[n] = make([]MapElement, maxY)
	}
	m := Map{Fields: fields, MaxX: maxX, MaxY: maxY}
	for x := 0; x < maxX; x++ {
		for y := 0; y < maxY; y++ {
			field := NewField(FIELD, x, y)
			m.Place(field)
		}
	}
	return m

}

func (m Map) Print() string {
	var s string
	for y := 9; y >= 0; y-- {
		for x := 0; x < 10; x++ {
			s += fmt.Sprintf(string(m.Fields[x][y].GetSymbol()))
		}
		s += "\n"
	}
	return s
}

func (m Map) Place(elem MapElement) MapElement {
	prevElem := m.Fields[elem.GetX()][elem.GetY()]
	m.Fields[elem.GetX()][elem.GetY()] = elem
	return prevElem
}

type Field struct {
	Symbol rune
	X      int
	Y      int
}

func NewField(symbol rune, x, y int) Field {
	return Field{
		Symbol: symbol,
		X:      x,
		Y:      y,
	}
}

func (f Field) GetSymbol() rune {
	return f.Symbol
}

func (f Field) GetX() int {
	return f.X
}

func (f Field) GetY() int {
	return f.Y
}

type Bot struct {
	Symbol rune
	X      int
	Y      int
}

func NewBot(symbol rune, x, y int) *Bot {
	return &Bot{
		Symbol: symbol,
		X:      x,
		Y:      y,
	}
}

func (b Bot) GetSymbol() rune {
	return b.Symbol
}

func (b Bot) GetX() int {
	return b.X
}

func (b Bot) GetY() int {
	return b.Y
}
