package models

type Field struct {
	Symbol rune
	X      int
	Y      int
}

func NewField(symbol rune, x, y int) Field {
	return Field{
		Symbol: symbol,
		X:      x,
		Y:      y,
	}
}

func (f Field) GetSymbol() rune {
	return f.Symbol
}

func (f Field) GetX() int {
	return f.X
}

func (f Field) GetY() int {
	return f.Y
}
