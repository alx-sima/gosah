package piese

import (
	"fmt"
	"math/rand"
	"time"

	"github.com/hajimehoshi/ebiten"
)

const (
	Width  = 1920       // Lungimea ecranului
	Height = 1080       // Latimea ecranului
	L      = Height / 8 // Latura unui patrat
)

var (
	Board    [8][8]Piesa  // Board retine tabla de joc
	Selected PozitiePiesa //Selected retine piesa pe care s-a dat click
	Clicked  bool         // Clicked retine daca fost dat click pe o piesa
	Changed  bool         // Changed retine daca trebuie (re)desenat ecranul
	SahAlb   bool
	SahNegru bool
	Turn     rune // Turn retine 'W' daca e randul albului, sau 'B' daca e randul negrului
)

// Returneaza patratul in care se afla mouse-ul
func GetSquare() (int, int) {
	j, i := ebiten.CursorPosition()
	i, j = i/L, j/L
	Changed = true
	if i < 0 || j < 0 || i > 7 || j > 7 {
		panic("mouse in afara tablei")
	}
	return i, j
}
func Cronometru() {
	// TODO: adaugat cronometru
	for sec := 10; sec > 0; sec-- {
		fmt.Println(sec)
		time.Sleep(1 * time.Second)
	}
	fmt.Println("Ai ramas fara timp cioara")
}

func AfisarePatrateAtacate(x, y int) {
	if Board[x][y].Culoare == Turn {
		Board[x][y].Move(&Board, x, y, true, SahNegru || SahAlb)
		Selected = PozitiePiesa{Ref: &Board[x][y], X: x, Y: y}
	}
}

// Mutare muta piesa selectata pe pozitia ceruta
func Mutare() {
	if x, y := GetSquare(); Board[x][y].Atacat {
		Changed = true

		// Translateaza piesa din selected pe pozitia (x, y)
		Board[x][y] = *Selected.Ref
		Board[x][y].Mutat = true

		// Stergere pozitia initial selectat
		Board[Selected.X][Selected.Y] = Empty()
		Selected = PozitiePiesa{}

		// Transforma pionul in regina cand ajunge la capat
		if Board[x][y].Tip == 'P' {
			if Board[x][y].Culoare == 'W' && x == 0 {
				Board[x][y].Tip = 'Q'
			}
			if Board[x][y].Culoare == 'B' && x == 7 {
				Board[x][y].Tip = 'Q'
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

		// Reseteaza tabla de sah si de pozitii atacate
		SahAlb, SahNegru = false, false
		Clear(&Board)

		// Schimba tura de joc
		if Turn == 'W' {
			// Verifica daca regele negru e in sah
			ctrlRege := Board[RegeNegru.X][RegeNegru.Y].Control
			if ctrlRege == 1 || ctrlRege == 3 {
				SahNegru = true
			}
			Turn = 'B'
		} else {
			// Verifica daca regele alb e in sah
			ctrlRege := Board[RegeAlb.X][RegeAlb.Y].Control
			if ctrlRege == 2 || ctrlRege == 3 {
				SahAlb = true
			}
			Turn = 'W'
		}
	}
}

// InitializareMatrice genereaza piesele aleatoare pt. tabla de joc
func InitializareMatrice() {
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
