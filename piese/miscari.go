package piese

// inBound verifica daca pozitia se afla pe tabla
func inBound(a, b int) bool {
	return a >= 0 && b >= 0 && a < 8 && b < 8
}

// Clear curata tabla de miscari posibile + verifica pentru pozitiile actuale controlul asupra fieccarui patrat
func Clear(tabla *[8][8]Piesa) {
	for i := 0; i < 8; i++ {
		for j := 0; j < 8; j++ {
			tabla[i][j].Atacat = false
			tabla[i][j].Control = 0
		}
	}
	verifPatrateAtacate(tabla)
}

// setareControl seteaza controlul acelui patrat
func setareControl(patrat *Piesa, culoare rune) {
	if culoare == 'W' {
		if patrat.Control == 2 {
			patrat.Control = 3
		} else if patrat.Control == 0 {
			patrat.Control = 1
		}
	} else {
		if patrat.Control == 1 {
			patrat.Control = 3
		} else if patrat.Control == 0 {
			patrat.Control = 2
		}
	}
}

// verifPatrateAtacate verifica pentru fiecare piesa ce patrate ataca si formeaza in tabla.Control o matrice care arata ce culoare controleaza fiecare patrat
func verifPatrateAtacate(tabla *[8][8]Piesa) {
	for i := 0; i < 8; i++ {
		for j := 0; j < 8; j++ {
			tabla[i][j].Move(tabla, i, j, false, false)
		}
	}
}

// verifIesireSah verifica daca exista miscare care scoate regele din sah
func verifIesireSah(tabla *[8][8]Piesa, x, y, m, n int) {
	tabla[m][n].Tip = tabla[x][y].Tip
	tabla[m][n].Culoare = tabla[x][y].Culoare
	tabla[x][y].Tip = 0
	tabla[x][y].Culoare = 0
	verifPatrateAtacate(tabla)

	if (tabla[RegeNegru.X][RegeNegru.Y].Control == 1 || tabla[RegeNegru.X][RegeNegru.Y].Control == 3) && tabla[x][y].Culoare == 'B' {
		tabla[m][n].Atacat = true
	}
	if (tabla[RegeAlb.X][RegeAlb.Y].Control == 2 || tabla[RegeAlb.X][RegeAlb.Y].Control == 3) && tabla[x][y].Culoare == 'W' {
		tabla[m][n].Atacat = true
	}

	tabla[x][y].Tip = tabla[m][n].Tip
	tabla[x][y].Culoare = tabla[m][n].Culoare
	tabla[m][n].Tip = 0
	tabla[m][n].Culoare = 0
}

// miscareRege cauta miscarile posibile pt regele ales
func (p *Piesa) miscareRege(tabla *[8][8]Piesa, x, y int, mutare bool) {
	var dx = []int{x - 1, x - 1, x - 1, x, x + 1, x + 1, x + 1, x}
	var dy = []int{y - 1, y, y + 1, y + 1, y + 1, y, y - 1, y - 1}
	for i := 0; i < 8; i++ {
		m, n := dx[i], dy[i]
		if inBound(m, n) {

			// Atat timp cat culoarea e diferita verifica patratul
			if tabla[x][y].Culoare != tabla[m][n].Culoare {

				// Daca alegem sa mutam afiseaza patratele disponibile
				if mutare {
					// Daca regele este alb si patratul de coords m, n nu e controlat de negru, regele poate ataca acel patrat
					if tabla[x][y].Culoare == 'W' && tabla[m][n].Control < 2 {
						tabla[m][n].Atacat = true
					}
					// Daca regele este negru si patratul de coords m, n nu e controlat de alb, regele poate ataca acel patrat
					if tabla[x][y].Culoare == 'B' && (tabla[m][n].Control == 2 || tabla[m][n].Control == 0) {
						tabla[m][n].Atacat = true
					}
				} else {
					setareControl(&tabla[m][n], tabla[x][y].Culoare)
				}
			}
		}
	}
}

// miscarePion cauta miscarile posibile pentru pionul de la (x, y)
func (p *Piesa) miscarePion(tabla *[8][8]Piesa, x, y int, mutare, isSah bool) {
	var dy = []int{y - 1, y, y + 1}
	for i := 0; i < 3; i++ {

		// Daca piesa e alba, verifica patratele de sus
		if tabla[x][y].Culoare == 'W' {
			m, n := x-1, dy[i]
			if inBound(m, n) {
				// Daca se afla o piesa inamica in stanga- sau dreapta-susul pionului, verifica acel patrat
				if tabla[m][n].Tip != 0 && i != 1 {
					if tabla[x][y].Culoare != tabla[m][n].Culoare {
						if mutare {
							// Daca regele se afla in sah, verificam daca aceasta mutare il scaote din sah
							if isSah {
								verifIesireSah(tabla, x, y, m, n)
							} else {
								tabla[m][n].Atacat = true
							}
						} else {
							setareControl(&tabla[m][n], tabla[x][y].Culoare)
						}
					}
				}
				// Verifica daca poti muta pionul un patrat in fata
				if mutare {
					if tabla[m][n].Tip == 0 && i == 1 {
						// Daca regele se afla in sah, verificam daca aceasta mutare il scaote din sah
						if isSah {
							verifIesireSah(tabla, x, y, m, n)
						} else {
							tabla[m][n].Atacat = true
						}
					}
				}
			}
		}

		// Daca piesa e neagra, verifica patratele de jos
		if tabla[x][y].Culoare == 'B' {
			m, n := x+1, dy[i]
			if inBound(m, n) {
				// Daca se afla o piesa inamica in stanga- sau dreapta-josul pionului, verifica acel patrat
				if tabla[m][n].Tip != 0 && i != 1 {
					if tabla[x][y].Culoare != tabla[m][n].Culoare {
						if mutare {
							// Daca regele se afla in sah, verificam daca aceasta mutare il scaote din sah
							if isSah {
								verifIesireSah(tabla, x, y, m, n)
							} else {
								tabla[m][n].Atacat = true
							}
						} else {
							setareControl(&tabla[m][n], tabla[x][y].Culoare)
						}
					}
				}
				// Verifica daca poti muta pionul un patrat in fata
				if mutare {
					if tabla[m][n].Tip == 0 && i == 1 {
						// Daca regele se afla in sah, verificam daca aceasta mutare il scaote din sah
						if isSah {
							verifIesireSah(tabla, x, y, m, n)
						} else {
							tabla[m][n].Atacat = true
						}
					}
				}
			}
		}
	}

	// Verifica daca piesa poate parcurge 2 patrate
	if mutare {
		// Daca a fost mutat deja, nu mai poate parcurge 2 patrate
		if tabla[x][y].Mutat == false {
			if tabla[x][y].Culoare == 'W' {
				// Verifica daca urmatoarele doua patrate sunt libere
				if tabla[x-1][y].Tip == 0 && tabla[x-2][y].Tip == 0 {
					// Daca regele se afla in sah, verificam daca aceasta mutare il scaote din sah
					if isSah {
						verifIesireSah(tabla, x, y, x-2, y)
					} else {
						tabla[x-2][y].Atacat = true
					}
				}
			}
			if tabla[x][y].Culoare == 'B' {
				// Verifica daca urmatoarele doua patrate sunt libere
				if tabla[x+1][y].Tip == 0 && tabla[x+2][y].Tip == 0 {
					// Daca regele se afla in sah, verificam daca aceasta mutare il scaote din sah
					if isSah {
						verifIesireSah(tabla, x, y, x+2, y)
					} else {
						tabla[x+2][y].Atacat = true
					}
				}
			}
		}
	}
}

// miscareNebun cauta miscarile posibile pentru nebunul de la (x, y)
func (p *Piesa) miscareNebun(tabla *[8][8]Piesa, x, y int, mutare, isSah bool) {
	var dx = []int{-1, -1, 1, 1}
	var dy = []int{-1, 1, -1, 1}
	for i := 0; i < 4; i++ {
		for j := 1; j < 8; j++ {

			m, n := x+j*dx[i], y+j*dy[i]

			if inBound(m, n) {
				// Daca regele se afla in sah, verificam daca aceasta mutare il scaote din sah
				if isSah {
					verifIesireSah(tabla, x, y, m, n)
				}

				// Daca vede o piesa de aceeasi culoare, se termina cautarea pe acea diagonala
				if tabla[x][y].Culoare == tabla[m][n].Culoare {

					if !mutare {
						setareControl(&tabla[m][n], tabla[x][y].Culoare)
					}
					break
				} else {
					if mutare {
						// Daca regele se afla in sah, verificam daca aceasta mutare il scaote din sah
						if isSah {
							verifIesireSah(tabla, x, y, m, n)
						} else {
							tabla[m][n].Atacat = true
						}
					} else {
						setareControl(&tabla[m][n], tabla[x][y].Culoare)
					}

					// Daca vede o piesa de alta culoare, afiseaza ca poate ataca acel patrat, dupa care opreste cautarea pe acea diagonala
					if tabla[x][y].Culoare != tabla[m][n].Culoare && tabla[m][n].Culoare != 0 {
						break
					}
				}
			}
		}
	}
}

// miscareCal cauta miscarile posibile pentru calul de la (x, y)
func (p *Piesa) miscareCal(tabla *[8][8]Piesa, x, y int, mutare, isSah bool) {
	var dx = []int{-2, -1, -1, -2, 2, 1, 1, 2}
	var dy = []int{-1, -2, 2, 1, -1, -2, 2, 1}
	for i := 0; i < 8; i++ {

		m, n := x+dx[i], y+dy[i]

		if inBound(m, n) {
			// Daca patratul m, n e de culoare diferita sau liber, ataca acel patrat
			if tabla[x][y].Culoare != tabla[m][n].Culoare {
				if mutare {
					// Daca regele se afla in sah, verificam daca aceasta mutare il scaote din sah
					if isSah {
						verifIesireSah(tabla, x, y, m, n)
					} else {
						tabla[m][n].Atacat = true
					}
				} else {
					setareControl(&tabla[m][n], tabla[x][y].Culoare)
				}
			} else {
				setareControl(&tabla[m][n], tabla[x][y].Culoare)
			}
		}
	}
}

// miscareTura cauta miscarile posibile pentru tura de la (x, y)
func (p *Piesa) miscareTura(tabla *[8][8]Piesa, x, y int, mutare, isSah bool) {
	var dx = []int{-1, 0, 1, 0}
	var dy = []int{0, 1, 0, -1}

	for i := 0; i < 4; i++ {
		for j := 1; j < 8; j++ {
			m, n := x+j*dx[i], y+j*dy[i]

			if inBound(m, n) {
				// Daca vede o piesa de aceeasi culoare, se termina cautarea pe acea linie/coloana
				if tabla[x][y].Culoare == tabla[m][n].Culoare {
					if !mutare {
						setareControl(&tabla[m][n], tabla[x][y].Culoare)
					}
					break
				} else {
					if mutare {
						// Daca regele se afla in sah, verificam daca aceasta mutare il scaote din sah
						if isSah {
							verifIesireSah(tabla, x, y, m, n)
						} else {
							tabla[m][n].Atacat = true
						}
					} else {
						setareControl(&tabla[m][n], tabla[x][y].Culoare)
					}

					// Daca vede o piesa de alta culoare, afiseaza ca poate ataca acel patrat, dupa care opreste cautarea pe acea linie/coloana
					if tabla[x][y].Culoare != tabla[m][n].Culoare && tabla[m][n].Culoare != 0 {
						break
					}
				}
			}
		}
	}
}
// miscareRegina cauta miscarile posibile pentru regina de la (x, y)
func (p *Piesa) miscareRegina(tabla *[8][8]Piesa, x, y int, mutare, isSah bool) {
	// regina se misca ca nebunul si tura
	p.miscareNebun(tabla, x, y, mutare, isSah)
	p.miscareTura(tabla, x, y, mutare, isSah)
}
