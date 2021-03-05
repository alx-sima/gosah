package game

import "gosah/piese"

// Selecteaza ce text trebuie afisat
func updateHelpText() {
	if piese.Editing {
		editingHelp()
	} else if piese.Started {
		meciHelp()
	} else {
		startHelp()
	}
}

func editingHelp() {
	helpText = `Click - amplaseaza piesa alba
Click dreapta - amplaseaza piesa neagra
Click mijloc - sterge piesa
ESC - iesi fara a salva
CTRL + S - iesi, salvand
h - help
Pentru a selecta un tip de piesa, apasa:
-> K - pt. rege
-> Q - pt. regina
-> B - pt. nebun
-> N - pt. cal
-> R - pt. tura
-> P - pt. pion
Imaginea din stanga jos arata ce piesa e selectata
`
}
func meciHelp() {
	helpText = `
Click pe o piesa pt. a o selecta si a arata ce patrate sunt disponibile (cu galben)
Click pe un patrat galben pt. a muta piesa selectata pe acea pozitie
Daca regele este intr-un patrat rosu, este in sah
In stanga sunt timerele, iar in dreapta istoricul miscarilor
ESC - inchidere joc
h - help
`
}
func startHelp() {
	helpText = `
Click pe tabla (sau ENTER) - incarca configuratia selectata
Click in stanga tablei (sau sageata-stanga) - selecteaza configuratia anterioara
Click in dreapta tablei (sau sageata-dreapta) - selecteaza configuratia posterioara
h - help
`
}
