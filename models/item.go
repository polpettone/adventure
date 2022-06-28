package models

type Item struct {
	Symbol rune
	X      int
	Y      int
}

func NewItem(symbol rune, x, y int) Field {
	return Field{
		Symbol: symbol,
		X:      x,
		Y:      y,
	}
}

func (i Item) GetSymbol() rune {
	return i.Symbol
}

func (i Item) GetX() int {
	return i.X
}

func (i Item) GetY() int {
	return i.Y
}
