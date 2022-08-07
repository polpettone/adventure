package models

import (
	"fmt"

	"github.com/google/uuid"
)

type MapElement interface {
	GetSymbol() rune
	GetX() int
	GetY() int
	GetID() uuid.UUID
}

type Map struct {
	MapElements []MapElement
	Positions   [][]MapPosition
	MaxX        int
	MaxY        int
	StatusLines []string
	Enemies     []Enemy
}

type MapPosition struct {
	Element MapElement
}

func NewMap(maxX, maxY int, mapElements []MapElement) *Map {

	statusLines := []string{
		"Dummy status line one",
		"Dummy status line two",
		"Dummy status line three",
	}
	positions := make([][]MapPosition, maxX)
	for n := range positions {
		positions[n] = make([]MapPosition, maxY)
	}

	m := &Map{
		Positions:   positions,
		MaxX:        maxX,
		MaxY:        maxY,
		StatusLines: statusLines,
		MapElements: mapElements,
	}

	m.Clear()

	return m
}

func (m *Map) Clear() {
	for x := 0; x < m.MaxX; x++ {
		for y := 0; y < m.MaxY; y++ {
			field := NewField(FIELD, x, y)
			m.place(field)
		}
	}
}

func (m *Map) Update() {
	m.Clear()
	for _, elem := range m.MapElements {
		m.place(elem)
	}
}

func (m *Map) GetElementFromPos(x, y int) MapElement {
	return m.Positions[x][y].Element
}

func (m *Map) SetStatusLine(number int, text string) {
	if len(m.StatusLines) <= number {
		//TODO: logging
		return
	}
	m.StatusLines[number] = text
}

func (m *Map) Print() string {
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

func (m *Map) place(elem MapElement) MapElement {
	prevElem := m.Positions[elem.GetX()][elem.GetY()].Element
	m.Positions[elem.GetX()][elem.GetY()].Element = elem
	return prevElem
}
