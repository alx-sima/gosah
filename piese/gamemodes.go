package piese

import (
	"math/rand"
	"time"
)

// init pentru package-ul piese
func init() {
	// Initializeaza seedul rand-ului
	rand.Seed(time.Now().Unix())

	// Albul incepe
	Turn = 'W'

	// Genereaza nivelele din fisiera
	Nivele = citireNivele()
}

// IncarcaNivel incarca nivel in Tabla
func IncarcaNivel(nivel string) {
	switch nivel {
	case "random":
		initializareMatriceRandomOglindit()
	case "editor":
		go editor()
		Editing = true
	// Daca nu gaseste in modurile prestabilite, verifica in folderul nivele
	default:
		initializareFisier(nivel)
	}

	// Daca nivelul e editor, cronometrul nu incepe
	if nivel != "editor" {
		TimpRamas.Alb.Min = 10
		TimpRamas.Negru.Min = 10
		go Cronometru()
	}
}

// initializareMatriceRandomOglindit genereaza piesele aleatoare pt. tabla de joc
func initializareMatriceRandomOglindit() {
	// Initializeaza regii
	generarePiesa(0, 4, 'K', 'B')
	generarePiesa(7, 4, 'K', 'W')

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
