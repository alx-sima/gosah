package piese

import (
	"github.com/hajimehoshi/ebiten"
	"github.com/hajimehoshi/ebiten/ebitenutil"
)

type Piesa struct {
	x, y         int
	tip, culoare rune
}

func NewPiesa(x, y int, tip, culoare rune) Piesa {
	e := Piesa{x, y, tip, culoare}
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