package models

type Player struct {
	Symbol rune
	X      int
	Y      int

	MoveUpKey    string
	MoveDownKey  string
	MoveRightKey string
	MoveLeftKey  string
	ActionKey    string

	Items []Item
}

func NewPlayer(symbol rune, x, y int, u, d, r, l, a string) *Player {
	return &Player{
		Symbol:       symbol,
		X:            x,
		Y:            y,
		MoveUpKey:    u,
		MoveDownKey:  d,
		MoveRightKey: r,
		MoveLeftKey:  l,
		ActionKey:    a,
	}
}

func (p *Player) addItem(item Item) {
	p.Items = append(p.Items, item)
}

func (p *Player) MoveUp() {
	p.Y += 1
}
func (p *Player) MoveDown() {
	p.Y -= 1
}
func (p *Player) MoveLeft() {
	p.X -= 1
}
func (p *Player) MoveRight() {
	p.X += 1
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
