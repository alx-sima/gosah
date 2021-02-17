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
	generarePiesa(0, 4, 'K', 'B')
	generarePiesa(7, 4, 'K', 'W')

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
					generarePiesa(i, j, 'P', 'B')
					generarePiesa(7-i, j, 'P', 'W')
				case 1:
					// Nebun
					generarePiesa(i, j, 'B', 'B')
					generarePiesa(7-i, j, 'B', 'W')
				case 2:
					// Cal
					generarePiesa(i, j, 'N', 'B')
					generarePiesa(7-i, j, 'N', 'W')
				case 3:
					// Tura
					generarePiesa(i, j, 'R', 'B')
					generarePiesa(7-i, j, 'R', 'W')
				case 4:
					// Regina
					generarePiesa(i, j, 'Q', 'B')
					generarePiesa(7-i, j, 'Q', 'W')
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
