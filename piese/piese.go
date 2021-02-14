package piese

import (
	"fmt"
	"github.com/hajimehoshi/ebiten"
	"github.com/hajimehoshi/ebiten/ebitenutil"
)

// Piesa tine informatii despre un patrat de pe tabla
type Piesa struct {
	Atacat  bool // Atacat retine daca in acel patrat poate ajunge piesa selectata (util.Selected)
	Mutat   bool // Mutat retine daca piesa a fost mutata pana acum
	Tip     rune // Tip retine initiala piesei (in engleza)
	Culoare rune // Culoare: W inseamna piesa alba, B inseamna piesa neagra
	Control int  // Control: 0 inseamna ca nu e controlat de nimeni acel patrat; 1 inseamna ca e controlat de alb, 2 inseamna ca e controlat de negru, 3 inseamna ca e controlat de ambele
}

// PozitiePiesa tine piesa si pozitia ei
type PozitiePiesa struct {
	Ref  *Piesa // Referinta la piesa memorata
	X, Y int    // Pozitia piesei pe tabla
}

var (
	RegeNegru PozitiePiesa // Pozitia regelui negru
	RegeAlb   PozitiePiesa // Pozitia regelui alb
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
	var culoare, tip string
	if p.Culoare == 'W' {
		culoare = "w"
	} else {
		culoare = "b"
	}
	switch p.Tip {
	case 'K':
		tip = "king"
	case 'P':
		tip = "pawn"
	case 'B':
		tip = "bishop"
	case 'N':
		tip = "knight"
	case 'R':
		tip = "rook"
	case 'Q':
		tip = "queen"
	default:
		return nil
	}
	// Creeaza path-ul imaginii de randat
	path := fmt.Sprintf("imagini/%s_%s_png_128px.png", culoare, tip)
	img, _, _ := ebitenutil.NewImageFromFile(path, ebiten.FilterNearest)
	return img
}

// Verifica pe ce patrate se poate misca o anumita piesa si apeleaza functia corespunzatoare
func (p Piesa) Move(tabla *[8][8]Piesa, x, y int, mutare, isSah bool) {
	switch p.Tip {
	case 'K':
		p.miscareRege(tabla, x, y, mutare)
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
