package piese

import (
	"encoding/json"
	"fmt"
	"io/fs"
	"io/ioutil"
	"os"

	"strings"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
)

//data tine structura pentru fisierele .json
type data struct {
	// TODO: height retine inaltimea tablei
	Height int `json:"height"`
	// TODO: width retine latimea tablei
	Width int `json:"width"`
	//Tabla retine randurile tablei
	Tabla []string `json:"tabla"`
}

// initializareFisier initializeaza nivele speciale din fisierul clasic.json
func initializareFisier(fileName string) bool {
	// Deschide fisierul si verifica daca e valid
	f, _ := ioutil.ReadFile("nivele/" + fileName + ".json")
	if !json.Valid(f) {
		return false
	}

	// Desface fisierul
	var niv data
	if err := json.Unmarshal(f, &niv); err != nil {
		panic(err)
	}

	// Parcurge rand cu rand fieldul Tabla
	for i := 0; i < 8; i++ {
		for j := 0; j < 8; j++ {
			// Ia caracter cu caracter
			chr := rune(niv.Tabla[i][j])
			// Verifica daca e litera mica (inseamna ca e piesa neagra)
			if strings.ToLower(string(chr)) == string(chr) {
				generarePiesa(i, j, chr-'a'+'A', 'B')
				// Altfel piesa alba
			} else {
				generarePiesa(i, j, chr, 'W')
			}
		}
	}
	// Initializat cu succes
	return true
}

// listare cauta toate fisierele din folderul /nivele/ si le afiseaza fara extensia ".json"
func listare() {
	fmt.Println("?")
	fmt.Println("random")
	fmt.Println("editor")
	files, _ := os.ReadDir("./nivele")
	for _, f := range files {
		numeFisier := strings.ReplaceAll(f.Name(), ".json", "")
		fmt.Println(numeFisier)
	}
}
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
		// Daca apesi ESC salveaza si iese
		if inpututil.IsKeyJustPressed(ebiten.KeyEscape) {
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

			// Printeaza informatiile in format json (indentat)
			niv := data{8, 8, tabla}
			text, _ := json.MarshalIndent(niv, "", "\t")

			// TODO: alege numele fisierului cand salvezi
			os.WriteFile("nivele/edited.json", text, fs.ModePerm)

			// Inchide programul
			os.Exit(0)
		}

		// Click-stanga pune piese albe
		if inpututil.IsMouseButtonJustPressed(ebiten.MouseButtonLeft) {
			if x, y := GetSquare(); inBound(x, y) {
				if RegeAlb.Ref != nil && tip == 'K' {
				}
					generarePiesa(x, y, tip, 'W')
			
			}
		}

		// Click-dreapta pune piese negre
		if inpututil.IsMouseButtonJustPressed(ebiten.MouseButtonRight) {
			if x, y := GetSquare(); inBound(x, y) {
				if RegeNegru.Ref == nil {
					generarePiesa(x, y, tip, 'B')
				}
			}
		}

		// Click-rotita sterge piese
		if inpututil.IsMouseButtonJustPressed(ebiten.MouseButtonMiddle) {
			if x, y := GetSquare(); inBound(x, y) {
				if Board[x][y].Tip == 'K' {
					if Board[x][y].Culoare == 'W' {
						RegeAlb = PozitiePiesa{}
					} else {
						RegeNegru = PozitiePiesa{}
					}
				}
				Board[x][y] = Empty()
			}
		}
	}
}
