package piese

import (
	"bytes"
	images "gosah/imagini"
	"image"

	// importat ca _ ca altfel nu merge
	_ "image/png"
	"log"

	"github.com/hajimehoshi/ebiten/v2"
)

// Piesa tine informatii despre un patrat de pe tabla
type Piesa struct {
	// Atacat retine daca in acel patrat poate ajunge piesa selectata (util.Selected)
	Atacat bool
	// Mutat retine daca piesa a fost mutata pana acum
	Mutat bool
	// EnPassant retine daca ultima miscare a pionului a fost de 2 patrate, astfel incat sa fie posibila capturarea prin en passant
	EnPassant bool
	// Tip retine initiala piesei (in engleza)
	Tip rune
	// Culoare: W inseamna piesa alba, B inseamna piesa neagra
	Culoare rune
	// Control: 0 inseamna ca nu e controlat de nimeni acel patrat; 1 inseamna ca e controlat de alb, 2 inseamna ca e controlat de negru, 3 inseamna ca e controlat de ambele
	Control int
}

// PozitiePiesa tine piesa si pozitia ei
type PozitiePiesa struct {
	// Referinta la piesa memorata
	Ref *Piesa
	// Pozitia piesei pe tabla
	X, Y int
}

/// Constructori

// NewPiesa returneaza o noua piesa, initializata
func NewPiesa(tip, culoare rune) Piesa {
	e := Piesa{false, false, false, tip, culoare, 0}
	return e
}

// Empty returneaza o noua piesa "goala"
func Empty() Piesa {
	e := Piesa{false, false, false, 0, 0, 0}
	return e
}

// generarePiesa adauga pe tabla si in vectorii de piese o noua piesa la pozitia (i, j), de <tip> si <culoare>
func generarePiesa(i, j int, tip, culoare rune) {
	// Functie anonima care verifica daca piesa care trebuie mentionata este valida
	tipCorect := func(x rune) bool {
		for _, i := range "RNBQKP" {
			if x == i {
				return true
			}
		}
		return false
	}
	if !tipCorect(tip) {
		return
	}
	Board[i][j] = NewPiesa(tip, culoare)
	if culoare == 'W' {
		PieseAlbe = append(PieseAlbe, tip)
		if tip == 'K' {
			RegeAlb = PozitiePiesa{&Board[i][j], i, j}
		}
	} else {
		PieseNegre = append(PieseNegre, tip)
		if tip == 'K' {
			RegeNegru = PozitiePiesa{&Board[i][j], i, j}
		}
	}
}

/// Metode

// loadImageFromBytes incarca din vectorul octeti imaginea si o returneaza
func loadImageFromBytes(octeti []byte) *ebiten.Image {
	img, _, err := image.Decode(bytes.NewReader(octeti))
	if err != nil {
		log.Fatal(err)
	}
	return ebiten.NewImageFromImage(img)
}

// DrawPiece returneaza imaginea piesei ce trebuie desenata, nil daca nu gaseste nimic
func (p *Piesa) DrawPiece() *ebiten.Image {
	if !Started {
		return nil
	} 
	switch p.Tip {
	case 'K':
		if p.Culoare == 'W' {
			return loadImageFromBytes(images.WhiteKing)
		}
		return loadImageFromBytes(images.BlackKing)
	case 'P':
		if p.Culoare == 'W' {
			return loadImageFromBytes(images.WhitePawn)
		}
		return loadImageFromBytes(images.BlackPawn)
	case 'B':
		if p.Culoare == 'W' {
			return loadImageFromBytes(images.WhiteBishop)
		}
		return loadImageFromBytes(images.BlackBishop)
	case 'N':
		if p.Culoare == 'W' {
			return loadImageFromBytes(images.WhiteKnight)
		}
		return loadImageFromBytes(images.BlackKnight)
	case 'R':
		if p.Culoare == 'W' {
			return loadImageFromBytes(images.WhiteRook)
		}
		return loadImageFromBytes(images.BlackRook)
	case 'Q':
		if p.Culoare == 'W' {
			return loadImageFromBytes(images.WhiteQueen)
		}
		return loadImageFromBytes(images.BlackQueen)
	default:
		return nil
	}
}

// Move verifica pe ce patrate se poate misca o anumita piesa si apeleaza functia corespunzatoare
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

// Clear curata tabla de miscari posibile + verifica pentru pozitiile actuale controlul asupra fieccarui patrat
func Clear(tabla *[8][8]Piesa, moved bool) {
	for i := 0; i < 8; i++ {
		for j := 0; j < 8; j++ {
			tabla[i][j].Atacat = false
			if moved {
				tabla[i][j].EnPassant = false
			}
		}
	}
	verifPatrateAtacate(tabla)
}

// setareControl seteaza controlul acelui patrat
func setareControl(patrat *Piesa, culoare rune) {
	// TODO: functia poate fi scurtata
	if culoare == 'W' {
		if patrat.Control == 2 {
			patrat.Control = 3
		} else if patrat.Control == 0 {
			patrat.Control = 1
		}
	} else {
		if patrat.Control == 1 {
			patrat.Control = 3
		} else if patrat.Control == 0 {
			patrat.Control = 2
		}
	}
}

// verifPatrateAtacate verifica pentru fiecare piesa ce patrate ataca si formeaza in tabla.Control o matrice care arata ce culoare controleaza fiecare patrat
func verifPatrateAtacate(tabla *[8][8]Piesa) {
	for i := 0; i < 8; i++ {
		for j := 0; j < 8; j++ {
			tabla[i][j].Control = 0
		}
	}
	for i := 0; i < 8; i++ {
		for j := 0; j < 8; j++ {
			tabla[i][j].Move(tabla, i, j, false, false)
		}
	}
}
