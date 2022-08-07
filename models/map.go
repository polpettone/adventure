package models

import (
	"fmt"
)

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
	Player1     Player
	Player2     Player
	Enemies     []Enemy
}

type MapPosition struct {
	Element MapElement
	X       int
	Y       int
}

func NewMap(maxX, maxY int, player1, player2 Player) *Map {

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
		Player1:     player1,
		Player2:     player2,
	}

	for x := 0; x < maxX; x++ {
		for y := 0; y < maxY; y++ {
			field := NewField(FIELD, x, y)
			m.Place(field)
		}
	}

	return m
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

func (m *Map) Place(elem MapElement) MapElement {
	prevElem := m.Positions[elem.GetX()][elem.GetY()].Element
	m.Positions[elem.GetX()][elem.GetY()].Element = elem
	return prevElem
}

func (m *Map) UpdatePlayer1(player Player) {
	m.Place(NewField(FIELD, m.Player1.X, m.Player1.Y))
	m.Place(player)
	m.Player1 = player
}

func (m *Map) UpdatePlayer2(player Player) {
	m.Place(NewField(FIELD, m.Player2.X, m.Player2.Y))
	m.Place(player)
	m.Player2 = player
}
