package piese

import (
	"fmt"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
	"strings"
)

// editor este functie asincrona care implementeaza modul de editare tabla de joc
// game.Update trebuie ignorat pentru a nu fi conflicte
func editor() {
	// Default e pionul
	tip := 'P'
	for {
		// R pentru tura
		if inpututil.IsKeyJustPressed(ebiten.KeyR) {
			tip = 'R'
		}
		// N pentru cal
		if inpututil.IsKeyJustPressed(ebiten.KeyN) {
			tip = 'N'
		}
		// B pentru nebun
		if inpututil.IsKeyJustPressed(ebiten.KeyB) {
			tip = 'B'
		}
		// Q pentru regina
		if inpututil.IsKeyJustPressed(ebiten.KeyQ) {
			tip = 'Q'
		}
		// K pentru rege
		if inpututil.IsKeyJustPressed(ebiten.KeyK) {
			tip = 'K'
		}
		// P pentru pion
		if inpututil.IsKeyJustPressed(ebiten.KeyP) {
			tip = 'P'
		}
		// Salveaza global tipul selectat
		TipSelectat = tip

		// Daca apesi Ctrl+S salveaza si iese
		if inpututil.KeyPressDuration(ebiten.KeyControl) > 0 && inpututil.KeyPressDuration(ebiten.KeyS) > 0 {
			// Verifica daca nivelul in starea curenta e valid
			if !checkSave() {
				continue
			}

			var tabla []string

			for i := 0; i < 8; i++ {
				var rand string
				for j := 0; j < 8; j++ {

					// Daca piesa nu exista, printeaza ' '
					if Board[i][j].Tip == 0 {
						rand += " "
						continue
					}

					piesa := string(Board[i][j].Tip)

					// Daca piesa e neagra, o printeaza cu litera mica
					if Board[i][j].Culoare == 'B' {
						piesa = strings.ToLower(piesa)
					}
					rand += piesa
				}
				tabla = append(tabla, rand)
			}

			saveToJson(data{
				Height: 8,
				Width:  8,
				Tabla:  tabla,
			})
		}

		// Click-stanga pune piese albe
		if inpututil.IsMouseButtonJustPressed(ebiten.MouseButtonLeft) {
			if x, y, err := GetSquare(); err == 0 {
				deletePiesa(x, y)

				if RegeAlb.Ref == nil || tip != 'K' {
					generarePiesa(x, y, tip, 'W')
				}
			}
		}

		// Click-dreapta pune piese negre
		if inpututil.IsMouseButtonJustPressed(ebiten.MouseButtonRight) {
			if x, y, err := GetSquare(); err == 0 {
				deletePiesa(x, y)

				if RegeNegru.Ref == nil || tip != 'K' {
					generarePiesa(x, y, tip, 'B')
				}
			}
		}

		// Click-rotita sterge piese
		if inpututil.IsMouseButtonJustPressed(ebiten.MouseButtonMiddle) {
			if x, y, err := GetSquare(); err == 0 {
				deletePiesa(x, y)
			}
		}
	}
}

// deletePiesa sterge piesa de pe pozitia x, y
func deletePiesa(x, y int) {
	if Board[x][y].Culoare == 'W' {
		ramaseAlbe.edit(Board[x][y].Tip, -1)
		if Board[x][y].Tip == 'K' {
			RegeAlb = PozitiePiesa{}
		}
	} else {
		ramaseNegre.edit(Board[x][y].Tip, -1)
		if Board[x][y].Tip == 'K' {
			RegeNegru = PozitiePiesa{}
		}
	}
	Board[x][y] = Empty()
}

// checkSave verifica daca nivelul este valid pt salvare
func checkSave() bool {
	// Trebuie sa existe ambii regi
	if RegeAlb.Ref == nil || RegeNegru.Ref == nil {
		fmt.Println("nu se poate salva: nu exista ambii regi")
		return false
	}
	// Trebuie sa nu fie pat
	if VerifPat(); Pat {
		fmt.Println("nu se poate salva: pat")
		return false
	}
	// FIXME
	// Trebuie sa nu fie mat
	return true
}
