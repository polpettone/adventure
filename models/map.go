package models

import (
	"fmt"
)

type Map struct {
	MaxX        int
	MaxY        int
	StatusLines []string

	Elements  []Element
	Positions [][]Element
}

func NewMap(maxX, maxY int, elements []Element) *Map {

	statusLines := []string{
		"Dummy status line 1",
		"Dummy status line 2",
		"Dummy status line 3",
		"Dummy status line 4",
	}

	positions := make([][]Element, maxX)
	for n := range positions {
		positions[n] = make([]Element, maxY)
	}

	m := &Map{
		MaxX:        maxX,
		MaxY:        maxY,
		StatusLines: statusLines,
		Elements:    elements,
		Positions:   positions,
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

func (m *Map) place(elem Element) Element {
	prevElem := m.Positions[elem.GetX()][elem.GetY()]
	m.Positions[elem.GetX()][elem.GetY()] = elem
	return prevElem
}

func (m *Map) Update(elements []Element) {
	m.Clear()
	for _, elem := range elements {
		m.place(elem)
	}
}

func (m *Map) GetElementFromPos(x, y int) Element {
	return m.Positions[x][y]
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

			e := m.Positions[x][y]

			if e.IsDisplayed() {
				s += fmt.Sprintf(string(e.GetSymbol()))
			} else {
				s += string(FIELD)
			}
		}
		s += "\n"
	}
	s += "\n"
	for _, l := range m.StatusLines {
		s += fmt.Sprintf("%s \n", l)
	}
	return s
}
