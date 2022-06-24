package main

import (
	"fmt"
)

type MapElement interface {
	GetSymbol() rune
}

type Map struct {
	Fields [][]MapElement
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

func (m Map) Place(elem MapElement, x, y int) MapElement {
	prevElem := m.Fields[x][y]
	m.Fields[x][y] = elem
	return prevElem
}

type Field struct {
	Symbol rune
}

func NewField(symbol rune) Field {
	return Field{
		Symbol: symbol,
	}
}

func (f Field) GetSymbol() rune {
	return f.Symbol
}

type Bot struct {
	Symbol rune
}

func NewBot(symbol rune) Field {
	return Field{
		Symbol: symbol,
	}
}

func (b Bot) GetSymbol() rune {
	return b.Symbol
}
