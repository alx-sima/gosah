package piese

import (
	"fmt"

	"github.com/hajimehoshi/ebiten"
	"github.com/hajimehoshi/ebiten/ebitenutil"
)

// Piesa tine informatii despre un patrat de pe tabla
type Piesa struct {
	Atacat    bool // Atacat retine daca in acel patrat poate ajunge piesa selectata (util.Selected)
	Mutat     bool // Mutat retine daca piesa a fost mutata pana acum
	EnPassant bool // EnPassant retine daca ultima miscare a pionului a fost de 2 patrate, astfel incat sa fie posibila capturarea prin en passant
	Tip       rune // Tip retine initiala piesei (in engleza)
	Culoare   rune // Culoare: W inseamna piesa alba, B inseamna piesa neagra
	Control   int  // Control: 0 inseamna ca nu e controlat de nimeni acel patrat; 1 inseamna ca e controlat de alb, 2 inseamna ca e controlat de negru, 3 inseamna ca e controlat de ambele
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
	tipCorect := func(x rune) bool{
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

// DrawPiece returneaza imaginea piesei ce trebuie desenata, nil daca nu gaseste nimc
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

// inBound verifica daca pozitia se afla pe tabla
func inBound(a, b int) bool {
	return a >= 0 && b >= 0 && a < 8 && b < 8
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

// verifIesireSah verifica daca exista miscare care scoate regele din sah
func verifIesireSah(tabla *[8][8]Piesa, x, y, m, n int) {

	if !verifSah(tabla, x, y, m, n) {
		tabla[m][n].Atacat = true
	}
}

// verifSah verifica daca mutarea piesei alese din patratul x, y in patratul m, n scoate regele din sah. Returneaza true daca ramane in sah, false daca iese din sah
func verifSah(tabla *[8][8]Piesa, x, y, m, n int) bool {
	// Muta piesa pe patratul (m, n) (temporar)
	piesa, culoare, ok := tabla[m][n].Tip, tabla[m][n].Culoare, false

	tabla[m][n].Tip = tabla[x][y].Tip
	tabla[m][n].Culoare = tabla[x][y].Culoare
	tabla[x][y].Tip = 0
	tabla[x][y].Culoare = 0

	// Resetam matricea care arata controlul fiecarui patrat
	verifPatrateAtacate(tabla)

	if tabla[m][n].Tip == 'K' {
		if tabla[m][n].Culoare == 'B' {
			if tabla[m][n].eControlatDeCuloare('W') {
				ok = true
			}
		} else {
			if tabla[m][n].eControlatDeCuloare('B') {
				ok = true
			}
		}
	} else {
		// Daca regele nu se mai afla in sah, notam mutarea efectuata drept posibila
		ctrlRegeNegru := tabla[RegeNegru.X][RegeNegru.Y]
		ctrlRegeAlb := tabla[RegeAlb.X][RegeAlb.Y]

		if tabla[m][n].Culoare == 'B' {
			if ctrlRegeNegru.eControlatDeCuloare('W') {
				ok = true
			}
		} else {
			if ctrlRegeAlb.eControlatDeCuloare('B') {
				ok = true
			}
		}
	}

	// Punem piesa inapoi unde era
	tabla[x][y].Tip = tabla[m][n].Tip
	tabla[x][y].Culoare = tabla[m][n].Culoare
	tabla[m][n].Tip = piesa
	tabla[m][n].Culoare = culoare

	// Resetam matricea care arata controlul fiecarui patrat la starea originala
	verifPatrateAtacate(tabla)
	return ok
}
