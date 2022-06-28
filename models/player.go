package models

type Player struct {
	Symbol rune
	X      int
	Y      int

	MoveUp    string
	MoveDown  string
	MoveRight string
	MoveLeft  string
}

func NewPlayer(symbol rune, x, y int, u, d, r, l string) *Player {
	return &Player{
		Symbol:    symbol,
		X:         x,
		Y:         y,
		MoveUp:    u,
		MoveDown:  d,
		MoveRight: r,
		MoveLeft:  l,
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
