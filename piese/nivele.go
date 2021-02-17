package piese

import (
	"encoding/json"
	"io/ioutil"

	"strings"
)

//data tine structura pentru fisierele .json
type data struct {
	height, width int      //TODO: height si width retin dimensiunile tablei
	Tabla         []string //Tabla retine randurile tablei
}

// initializareFisier initializeaza nivele speciale din fisierul nivel.json
func initializareFisier() {
	// Deschide fisierul
	f, _ := ioutil.ReadFile("nivel.json")

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
				generarePiesa(i, j, rune(chr-'a'+'A'), 'B')
			// Altfel piesa alba	
			} else {
				generarePiesa(i, j, rune(chr), 'W')
			}
		}
	}
}
