package piese

import (
	"github.com/hajimehoshi/ebiten"
	"github.com/hajimehoshi/ebiten/ebitenutil"
)

type Piesa struct {
	Mutabil      bool
	tip, culoare rune
}

func NewPiesa(tip, culoare rune) Piesa {
	e := Piesa{false, tip, culoare}
	return e
}

func (p *Piesa) Draw() *ebiten.Image {
	path := "imagini/"
	if p.culoare == 'W' {
		path += "w_"
	} else {
		path += "b_"
	}
	switch p.tip {
	case 'K':
		path += "king"
	case 'P':
		path += "pawn"
	case 'B':
		path += "bishop"
	case 'N':
		path += "knight"
	case 'R':
		path += "rook"
	case 'Q':
		path += "queen"
	default:
		return nil
	}
	path += "_png_128px.png"
	img, _, _ := ebitenutil.NewImageFromFile(path, ebiten.FilterNearest)
	return img
}

func (p Piesa) Move(tabla *[8][8]Piesa, x, y int) {
	switch p.tip {
	case 'K':
		p.miscareRege(tabla, x, y)
	case 'P':
		p.miscarePion(tabla, x, y)
	case 'B':
		p.miscareNebun(tabla, x, y)
	case 'N':
		p.miscareCal(tabla, x, y)
	case 'R':
		p.miscareTura(tabla, x, y)
	case 'Q':
		p.miscareRegina(tabla, x, y)
	default:
		return
	}
}
