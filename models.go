package main

import (
	"fmt"
)

type MapElement interface {
	GetSymbol() rune
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
