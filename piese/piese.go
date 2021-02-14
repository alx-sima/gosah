package piese

import (
	"github.com/hajimehoshi/ebiten"
	"github.com/hajimehoshi/ebiten/ebitenutil"
)

// Piesa tine informatii despre un patrat de pe tabla
type Piesa struct {
	Atacat, Mutat bool
	Tip, Culoare  rune // Culoare: W inseamna piesa alba, B inseamna piesa neagra
	Control       int  // 0 inseamna ca nu e controlat de nimeni acel patrat; 1 inseamna ca e controlat de alb, 2 inseamna ca e controlat de negru, 3 inseamna ca e controlat de ambele
}

// PozitiePiesa tine piesa si pozitia ei
type PozitiePiesa struct {
	Ref  *Piesa // Referinta la piesa memorata
	X, Y int    // Pozitia piesei pe tabla
}

var (
	RegeNegru PozitiePiesa	// Pozitia regelui negru
	RegeAlb   PozitiePiesa	// Pozitia regelui alb
)

/// Constructori

// NewPiesa returneaza o noua piesa, initializata
func NewPiesa(tip, culoare rune) Piesa {
	e := Piesa{false, false, tip, culoare, 0}
	return e
}

// Empty eturneaza o noua piesa "goala"
func Empty() Piesa {
	e := Piesa{false, false, 0, 0, 0}
	return e
}

/// Metode

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
func (p Piesa) Move(tabla *[8][8]Piesa, x, y int, mutare, isSah bool) {
	switch p.Tip {
	case 'K':
		p.miscareRege(tabla, x, y, mutare, isSah)
	case 'P':
		p.miscarePion(tabla, x, y, mutare, isSah)
	case 'B':
		p.miscareNebun(tabla, x, y, mutare, isSah)
	case 'N':
		p.miscareCal(tabla, x, y, mutare, isSah)
	case 'R':
		p.miscareTura(tabla, x, y, mutare, isSah)
	case 'Q':
		p.miscareRegina(tabla, x, y, mutare, isSah)
	default:
		return
	}
}
