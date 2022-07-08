package models

import "fmt"

type MapElement interface {
	GetSymbol() rune
	GetX() int
	GetY() int
}

type Map struct {
	Positions   [][]MapPosition
	MaxX        int
	MaxY        int
	StatusLines []string
}

type MapPosition struct {
	Element MapElement
	X       int
	Y       int
}

func NewMap(maxX, maxY int) Map {

	statusLines := []string{"Dummy status line one", "Dummy status line two"}
	positions := make([][]MapPosition, maxX)
	for n := range positions {
		positions[n] = make([]MapPosition, maxY)
	}

	m := Map{
		Positions:   positions,
		MaxX:        maxX,
		MaxY:        maxY,
		StatusLines: statusLines,
	}

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
			s += fmt.Sprintf(string(m.Positions[x][y].Element.GetSymbol()))
		}
		s += "\n"
	}
	s += "\n"
	for _, l := range m.StatusLines {
		s += fmt.Sprintf("%s \n", l)
	}
	return s
}

func (m Map) Place(elem MapElement) MapElement {
	prevElem := m.Positions[elem.GetX()][elem.GetY()].Element
	m.Positions[elem.GetX()][elem.GetY()].Element = elem
	return prevElem
}
