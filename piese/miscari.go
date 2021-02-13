package piese

func inBound(a, b int) bool {
	return a >= 0 && b >= 0 && a < 8 && b < 8
}
func Clear(tabla *[8][8]Piesa) {
	for i := 0; i < 8; i++ {
		for j := 0; j < 8; j++ {
			tabla[i][j].Atacat = false
		}
	}
}
// TODO: verificare pozitii controlate kinda buggy needs fix but works kinda
// Verifica pentru fiecare piesa ce patrate ataca si formeaza in tabla.Control o matrice care arata ce culoare controleaza fiecare patrat
func VerifPatrateAtacate(tabla *[8][8]Piesa) {
	for i := 0; i < 8; i++ {
		for j := 0; j < 8; j++ {
			if tabla[i][j].Tip != 0 {
				switch tabla[i][j].Tip {
				case 'K':
					tabla[i][j].miscareRege(tabla, i, j, 'V')
				case 'P':
					tabla[i][j].miscarePion(tabla, i, j, 'V')
				case 'B':
					tabla[i][j].miscareNebun(tabla, i, j, 'V')
				case 'N':
					tabla[i][j].miscareCal(tabla, i, j, 'V')
				case 'R':
					tabla[i][j].miscareTura(tabla, i, j, 'V')
				case 'Q':
					tabla[i][j].miscareRegina(tabla, i, j, 'V')
				default:
					return
				}
			}
		}
	}
}

func (p *Piesa) miscareRege(tabla *[8][8]Piesa, x, y int, utilizare rune) {
	var dx = []int{x - 1, x - 1, x - 1, x, x + 1, x + 1, x + 1, x}
	var dy = []int{y - 1, y, y + 1, y + 1, y + 1, y, y - 1, y - 1}
	for i := 0; i < 8; i++ {
		m, n := dx[i], dy[i]
		if inBound(m, n) {
			if tabla[x][y].Culoare != tabla[m][n].Culoare {
				if utilizare == 'M' {
					if tabla[x][y].Culoare == 'W' && tabla[m][n].Control < 2 {
						tabla[m][n].Atacat = true
					}
					if tabla[x][y].Culoare == 'B' && (tabla[m][n].Control == 2 || tabla[m][n].Control == 0) {
						tabla[m][n].Atacat = true
					}
				} else {
					if tabla[x][y].Culoare == 'W' {
						if tabla[m][n].Control == 2 {
							tabla[m][n].Control = 3
						} else {
							tabla[m][n].Control = 1
						}
					} else {
						if tabla[m][n].Control == 1 {
							tabla[m][n].Control = 3
						} else {
							tabla[m][n].Control = 2
						}
					}
				}
			}
		}
	}
}

func (p *Piesa) miscarePion(tabla *[8][8]Piesa, x, y int, utilizare rune) {
	var dy = []int{y - 1, y, y + 1}
	for i := 0; i < 3; i++ {

		// Daca piesa e alba, verifica patratele de sus
		if tabla[x][y].Culoare == 'W' {
			m, n := x-1, dy[i]
			if inBound(m, n) {
				if tabla[m][n].Tip != 0 && i != 1 {
					if tabla[x][y].Culoare != tabla[m][n].Culoare {
						if utilizare == 'M' {
							tabla[m][n].Atacat = true
						} else {
							if tabla[m][n].Control == 2 {
								tabla[m][n].Control = 3
							} else {
								tabla[m][n].Control = 1
							}
						}
					}
				}
				if tabla[m][n].Tip == 0 && i == 1 {
					tabla[m][n].Atacat = true
				}
			}
		}

		// Daca piesa e neagra, verifica patratele de jos
		if tabla[x][y].Culoare == 'B' {
			m, n := x+1, dy[i]
			if inBound(m, n) {
				if tabla[m][n].Tip != 0 && i != 1 {
					if tabla[x][y].Culoare != tabla[m][n].Culoare {
						if utilizare == 'M' {
							tabla[m][n].Atacat = true
						} else {
							if tabla[m][n].Control == 1 {
								tabla[m][n].Control = 3
							} else {
								tabla[m][n].Control = 2
							}
						}
					}
					if tabla[m][n].Tip == 0 && i == 1 {
						tabla[m][n].Atacat = true
					}
				}
			}
		}
	}

	// Verifica daca piesa poate parcurge 2 patrate
	if tabla[x][y].Mutat == false {
		if tabla[x][y].Culoare == 'W' {
			if tabla[x-1][y].Tip == 0 && tabla[x-2][y].Tip == 0 {
				tabla[x-2][y].Atacat = true
			}
		}
		if tabla[x][y].Culoare == 'B' {
			if tabla[x+1][y].Tip == 0 && tabla[x+2][y].Tip == 0 {
				tabla[x+2][y].Atacat = true
			}
		}
	}
}

func (p *Piesa) miscareNebun(tabla *[8][8]Piesa, x, y int, utilizare rune) {
	var dx = []int{-1, -1, 1, 1}
	var dy = []int{-1, 1, -1, 1}
	for i := 0; i < 4; i++ {
		for j := 1; j < 8; j++ {

			m, n := x+j*dx[i], y+j*dy[i]

			if inBound(m, n) {

				// Daca vede o piesa de aceeasi culoare, se termina cautarea pe acea diagonala
				if tabla[x][y].Culoare == tabla[m][n].Culoare {
					if utilizare == 'V' {
						if tabla[x][y].Culoare == 'W' {
							if tabla[m][n].Control == 2 {
								tabla[m][n].Control = 3
							} else {
								tabla[m][n].Control = 1
							}
						} else {
							if tabla[m][n].Control == 1 {
								tabla[m][n].Control = 3
							} else {
								tabla[m][n].Control = 2
							}
						}
					}
					break
				} else {
					if utilizare == 'M' {
						tabla[m][n].Atacat = true
					} else {
						if tabla[x][y].Culoare == 'W' {
							if tabla[m][n].Control == 2 {
								tabla[m][n].Control = 3
							} else {
								tabla[m][n].Control = 1
							}
						} else {
							if tabla[m][n].Control == 1 {
								tabla[m][n].Control = 3
							} else {
								tabla[m][n].Control = 2
							}
						}
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

func (p *Piesa) miscareCal(tabla *[8][8]Piesa, x, y int, utilizare rune) {
	var dx = []int{-2, -1, -1, -2, 2, 1, 1, 2}
	var dy = []int{-1, -2, 2, 1, -1, -2, 2, 1}
	for i := 0; i < 8; i++ {

		m, n := x+dx[i], y+dy[i]

		if inBound(m, n) {
			if tabla[x][y].Culoare != tabla[m][n].Culoare {
				if utilizare == 'M' {
					tabla[m][n].Atacat = true
				} else {
					if tabla[x][y].Culoare == 'W' {
						if tabla[m][n].Control == 2 {
							tabla[m][n].Control = 3
						} else {
							tabla[m][n].Control = 1
						}
					} else {
						if tabla[m][n].Control == 1 {
							tabla[m][n].Control = 3
						} else {
							tabla[m][n].Control = 2
						}
					}
				}
			} else {
				if utilizare == 'V' {
					if tabla[x][y].Culoare == 'W' {
						if tabla[m][n].Control == 2 {
							tabla[m][n].Control = 3
						} else {
							tabla[m][n].Control = 1
						}
					} else {
						if tabla[m][n].Control == 1 {
							tabla[m][n].Control = 3
						} else {
							tabla[m][n].Control = 2
						}
					}
				}
			}
		}
	}
}

func (p *Piesa) miscareTura(tabla *[8][8]Piesa, x, y int, utilizare rune) {
	var dx = []int{-1, 0, 1, 0}
	var dy = []int{0, 1, 0, -1}
	for i := 0; i < 4; i++ {
		for j := 1; j < 8; j++ {

			m, n := x+j*dx[i], y+j*dy[i]

			if inBound(m, n) {

				// Daca vede o piesa de aceeasi culoare, se termina cautarea pe acea linie/coloana
				if tabla[x][y].Culoare == tabla[m][n].Culoare {
					if utilizare == 'V' {
						if tabla[x][y].Culoare == 'W' {
							if tabla[m][n].Control == 2 {
								tabla[m][n].Control = 3
							} else {
								tabla[m][n].Control = 1
							}
						} else {
							if tabla[m][n].Control == 1 {
								tabla[m][n].Control = 3
							} else {
								tabla[m][n].Control = 2
							}
						}
					}
					break
				} else {
					if utilizare == 'M' {
						tabla[m][n].Atacat = true
					} else {
						if tabla[x][y].Culoare == 'W' {
							if tabla[m][n].Control == 2 {
								tabla[m][n].Control = 3
							} else {
								tabla[m][n].Control = 1
							}
						} else {
							if tabla[m][n].Control == 1 {
								tabla[m][n].Control = 3
							} else {
								tabla[m][n].Control = 2
							}
						}
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

func (p *Piesa) miscareRegina(tabla *[8][8]Piesa, x, y int, utilizare rune) {
	p.miscareNebun(tabla, x, y, utilizare)
	p.miscareTura(tabla, x, y, utilizare)
}
