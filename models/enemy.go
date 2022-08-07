package models

type Enemy struct {
	Symbol rune
	X      int
	Y      int
}

func NewEnemy(symbol rune, x, y int) *Enemy {
	return &Enemy{
		Symbol: symbol,
		X:      x,
		Y:      y,
	}
}

func (e *Enemy) MoveUp() {
	e.Y += 1
}
func (e *Enemy) MoveDown() {
	e.Y -= 1
}
func (e *Enemy) MoveLeft() {
	e.X -= 1
}
func (e *Enemy) MoveRight() {
	e.X += 1
}

func (e Enemy) GetSymbol() rune {
	return e.Symbol
}

func (e Enemy) GetX() int {
	return e.X
}

func (e Enemy) GetY() int {
	return e.Y
}
