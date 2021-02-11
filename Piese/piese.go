package piese

type Piesa struct {
	x, y         int
	tip, culoare rune
}

func NewPiesa(x, y int, tip, culoare rune) Piesa {
	e := Piesa{x, y, tip, culoare}
	return e
}
