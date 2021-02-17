package piese

import (
	"encoding/json"
	"fmt"
	"io/ioutil"

	"strings"
)

//data tine structura pentru fisierele .json
type data struct {
	height, width int      //TODO: height si width retin dimensiunile tablei
	Tabla         []string //Tabla retine randurile tablei
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
	files, _ := ioutil.ReadDir("./nivele")
	for _, f := range files {
		numeFisier := strings.ReplaceAll(f.Name(), ".json", "")
		fmt.Println(numeFisier)
	}
}
