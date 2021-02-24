package piese

import "fmt"

// verifInBound verifica daca pozitia se ala pe tabla
func verifInBound(a, b int) bool {
	return a >= 0 && b >= 0 && a < 8 && b < 8
}

// verifIesireSah verifica daca exista miscare care scoate regele din sah
func verifIesireSah(tabla *[8][8]Piesa, x, y, m, n int) {

	if !verifSah(tabla, x, y, m, n) {
		tabla[m][n].Atacat = true
	}
}

// verifSah verifica daca mutarea piesei alee din patratul x, y in patratul m, n scoate regele din sah. Returneaza true daca ramane in sah, false daca iese din sah
func verifSah(tabla *[8][8]Piesa, x, y, m, n int) bool {
	// Muta piesa pe patratul (m, n) (temporar)
	piesa, culoare, ok := tabla[m][n].Tip, tabla[m][n].Culoare, false

	tabla[m][n].Tip = tabla[x][y].Tip
	tabla[m][n].Culoare = tabla[x][y].Culoare
	tabla[x][y].Tip = 0
	tabla[x][y].Culoare = 0

	// Resetam matricea care arata controlul fiecarui patrat
	verifPatrateAtacate(tabla)

	if tabla[m][n].Tip == 'K' {
		if tabla[m][n].Culoare == 'B' {
			if tabla[m][n].eControlatDeCuloare('W') {
				ok = true
			}
		} else {
			if tabla[m][n].eControlatDeCuloare('B') {
				ok = true
			}
		}
	} else {
		// Daca regele nu se mai afla in sah, noteaza mutarea efectuata drept posibila
		ctrlRegeNegru := tabla[RegeNegru.X][RegeNegru.Y]
		ctrlRegeAlb := tabla[RegeAlb.X][RegeAlb.Y]

		if tabla[m][n].Culoare == 'B' {
			if ctrlRegeNegru.eControlatDeCuloare('W') {
				ok = true
			}
		} else {
			if ctrlRegeAlb.eControlatDeCuloare('B') {
				ok = true
			}
		}
	}

	// Punem piesa inapoi unde era
	tabla[x][y].Tip = tabla[m][n].Tip
	tabla[x][y].Culoare = tabla[m][n].Culoare
	tabla[m][n].Tip = piesa
	tabla[m][n].Culoare = culoare

	// Reseta matricea care arata controlul fiecarui patrat la starea originala
	verifPatrateAtacate(tabla)
	return ok
}

// VerifPat verifica toate conditiile de egalitate, stabilind daca jocul mai poate continua
func VerifPat() {
	// Daca dup 50 de miscari (alb + negru) nu se captureaza nicio piesa, meciul se termina in egal
	if MutariUltimaCapturare == 100 {
		Pat = true
	} else {
		// Daca nu xista material suficient pentru sah mat, meciul se termina in egal
		if len(PieseAlbe) <= 2 && len(PieseNegre) <= 2 {
			Pat = true
			// Cazul Rege + Nebun vs ege + Nebun. Daca nebunii se afla pe patrate de aceeasi culoare meciul se termina in egal
			if len(PieseAlbe) == len(PieseNegre) && len(PieseNegre) == 2 {
				for i := 0; i < 2; i++ {
					if PieseAlbe[i] != 'K' && PieseAlbe[i] != 'B' && PieseNegre[i] != 'K' && PieseNegre[i] != 'B' {
						Pat = false
						break
					}
				}
				if Pat == true {
					culoare := 0 // culoare rprezinta culoarea patratului pe care se afla nebunul
					for i := 0; i < 8; i++ {
						for j := 0; j < 8; j++ {
							if Board[i][j].Tip == 'B' {
								if culoare != 0 {
									if ((i+j)%2 == 0 && culoare != 'A') || ((i+j)%2 == 1 && culoare != 'N') {
										Pat = false
									}
								} else {
									if (i+j)%2 == 0 {
										culoare = 'A'
									} else {
										culoare = 'N'
									}
								}
							}
						}
					}
				}
				// Verifica celelalte 3 cazuri de insuficiena: Rege vs Rege, Rege + Cal vs Rege, Rege + Nebun vs Rege
			} else {
				for i := 0; i < len(PieseAlbe) && Pat; i++ {
					if PieseAlbe[i] != 'K' && PieseAlbe[i] != 'N' && PieseAlbe[i] != 'B' {
						Pat = false
					}
				}
				for i := 0; i < len(PieseNegre) && Pat; i++ {
					if PieseNegre[i] != 'K' && PieseNegre[i] != 'N' && PieseNegre[i] != 'B' {
						Pat = false
					}
				}
			}
		}
		// Daca nu exista misari legale posibile, meciul se termina in sah mat
		if Pat == false {
			existaMutare = false
			for i := 0; i < 8 && !existaMutare; i++ {
				for j := 0; j < 8 && !existaMutare; j++ {
					if Board[i][j].Culoare == Turn {
						Board[i][j].Move(&Board, i, j, false, true)
					}
				}
			}
			if existaMutare == false {
				Pat = true
			}
		}
	}
}

// eControlateCuloar verifica daca echipa culoare controleaza patratul dat
func (p *Piesa) eControlatDeCuloare(culoare rune) bool {
	if culoare == 'W' {
		return p.Control == 1 || p.Control == 3
	} else if culoare == 'B' {
		return p.Control == 2 || p.Control == 3
	}
	return false
}

// verifMatverifica toate piesele cautand miscari legale care sa te scoata din sah
func verifMat(culoare rune) {
	Mat = true
	for i := 0; i < 8 && Mat; i++ {
		for j := 0; j < 8 && Mat; j++ {
			if Board[i][j].Culoare == culoare {
				Board[i][j].Move(&Board, i, j, false, true)
			}
		}
	}
	if Mat {
		if culoare == 'W' {
			fmt.Print("Alb ")
		} else {
			fmt.Print("Negru ")

			fmt.Print("a pierdut")
		}
	}
}
