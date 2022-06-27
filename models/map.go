package models

import "fmt"

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
	for y := m.MaxY - 1; y >= 0; y-- {
		for x := 0; x < m.MaxX; x++ {
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
