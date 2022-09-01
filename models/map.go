package models

import (
	"fmt"
)

type Map struct {
	MaxX        int
	MaxY        int
	StatusLines map[string]string

	Positions [][]Element
}

func NewMap(maxX, maxY int) *Map {

	statusLines := map[string]string{}

	positions := make([][]Element, maxX)
	for n := range positions {
		positions[n] = make([]Element, maxY)
	}

	m := &Map{
		MaxX:        maxX,
		MaxY:        maxY,
		StatusLines: statusLines,
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

func (m *Map) Update(elements []Element) {
	m.Clear()
	for _, elem := range elements {
		m.place(elem)
	}
}

func (m *Map) GetElementFromPos(x, y int) Element {
	return m.Positions[x][y]
}

func (m *Map) SetStatusLine(key string, text string) {
	m.StatusLines[key] = text
}

func (m *Map) Print() string {
	var s string
	for y := m.MaxY - 1; y >= 0; y-- {
		for x := 0; x < m.MaxX; x++ {
			e := m.Positions[x][y]
			s += fmt.Sprintf(string(e.GetSymbol()))
		}
		s += "\n"
	}
	s += "\n"

	s += printStatusLine(m.StatusLines)

	return s
}

func printStatusLine(statusLines map[string]string) string {

	keys := make([]string, len(statusLines))
	for k := range statusLines {
		keys = append(keys, k)
	}

	s := ""
	for _, k := range keys {
		s += fmt.Sprintf("%s \n", statusLines[k])
	}

	return s

}

func (m *Map) place(elem Element) {
	if elem.GetX() < m.MaxX &&
		elem.GetY() < m.MaxY {
		m.Positions[elem.GetX()][elem.GetY()] = elem
	}
}
