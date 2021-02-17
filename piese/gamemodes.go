package piese

import (
	"math/rand"
	"os"
	"time"
)

// Init implementeaza gamemodeul
func Init(mod string) {
	Turn = 'W'
	// Cauta fisierul nivel.json
	_, err := os.Stat("nivel.json")
	// Daca nu exista, initiem normal
	if os.IsNotExist(err) {
		switch mod {
		case "clasic":
			initializareMatriceClasic()
		case "random":
			initializareMatriceRandomOglindit()
		case "sandbox":
			initializareMatriceSandbox()
		}

		// FIXME: Cronometru
		// go cronometru()
	// Daca nu exista erori, initializam nivelul din fisier
	} else if err == nil {
		initializareFisier()
	}
}

// initializareMatriceRandomOglindit genereaza piesele aleatoare pt. tabla de joc
func initializareMatriceRandomOglindit() {
	// Initializeaza regii
	Board[0][4] = NewPiesa('K', 'B')
	RegeNegru = PozitiePiesa{Ref: &Board[0][4], Y: 4}
	PieseNegre = append(PieseNegre, 'K')

	Board[7][4] = NewPiesa('K', 'W')
	RegeAlb = PozitiePiesa{Ref: &Board[7][4], X: 7, Y: 4}
	PieseAlbe = append(PieseAlbe, 'K')

	// Initializeaza seedul rand-ului
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
					PieseAlbe = append(PieseAlbe, 'P')
					PieseNegre = append(PieseNegre, 'P')
				case 1:
					// Nebun
					Board[i][j], Board[7-i][j] = NewPiesa('B', 'B'), NewPiesa('B', 'W')
					PieseAlbe = append(PieseAlbe, 'B')
					PieseNegre = append(PieseNegre, 'B')
				case 2:
					// Cal
					Board[i][j], Board[7-i][j] = NewPiesa('N', 'B'), NewPiesa('N', 'W')
					PieseAlbe = append(PieseAlbe, 'N')
					PieseNegre = append(PieseNegre, 'N')
				case 3:
					// Tura
					Board[i][j], Board[7-i][j] = NewPiesa('R', 'B'), NewPiesa('R', 'W')
					PieseAlbe = append(PieseAlbe, 'R')
					PieseNegre = append(PieseNegre, 'R')
				case 4:
					// Regina
					Board[i][j], Board[7-i][j] = NewPiesa('Q', 'B'), NewPiesa('Q', 'W')
					PieseAlbe = append(PieseAlbe, 'Q')
					PieseNegre = append(PieseNegre, 'Q')
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
	PieseAlbe = append(PieseAlbe, 'R', 'N', 'B', 'Q', 'K', 'B', 'N', 'R', 'P', 'P', 'P', 'P', 'P', 'P', 'P', 'P')
	PieseNegre = append(PieseNegre, 'R', 'N', 'B', 'Q', 'K', 'B', 'N', 'R', 'P', 'P', 'P', 'P', 'P', 'P', 'P', 'P')
}

// FIXME: will be removed
func initializareMatriceSandbox() {
	Board[0][4] = NewPiesa('K', 'B')
	RegeNegru = PozitiePiesa{Ref: &Board[0][4], Y: 4}
	PieseNegre = append(PieseNegre, 'K')

	Board[7][4] = NewPiesa('K', 'W')
	RegeAlb = PozitiePiesa{Ref: &Board[7][4], X: 7, Y: 4}
	PieseAlbe = append(PieseAlbe, 'K')

	Board[7][5] = NewPiesa('N', 'W')
	RegeAlb = PozitiePiesa{Ref: &Board[7][4], X: 7, Y: 4}
	PieseAlbe = append(PieseAlbe, 'N')

	/*Board[0][3] = NewPiesa('N', 'B')
	RegeNegru = PozitiePiesa{Ref: &Board[0][4], Y: 4}
	PieseNegre = append(PieseNegre, 'N')*/
}

// initializare nivele speciale din fisierul nivel.json
func initializareFisier() {
	
}

