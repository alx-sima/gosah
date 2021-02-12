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
func (p *Piesa) miscareRege(tabla *[8][8]Piesa, x, y int) {
	var dx = []int{x - 1, x - 1, x - 1, x, x + 1, x + 1, x + 1, x}
	var dy = []int{y - 1, y, y + 1, y + 1, y + 1, y, y - 1, y - 1}
	for i := 0; i < 8; i++ {
		if inBound(dx[i], dy[i]) {
			tabla[dx[i]][dy[i]].Atacat = true
		}
	}
	// TODO: evitare sah
	// TODO: evitare rege inamic
	// TODO: evitare piese din aceeasi echipa
}

func (p *Piesa) miscarePion(tabla *[8][8]Piesa, x, y int) {
	var dy = []int{y - 1, y, y + 1}
	for i := 0; i < 3; i++ {

		// Daca piesa e alba, verifica patratele de sus
		if tabla[x][y].Culoare == 'W' {
			if inBound(x-1, dy[i]) {
				if (tabla[x-1][dy[i]].Tip != 0 && i != 1) || (tabla[x-1][dy[i]].Tip == 0 && i == 1) {
					tabla[x-1][dy[i]].Atacat = true
				}
			}
		}

		// Daca piesa e neagra, verifica patratele de jos
		if tabla[x][y].Culoare == 'B' {
			if inBound(x+1, dy[i]) {
				if (tabla[x+1][dy[i]].Tip != 0 && i != 1) || (tabla[x+1][dy[i]].Tip == 0 && i == 1) {
					tabla[x+1][dy[i]].Atacat = true
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

func (p *Piesa) miscareNebun(tabla *[8][8]Piesa, x, y int) {
	var dx = []int{-1, -1, 1, 1}
	var dy = []int{-1, 1, -1, 1}
	for i := 1; i < 8; i++ {
		for j := 0; j < 4; j++ {
			m, n := x+i*dx[j], y+i*dy[j]
			if inBound(m, n) {
				if tabla[x][y].Culoare == tabla[m][n].Culoare {
					dx[j], dy[j] = 0, 0
				} else{
					tabla[m][n].Atacat = true
				}
			}
		}
	}
}

func (p *Piesa) miscareCal(tabla *[8][8]Piesa, x, y int) {
	var dx = []int{-2, -1, -1, -2, 2, 1, 1, 2}
	var dy = []int{-1, -2, 2, 1, -1, -2, 2, 1}
	for i := 0; i < 8; i++ {
		if inBound(x+dx[i], y+dy[i]) {
			tabla[x+dx[i]][y+dy[i]].Atacat = true
		}
	}
}

func (p *Piesa) miscareTura(tabla *[8][8]Piesa, x, y int) {
	var dx = []int{-1, 0, 1, 0}
	var dy = []int{0, 1, 0, -1}
	for i := 1; i < 8; i++ {
		for j := 0; j < 4; j++ {
			m, n := x+i*dx[j], y+i*dy[j]
			if inBound(m, n) {
				// Daca gaseste piesa friendly, nu mai merge pe directia aceea
				if tabla[x][y].Culoare == tabla[m][n].Culoare {
					dx[j], dy[j] = 0, 0
				} else {
					tabla[m][n].Atacat = true
				}
			}
		}
	}
}

func (p *Piesa) miscareRegina(tabla *[8][8]Piesa, x, y int) {
	p.miscareNebun(tabla, x, y)
	p.miscareTura(tabla, x, y)
}
