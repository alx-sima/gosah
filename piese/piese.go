package piese

import (
	"github.com/hajimehoshi/ebiten"
	"github.com/hajimehoshi/ebiten/ebitenutil"
)

type Piesa struct {
	Atacat, Mutat bool
	Tip, Culoare  rune
	Control       int // 0 inseamna ca nu e controlat de nimeni acel patrat; 1 inseamna ca e controlat de alb, 2 inseamna ca e controlat de negru, 3 inseamna ca e controlat de ambele
}

// Constructori
func NewPiesa(tip, culoare rune) Piesa {
	e := Piesa{false, false, tip, culoare, 0}
	return e
}

func Empty() Piesa {
	e := Piesa{false, false, 0, 0, 0}
	return e
}

// Metode

// returneaza imaginea piesei ce trebuie desenata, nil daca nu gaseste nimc
func (p *Piesa) DrawPiece() *ebiten.Image {
	path := "imagini/"
	if p.Culoare == 'W' {
		path += "w_"
	} else {
		path += "b_"
	}
	switch p.Tip {
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

// Verifica pe ce patrate se poate misca o anumita piesa
func (p Piesa) Move(tabla *[8][8]Piesa, x, y int) {
	switch p.Tip {
	case 'K':
		p.miscareRege(tabla, x, y, 'M')
	case 'P':
		p.miscarePion(tabla, x, y, 'M')
	case 'B':
		p.miscareNebun(tabla, x, y, 'M')
	case 'N':
		p.miscareCal(tabla, x, y, 'M')
	case 'R':
		p.miscareTura(tabla, x, y, 'M')
	case 'Q':
		p.miscareRegina(tabla, x, y, 'M')
	default:
		return
	}
}
