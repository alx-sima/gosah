package piese

import (
	"fmt"
	"math/rand"
	"time"

	"github.com/hajimehoshi/ebiten"
)

const (
	Width  = 1920       // Width retine lungimea ecranului
	Height = 1080       // Height retine inaltimea ecranului
	L      = Height / 8 // L retine latura unui patrat
)

var (
	Board    [8][8]Piesa  // Board retine tabla de joc
	Selected PozitiePiesa //Selected retine piesa pe care s-a dat click
	Clicked  bool         // Clicked retine daca fost dat click pe o piesa
	Changed  bool         // Changed retine daca trebuie (re)desenat ecranul
	SahAlb   bool         // SahAlb retine daca regele alb e in sah
	SahNegru bool         // SahNegru retine daca regele negru e in sah
	Turn     rune         // Turn retine 'W' daca e randul albului, sau 'B' daca e randul negrului
)

// GetSquare returneaza patratul in care se afla mouse-ul
func GetSquare() (int, int) {
	j, i := ebiten.CursorPosition()
	i, j = i/L, j/L
	Changed = true
	if i < 0 || j < 0 || i > 7 || j > 7 {
		panic("mouse in afara tablei")
	}
	return i, j
}

// Cronometru ar trebui sa numere 10 minute pt fiecare jucator
func Cronometru() {
	// TODO: adaugat cronometru
	for sec := 10; sec > 0; sec-- {
		fmt.Println(sec)
		time.Sleep(1 * time.Second)
	}
	fmt.Println("Ai ramas fara timp cioara")
}

// AfisarePatrateAtacate genereaza mutarile posibile pentru piesa din (x, y) si o memoreaza in Selected
func AfisarePatrateAtacate(x, y int) {
	if Board[x][y].Culoare == Turn {
		Board[x][y].Move(&Board, x, y, true, SahNegru || SahAlb)
		Selected = PozitiePiesa{Ref: &Board[x][y], X: x, Y: y}
	}
}

// Mutare muta piesa selectata pe pozitia ceruta
func Mutare() {
	if x, y := GetSquare(); Board[x][y].Atacat {

		// Translateaza piesa din selected pe pozitia (x, y)
		Board[x][y] = *Selected.Ref
		Board[x][y].Mutat = true
		// Verifica daca mutarea provoaca o rocada
		if Board[x][y].Tip == 'K' {
			// In dreapta
			if y-Selected.Y == 2 {
				Board[x][y-1], Board[x][y+1] = Board[x][y+1], Board[x][y-1]
				// In stanga
			} else if Selected.Y-y == 2 {
				Board[x][y+1], Board[x][y-2] = Board[x][y-2], Board[x][y+1]
			}
		}
		// IMPORTANT!  aceasta verificare pentru pion trebuie facuta inainte de clear
		// Daca piesa captureaza prin en passant, elimina piesa capturata de pe tabla
		if Board[x][y].Tip == 'P' {
			if inBound(x-1, y) {
				if Board[x-1][y].EnPassant && Selected.X-x == -1 && (Selected.Y-y == 1 || Selected.Y-y == -1) {
					Board[x-1][y] = Empty()
				}
			}
			if inBound(x+1, y) {
				if Board[x+1][y].EnPassant && Selected.X-x == 1 && (Selected.Y-y == 1 || Selected.Y-y == -1) {
					Board[x+1][y] = Empty()
				}
			}

			// Transforma pionul in regina cand ajunge la capat
			if Board[x][y].Culoare == 'W' && x == 0 {
				Board[x][y].Tip = 'Q'
			}
			if Board[x][y].Culoare == 'B' && x == 7 {
				Board[x][y].Tip = 'Q'
			}
		}
		// Stergere pozitia initial selectat
		Board[Selected.X][Selected.Y] = Empty()
		Selected = PozitiePiesa{}

		// Reseteaza tabla de sah si de pozitii atacate
		SahAlb, SahNegru = false, false
		Clear(&Board, true)

		// IMPORTANT!  aceasta verificare pentru pion trebuie facuta dupa de clear
		if Board[x][y].Tip == 'P' {

			// Daca pionul s-a mutat 2 patratele, retine ca e apt pt. en passant
			if Selected.X-x == 2 || Selected.X-x == -2 {
				Board[x][y].EnPassant = true
			}
		}

		// Ia pozitia regelui
		if Board[x][y].Tip == 'K' {
			if Board[x][y].Culoare == 'W' {
				RegeAlb = PozitiePiesa{Ref: &Board[x][y], X: x, Y: y}
			}
			if Board[x][y].Culoare == 'B' {
				RegeNegru = PozitiePiesa{Ref: &Board[x][y], X: x, Y: y}
			}
		}

		// Schimba tura de joc
		if Turn == 'W' {
			// Verifica daca regele negru e in sah
			if Board[RegeNegru.X][RegeNegru.Y].eControlatDeCuloare('W') {
				SahNegru = true
			}
			Turn = 'B'
		} else {
			// Verifica daca regele alb e in sah
			if Board[RegeAlb.X][RegeAlb.Y].eControlatDeCuloare('B') {
				SahAlb = true
			}
			Turn = 'W'
		}

		Changed = true
	}
}

// InitializareMatriceRandomOglindit genereaza piesele aleatoare pt. tabla de joc
func InitializareMatriceRandomOglindit() {
	// Initializeaza regii
	Board[0][4] = NewPiesa('K', 'B')
	RegeNegru = PozitiePiesa{Ref: &Board[0][4], Y: 4}
	Board[7][4] = NewPiesa('K', 'W')
	RegeAlb = PozitiePiesa{Ref: &Board[7][4], X: 7, Y: 4}

	// Initializeaza seedul rand-ului
	rand.Seed(time.Now().Unix())

	for i := 0; i < 2; i++ {
		for j := 0; j < 8; j++ {
			// Genereaza piese aleatoriu (mai putin pe pozitia regilor)
			if !(i == 0 && j == 4) {
				r := rand.Int()
				switch r % 5 {
				case 0:
					// Pion
					Board[i][j], Board[7-i][j] = NewPiesa('P', 'B'), NewPiesa('P', 'W')
				case 1:
					// Nebun
					Board[i][j], Board[7-i][j] = NewPiesa('B', 'B'), NewPiesa('B', 'W')
				case 2:
					// Cal
					Board[i][j], Board[7-i][j] = NewPiesa('N', 'B'), NewPiesa('N', 'W')
				case 3:
					// Tura
					Board[i][j], Board[7-i][j] = NewPiesa('R', 'B'), NewPiesa('R', 'W')
				case 4:
					// Regina
					Board[i][j], Board[7-i][j] = NewPiesa('Q', 'B'), NewPiesa('Q', 'W')
				}
			}
		}
	}

	// FIXME: FOR TESTING PURPOSES
	for i := 0; i < 8; i++ {
		for j := 0; j < 8; j++ {
			fmt.Printf("%c ", Board[i][j].Tip)
		}
		fmt.Println()
	}
	fmt.Println("================")

	// FIXME: Cronometru
	// go cronometru()
}

// InitializareMatriceClasic initializeaza tabla unui joc clasic de sah
func InitializareMatriceClasic() {
	// Initializare
	piese := "RNBQKBNR"
	for i := 0; i < 8; i++ {
		Board[0][i] = NewPiesa(rune(piese[i]), 'B')
		Board[7][i] = NewPiesa(rune(piese[i]), 'W')
		Board[1][i] = NewPiesa('P', 'B')
		Board[6][i] = NewPiesa('P', 'W')
	}
	RegeNegru = PozitiePiesa{Ref: &Board[0][4], Y: 4}
	RegeAlb = PozitiePiesa{Ref: &Board[7][4], X: 7, Y: 4}
}

// returneaza daca regele de la (x, y) poate face rocada la (x, y + n)
func verifRocada(x, y, n int) bool {
	// Daca regele sau tura au fost mutate, rocada nu e posibila
	if Board[x][y+n].Mutat {
		return false
	}
	// Daca piesa de la (m, n) nu e o tura, rocada nu e posibila
	if Board[x][y+n].Tip != 'R' {
		return false
	}
	// semn retine 1 daca n e pozitiv, -1 daca nu (pentru a cauta in ambele directii)
	var semn int
	if n >= 0 {
		semn = 1
	} else {
		semn = -1
	}
	// Verifica daca calea de la rege la tura e goala
	// FIXME: merge doar spre dreapta
	for i := 1; i < semn*n; i++ {
		if Board[x][y+semn*i].Tip != 0 {
			return false
		}
	}
	// Verifica daca regele nu va ajunge in sah
	for i := 0; i < semn*3; i++ {
		if Board[x][y].Culoare == 'W' {
			if Board[x][y+semn*i].eControlatDeCuloare('B') {
				return false
			}
		} else if Board[x][y].Culoare == 'B' {
			if Board[x][y+semn*i].eControlatDeCuloare('W') {
				return false
			}
		}
	}
	return true
}

// eControlatDeCuloare verifica daca echipa culoare controleaza patratul dat
func (p *Piesa) eControlatDeCuloare(culoare rune) bool {
	if culoare == 'W' {
		return p.Control == 1 || p.Control == 3
	} else if culoare == 'B' {
		return p.Control == 2 || p.Control == 3
	}
	return false
}
