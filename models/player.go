package models

type Player struct {
	Symbol rune
	X      int
	Y      int
}

func NewPlayer(symbol rune, x, y int) *Player {
	return &Player{
		Symbol: symbol,
		X:      x,
		Y:      y,
	}
}

func (b Player) GetSymbol() rune {
	return b.Symbol
}

func (b Player) GetX() int {
	return b.X
}

func (b Player) GetY() int {
	return b.Y
}
