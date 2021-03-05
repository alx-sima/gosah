package piese

import (
	"encoding/json"
	"io/ioutil"
	"os"
	"strings"
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
func initializareFisier(fileName string) {
	// Deschide fisierul si verifica daca e valid
	f, _ := ioutil.ReadFile("nivele/" + fileName + ".json")
	if !json.Valid(f) {
		panic("fisierul este invalid")
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
	return
}

// citireNivele cauta toate fisierele din folderul /nivele/ si le afiseaza fara extensia ".json"
func citireNivele() (nivele []string) {
	// Declaratii
	nivele = append(nivele, "random")
	files, _ := os.ReadDir("./nivele")

	// Parcurge fisierele din "./nivele"
	for _, f := range files {
		numeFisier := strings.ReplaceAll(f.Name(), ".json", "")
		nivele = append(nivele, numeFisier)
	}

	nivele = append(nivele, "editor")
	return
}

// saveToJson printeaza informatiile in format json (indentat)
func saveToJson(niv data) {
	text, _ := json.MarshalIndent(niv, "", "\t")
	// TODO: alege numele fisierului cand salvezi
	if err := os.WriteFile("nivele/custom.json", text, 0777); err != nil {
		panic(err)
	}

	// Inchide programul
	os.Exit(0)
}
