package piese

// miscareRege cauta miscarile posibile pt regele ales
func (p *Piesa) miscareRege(tabla *[8][8]Piesa, x, y int, mutare, isSah bool) {
	var dx = []int{x - 1, x - 1, x - 1, x, x + 1, x + 1, x + 1, x}
	var dy = []int{y - 1, y, y + 1, y + 1, y + 1, y, y - 1, y - 1}

	// Verific daca poate face rocada
	if mutare == true && !tabla[x][y].Mutat {
		if verifRocada(x, y, -4) {
			tabla[x][y-2].Atacat = true
		}
		if verifRocada(x, y, +3) {
			tabla[x][y+2].Atacat = true
		}
	}
	for i := 0; i < 8; i++ {
		m, n := dx[i], dy[i]
		if verifInBound(m, n) {

			// Atat timp cat culoarea e diferita verifica patratul
			if tabla[x][y].Culoare != tabla[m][n].Culoare {

				// Daca alegem sa mutam afiseaza patratele disponibile
				if mutare {
					// Daca regele este alb si patratul de coords m, n nu e controlat de negru, regele poate ataca acel patrat
					if tabla[x][y].Culoare == 'W' && (tabla[m][n].Control == 1 || tabla[m][n].Control == 0) {
						tabla[m][n].Atacat = true
					}
					// Daca regele este negru si patratul de coords m, n nu e controlat de alb, regele poate ataca acel patrat
					if tabla[x][y].Culoare == 'B' && (tabla[m][n].Control == 2 || tabla[m][n].Control == 0) {
						tabla[m][n].Atacat = true
					}
				} else {
					if isSah {
						if !verifSah(tabla, x, y, m, n) {
							Mat = false
							existaMutare = true
						}
					}
				}
			}
			if !mutare && !isSah {
				setareControl(&tabla[m][n], tabla[x][y].Culoare)
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
			if verifInBound(m, n) {
				// Daca se afla o piesa inamica in stanga- sau dreapta-susul pionului, verifica acel patrat
				if (tabla[m][n].Tip != 0 || (tabla[m+1][n].Tip != 0 && tabla[m+1][n].EnPassant)) && i != 1 {
					if tabla[x][y].Culoare != tabla[m][n].Culoare {
						if mutare {
							// Daca regele se afla in sah, verificam daca aceasta mutare il scaote din sah
							if isSah {
								verifIesireSah(tabla, x, y, m, n)
							} else {
								if !verifSah(tabla, x, y, m, n) {
									tabla[m][n].Atacat = true
								}
							}
						} else {
							if isSah {
								if !verifSah(tabla, x, y, m, n) {
									Mat = false
									existaMutare = true
								}
							}
						}
					}
				}
				if !mutare && !isSah && i != 1 {
					setareControl(&tabla[m][n], tabla[x][y].Culoare)
				}
				// Verifica daca poti muta pionul un patrat in fata
				if tabla[m][n].Tip == 0 && i == 1 {
					if mutare {
						// Daca regele se afla in sah, verificam daca aceasta mutare il scaote din sah
						if isSah {
							verifIesireSah(tabla, x, y, m, n)
						} else {
							if !verifSah(tabla, x, y, m, n) {
								tabla[m][n].Atacat = true
							}
						}
					} else {
						if isSah {
							if !verifSah(tabla, x, y, m, n) {
								Mat = false
								existaMutare = true
							}
						}
					}
				}
			}
		}

		// Daca piesa e neagra, verifica patratele de jos
		if tabla[x][y].Culoare == 'B' {
			m, n := x+1, dy[i]
			if verifInBound(m, n) {
				// Daca se afla o piesa inamica in stanga- sau dreapta-josul pionului, verifica acel patrat
				if (tabla[m][n].Tip != 0 || (tabla[m-1][n].Tip != 0 && tabla[m-1][n].EnPassant)) && i != 1 {
					if tabla[x][y].Culoare != tabla[m][n].Culoare {
						if mutare {
							// Daca regele se afla in sah, verificam daca aceasta mutare il scaote din sah
							if isSah {
								verifIesireSah(tabla, x, y, m, n)
							} else {
								if !verifSah(tabla, x, y, m, n) {
									tabla[m][n].Atacat = true
								}
							}
						} else {
							if isSah {
								if !verifSah(tabla, x, y, m, n) {
									Mat = false
									existaMutare = true
								}
							}
						}
					}
				}
				if !mutare && !isSah && i != 1 {
					setareControl(&tabla[m][n], tabla[x][y].Culoare)
				}
				// Verifica daca poti muta pionul un patrat in fata
				if tabla[m][n].Tip == 0 && i == 1 {
					if mutare {
						// Daca regele se afla in sah, verificam daca aceasta mutare il scaote din sah
						if isSah {
							verifIesireSah(tabla, x, y, m, n)
						} else {
							if !verifSah(tabla, x, y, m, n) {
								tabla[m][n].Atacat = true
							}
						}
					} else {
						if isSah {
							if !verifSah(tabla, x, y, m, n) {
								Mat = false
								existaMutare = true
							}
						}
					}
				}
			}
		}
	}

	// Verifica daca piesa poate parcurge 2 patrate
	// Daca a fost mutat deja, nu mai poate parcurge 2 patrate
	if tabla[x][y].Mutat == false {
		if tabla[x][y].Culoare == 'W' {
			if verifInBound(x-2, y) {
				// Verifica daca urmatoarele doua patrate sunt libere
				if tabla[x-1][y].Tip == 0 && tabla[x-2][y].Tip == 0 {
					// Daca regele se afla in sah, verificam daca aceasta mutare il scaote din sah
					if mutare {
						if isSah {
							verifIesireSah(tabla, x, y, x-2, y)
						} else {
							if !verifSah(tabla, x, y, x-2, y) {
								tabla[x-2][y].Atacat = true
							}
						}
					} else {
						if isSah {
							if !verifSah(tabla, x, y, x-2, y) {
								Mat = false
								existaMutare = true
							}
						}
					}
				}
			}
		}
		if tabla[x][y].Culoare == 'B' {
			if verifInBound(x+2, y) {
				// Verifica daca urmatoarele doua patrate sunt libere
				if tabla[x+1][y].Tip == 0 && tabla[x+2][y].Tip == 0 {
					// Daca regele se afla in sah, verificam daca aceasta mutare il scaote din sah
					if mutare {
						if isSah {
							verifIesireSah(tabla, x, y, x+2, y)
						} else {
							if !verifSah(tabla, x, y, x+2, y) {
								tabla[x+2][y].Atacat = true
							}
						}
					} else {
						if isSah {
							if !verifSah(tabla, x, y, x+2, y) {
								Mat = false
								existaMutare = true
							}
						}
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

			if verifInBound(m, n) {
				if mutare && !isSah {
					if j == 1 && verifSah(tabla, x, y, m, n) {
						break
					}
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
						if isSah {
							if !verifSah(tabla, x, y, m, n) {
								Mat = false
								existaMutare = true
							}
						} else {
							setareControl(&tabla[m][n], tabla[x][y].Culoare)
						}
					}

					// Daca vede o piesa de alta culoare, afiseaza ca poate ataca acel patrat, dupa care opreste cautarea pe acea diagonala
					if tabla[x][y].Culoare != tabla[m][n].Culoare && tabla[m][n].Culoare != 0 {
						if tabla[m][n].Tip == 'K' {
							m2, n2 := x+(j+1)*dx[i], y+(j+1)*dy[i]
							if verifInBound(m2, n2) {
								setareControl(&tabla[m2][n2], tabla[x][y].Culoare)
							}
						}
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

		if verifInBound(m, n) {
			// Daca patratul m, n e de culoare diferita sau liber, ataca acel patrat
			if tabla[x][y].Culoare != tabla[m][n].Culoare {
				if mutare {
					// Daca regele se afla in sah, verificam daca aceasta mutare il scaote din sah
					if isSah {
						verifIesireSah(tabla, x, y, m, n)
					} else {
						if !verifSah(tabla, x, y, m, n) {
							tabla[m][n].Atacat = true
						}
					}
				} else {
					if isSah {
						if !verifSah(tabla, x, y, m, n) {
							Mat = false
							existaMutare = true
						}
					} else {
						setareControl(&tabla[m][n], tabla[x][y].Culoare)
					}
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

			if verifInBound(m, n) {
				if mutare && !isSah {
					if j == 1 && verifSah(tabla, x, y, m, n) {
						break
					}
				}

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
						if isSah {
							if !verifSah(tabla, x, y, m, n) {
								Mat = false
								existaMutare = true
							}
						} else {
							setareControl(&tabla[m][n], tabla[x][y].Culoare)
						}
					}

					// Daca vede o piesa de alta culoare, afiseaza ca poate ataca acel patrat, dupa care opreste cautarea pe acea linie/coloana
					if tabla[x][y].Culoare != tabla[m][n].Culoare && tabla[m][n].Culoare != 0 {
						if tabla[m][n].Tip == 'K' {
							m2, n2 := x+(j+1)*dx[i], y+(j+1)*dy[i]
							if verifInBound(m2, n2) {
								setareControl(&tabla[m2][n2], tabla[x][y].Culoare)
							}
						}
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
