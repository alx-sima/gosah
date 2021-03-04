package game

import (
	"fmt"
	"gosah/piese"
	"os"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
)

// Update proceeds the game state.
// Update is called every tick (1/60 [s] by default).
func (g *game) Update() error {
	// Daca se apasa escape, jocul se termina
	if inpututil.IsKeyJustPressed(ebiten.KeyEscape) {
		os.Exit(0)
	}
	if inpututil.IsKeyJustPressed(ebiten.KeyH) {
		helping = !helping
		if helping {
			updateHelpText()
		}
	}

	// Daca jocul nu este inceput, se selecteaza nivelul
	if !piese.Started {
		// -1: muta la stanga, 0: selecteaza, 1: muta la dreapta, 2: default
		var mutare = 2

		// KeyLeft muta la stanga
		if inpututil.IsKeyJustPressed(ebiten.KeyLeft) {
			mutare = -1
			// KeyEnter selecteaza
		} else if inpututil.IsKeyJustPressed(ebiten.KeyEnter) {
			mutare = 0
			// KeyRight muta la dreapta
		} else if inpututil.IsKeyJustPressed(ebiten.KeyRight) {
			mutare = 1
			// MouseButtonLeft schimba nivelul in functie de pozitia mouseului
		} else if inpututil.IsMouseButtonJustPressed(ebiten.MouseButtonLeft) {
			_, _, mutare = piese.GetSquare()
		}

		switch mutare {
		// Schimba nivelul la stanga
		case -1:
			if selected > 0 {
				selected--
			}

		// Porneste nivelul
		case 0:
			piese.IncarcaNivel(piese.Nivele[selected])
			piese.Started = true

		// Schimba nivelul la dreapta
		case 1:
			if selected < len(piese.Nivele)-1 {
				selected++
			}
		}
		return nil
	}

	// Daca jocul este in modul editing, Update este overrided de piese.editor
	if piese.Editing {
		return nil
	}

	// Daca jocul este in pat, se termina
	if piese.Pat {
		piese.Turn = 'X'
		fmt.Println("Ai egalat")
		return nil
	}

	if inpututil.IsMouseButtonJustReleased(ebiten.MouseButtonLeft) {
		x, y, err := piese.GetSquare()
		if err != 0 {
			return nil
		}
		// Daca apesi pe un patrat atacat apelezi functia de mutare
		if piese.Board[x][y].Atacat == true {
			piese.Mutare()
			piese.Clicked = false
		}
		// Daca ultimul clic a fost pe o piesa, se reseteaza tabla inainte de a inregistra clicul curent
		if piese.Clicked {
			piese.Clear(&piese.Board, false)
		}
		// Daca clicul a fost pe o piesa afiseaza patratele pe care se poate misca
		if piese.Board[x][y].Tip != 0 {
			piese.Clicked = true
		} else {
			piese.Clicked = false
		}

		if piese.Clicked {
			piese.AfisarePatrateAtacate(x, y)
		}
	}
	return nil
}
