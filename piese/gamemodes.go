package piese

import (
	"fmt"
	"math/rand"
	"time"
)

// Init implementeaza gamemodeul
func Init() {
	// Albul incepe
	Turn = 'W'

	// Repeta pana cand se incarca un nivel valid
	fmt.Println("Ce nivel vei juca? (default: random, editor: deseneaza nivel, ?: listare variante)")
	for nivel, invalid := "random", true; invalid; {

		fmt.Scanf("%s", &nivel)
		switch nivel {
		case "?":
			listare()
		case "random":
			initializareMatriceRandomOglindit()
			invalid = false
		case "editor":
			go editor()
			Editing = true
			invalid = false
		// Daca nu gaseste in modurile prestabilite, verifica in folderul nivele
		default:
			if initializareFisier(nivel) {
				invalid = false
			} else {
				fmt.Println("format gresit")
			}
		}
	}
	// FIXME: Cronometru
	// go cronometru()
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
