package piese

import (
	"math/rand"
	"time"
)

// Init implementeaza gamemodeul
func Init(mod string) {
	Turn = 'W'
	switch mod {
	case "clasic":
		initializareMatriceClasic()
	case "random":
		initializareMatriceRandomOglindit()
	}

	// FIXME: Cronometru
	// go cronometru()
}

// initializareMatriceRandomOglindit genereaza piesele aleatoare pt. tabla de joc
func initializareMatriceRandomOglindit() {
	// Initializeaza regii
	Board[0][4] = NewPiesa('K', 'B')
	RegeNegru = PozitiePiesa{Ref: &Board[0][4], Y: 4}
	Board[7][4] = NewPiesa('K', 'W')
	RegeAlb = PozitiePiesa{Ref: &Board[7][4], X: 7, Y: 4}

	// Initializeaza sedul rand-ului
	rand.Seed(time.Now().Unix())

	for i := 0; i < 2; i++ {
		for j := 0; j < 8; j++ {
			// Generaza piese aleatoriu (mai putin pe pozitia regilor)
			if !(i == 0 && j == 4) {
				r := rand.Int()
				switch r % 5 {
				case 0:
					// Pion
					Board[i][j], Board[7-i][j] = NewPiesa('P', 'B'), NewPiesa('P', 'W')
				case 1:
					// Nebn
					Board[i][j], Board[7-i][j] = NewPiesa('B', 'B'), NewPiesa('B', 'W')
				case 2:
					// Cal
					Board[i][j], Board[7-i][j] = NewPiesa('N', 'B'), NewPiesa('N', 'W')
				case 3:
					// Tura
					Board[i][j], Board[7-i][j] = NewPiesa('R', 'B'), NewPiesa('R', 'W')
				case 4:
					// Reina
					Board[i][j], Board[7-i][j] = NewPiesa('Q', 'B'), NewPiesa('Q', 'W')
				}
			}
		}
	}
}

// InitializareMatriceClasic initilizeaza tabla unui joc clasic de sah
func initializareMatriceClasic() {
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

// eControlateCuloare verifica daca echipa culoare controleaza patratul dat
func (p *Piesa) eControlatDeCuloare(culoare rune) bool {
	if culoare == 'W' {
		return p.Control == 1 || p.Control == 3
	} else if culoare == 'B' {
		return p.Control == 2 || p.Control == 3
	}
	return false
}
