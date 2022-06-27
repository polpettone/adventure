package models

type Bot struct {
	Symbol rune
	X      int
	Y      int
}

func NewBot(symbol rune, x, y int) *Bot {
	return &Bot{
		Symbol: symbol,
		X:      x,
		Y:      y,
	}
}

func (b Bot) GetSymbol() rune {
	return b.Symbol
}

func (b Bot) GetX() int {
	return b.X
}

func (b Bot) GetY() int {
	return b.Y
}
