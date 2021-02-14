package piese

var sahAlb, sahNegru bool

func inBound(a, b int) bool {
	return a >= 0 && b >= 0 && a < 8 && b < 8
}

func Clear(tabla *[8][8]Piesa) {
	for i := 0; i < 8; i++ {
		for j := 0; j < 8; j++ {
			tabla[i][j].Atacat = false
			tabla[i][j].Control = 0
		}
	}
	VerifPatrateAtacate(tabla)
}

// Seteaza controlul acelui patrat
func SetareControl(tabla *[8][8]Piesa, culoare rune, x, y int) {
	if culoare == 'W' {
		if tabla[x][y].Control == 2 {
			tabla[x][y].Control = 3
		} else if tabla[x][y].Control == 0 {
			tabla[x][y].Control = 1
		}
	} else {
		if tabla[x][y].Control == 1 {
			tabla[x][y].Control = 3
		} else if tabla[x][y].Control == 0 {
			tabla[x][y].Control = 2
		}
	}
}

// Verifica pentru fiecare piesa ce patrate ataca si formeaza in tabla.Control o matrice care arata ce culoare controleaza fiecare patrat
func VerifPatrateAtacate(tabla *[8][8]Piesa) {
	for i := 0; i < 8; i++ {
		for j := 0; j < 8; j++ {
			tabla[i][j].Move(tabla, i, j, false, false)
		}
	}
}

func verifIesireSah(tabla *[8][8]Piesa, x, y, m, n int) {
	tabla[m][n].Tip = tabla[x][y].Tip
	tabla[m][n].Culoare = tabla[x][y].Culoare
	tabla[x][y].Tip = 0
	tabla[x][y].Culoare = 0
	VerifPatrateAtacate(tabla)

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

func (p *Piesa) miscareRege(tabla *[8][8]Piesa, x, y int, mutare, isSah bool) {
	var dx = []int{x - 1, x - 1, x - 1, x, x + 1, x + 1, x + 1, x}
	var dy = []int{y - 1, y, y + 1, y + 1, y + 1, y, y - 1, y - 1}
	for i := 0; i < 8; i++ {
		m, n := dx[i], dy[i]
		if inBound(m, n) {
			if tabla[x][y].Culoare != tabla[m][n].Culoare {
				if mutare {
					if tabla[x][y].Culoare == 'W' && tabla[m][n].Control < 2 {
						tabla[m][n].Atacat = true
					}
					if tabla[x][y].Culoare == 'B' && (tabla[m][n].Control == 2 || tabla[m][n].Control == 0) {
						tabla[m][n].Atacat = true
					}
				} else {
					SetareControl(tabla, tabla[x][y].Culoare, m, n)
				}
			}
		}
	}
}

func (p *Piesa) miscarePion(tabla *[8][8]Piesa, x, y int, mutare, isSah bool) {
	var dy = []int{y - 1, y, y + 1}
	for i := 0; i < 3; i++ {

		// Daca piesa e alba, verifica patratele de sus
		if tabla[x][y].Culoare == 'W' {
			m, n := x-1, dy[i]
			if inBound(m, n) {
				if tabla[m][n].Tip != 0 && i != 1 {
					if tabla[x][y].Culoare != tabla[m][n].Culoare {
						if mutare {
							if isSah {
								verifIesireSah(tabla, x, y, m, n)
							} else {
								tabla[m][n].Atacat = true
							}
						} else {
							SetareControl(tabla, tabla[x][y].Culoare, m, n)
						}
					}
				}
				if mutare {
					if tabla[m][n].Tip == 0 && i == 1 {
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
				if tabla[m][n].Tip != 0 && i != 1 {
					if tabla[x][y].Culoare != tabla[m][n].Culoare {
						if mutare {
							if isSah {
								verifIesireSah(tabla, x, y, m, n)
							} else {
								tabla[m][n].Atacat = true
							}
						} else {
							SetareControl(tabla, tabla[x][y].Culoare, m, n)
						}
					}
				}
				if mutare {
					if tabla[m][n].Tip == 0 && i == 1 {
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

	if mutare {
		// Verifica daca piesa poate parcurge 2 patrate
		if tabla[x][y].Mutat == false {
			if tabla[x][y].Culoare == 'W' {
				if tabla[x-1][y].Tip == 0 && tabla[x-2][y].Tip == 0 {
					if isSah {
						verifIesireSah(tabla, x, y, x-2, y)
					} else {
						tabla[x-2][y].Atacat = true
					}
				}
			}
			if tabla[x][y].Culoare == 'B' {
				if tabla[x+1][y].Tip == 0 && tabla[x+2][y].Tip == 0 {
					if isSah {
						verifIesireSah(tabla, x, y, x-2, y)
					} else {
						tabla[x-2][y].Atacat = true
					}
				}
			}
		}
	}
}

func (p *Piesa) miscareNebun(tabla *[8][8]Piesa, x, y int, mutare, isSah bool) {
	var dx = []int{-1, -1, 1, 1}
	var dy = []int{-1, 1, -1, 1}
	for i := 0; i < 4; i++ {
		for j := 1; j < 8; j++ {

			m, n := x+j*dx[i], y+j*dy[i]

			if inBound(m, n) {

				if isSah {
				}

				// Daca vede o piesa de aceeasi culoare, se termina cautarea pe acea diagonala
				if tabla[x][y].Culoare == tabla[m][n].Culoare {

					if !mutare {
						SetareControl(tabla, tabla[x][y].Culoare, m, n)
					}

					break
				} else {
					if mutare {
						if isSah {
							verifIesireSah(tabla, x, y, m, n)
						} else {
							tabla[m][n].Atacat = true
						}
					} else {
						SetareControl(tabla, tabla[x][y].Culoare, m, n)
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

func (p *Piesa) miscareCal(tabla *[8][8]Piesa, x, y int, mutare, isSah bool) {
	var dx = []int{-2, -1, -1, -2, 2, 1, 1, 2}
	var dy = []int{-1, -2, 2, 1, -1, -2, 2, 1}
	for i := 0; i < 8; i++ {

		m, n := x+dx[i], y+dy[i]

		if inBound(m, n) {
			if tabla[x][y].Culoare != tabla[m][n].Culoare {
				if mutare {
					if isSah {
						verifIesireSah(tabla, x, y, m, n)
					} else {
						tabla[m][n].Atacat = true
					}
				} else {
					SetareControl(tabla, tabla[x][y].Culoare, m, n)
				}
			} else {
				SetareControl(tabla, tabla[x][y].Culoare, m, n)
			}
		}
	}
}

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
						SetareControl(tabla, tabla[x][y].Culoare, m, n)
					}

					break
				} else {
					if mutare {
						if isSah {
							verifIesireSah(tabla, x, y, m, n)
						} else {
							tabla[m][n].Atacat = true
						}
					} else {
						SetareControl(tabla, tabla[x][y].Culoare, m, n)
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

func (p *Piesa) miscareRegina(tabla *[8][8]Piesa, x, y int, mutare, isSah bool) {
	p.miscareNebun(tabla, x, y, mutare, isSah)
	p.miscareTura(tabla, x, y, mutare, isSah)
}
