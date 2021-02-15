package piese

// returneaza daca regele de la (x, y) poate face rocada la (x, y + n)
func verifRocada(x, y, n int) bool {
	// Daca regele sau tura au fost mutate, rocada nu e posibila
	if Board[x][y+n].Mutat {
		return false
	}
	// Daca piesa de la (m, n) nu e o tura, rocada nu e posibila
	if Board[x][y+n].Tip != 'R' {
		return false
	}
	// semn retine 1 daca n e pozitiv, -1 daca nu (pentru a cauta in ambele directii)
	var semn int
	if n >= 0 {
		semn = 1
	} else {
		semn = -1
	}
	// Verifica daca calea de la rege la tura e goala
	for i := 1; i < semn*n; i++ {
		if Board[x][y+semn*i].Tip != 0 {
			return false
		}
	}
	// Verifica daca regele nu va ajunge in sah
	for i := 0; i < semn*3; i++ {
		if Board[x][y].Culoare == 'W' {
			if Board[x][y+semn*i].eControlatDeCuloare('B') {
				return false
			}
		} else if Board[x][y].Culoare == 'B' {
			if Board[x][y+semn*i].eControlatDeCuloare('W') {
				return false
			}
		}
	}
	return true
}
