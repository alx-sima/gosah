package piese

func inBound(a, b int) bool {
	return a >= 0 && b >= 0 && a < 8 && b < 8
}
func Clear(tabla *[8][8]Piesa) {
	for i := 0; i < 8; i++ {
		for j := 0; j < 8; j++ {
			tabla[i][j].Mutabil = false
		}
	}
}
func (p *Piesa) miscareRege(tabla *[8][8]Piesa, x, y int) {
	var dx = []int{x - 1, x - 1, x - 1, x, x + 1, x + 1, x + 1, x}
	var dy = []int{y - 1, y, y + 1, y + 1, y + 1, y, y - 1, y - 1}
	for i := 0; i < 8; i++ {
		if inBound(dx[i], dy[i]) {
			tabla[dx[i]][dy[i]].Mutabil = true
		}
	}
	// TODO: evitare sah
	// TODO: evitare rege inamic
	// TODO: evitare piese din aceeasi echipa
}

func (p *Piesa) miscarePion(tabla *[8][8]Piesa, x, y int) {
	//var dx = []int{x - 1, x, x + 1}
	//var dy = []int{y - 2, y - 1, y + 1, y + 2}
}

func (p *Piesa) miscareNebun(tabla *[8][8]Piesa, x, y int) {
	var dx = []int{-1, -1, 1, 1}
	var dy = []int{-1, 1, -1, 1}
	for i := 1; i < 8; i++ {
		for j := 0; j < 4; j++ {
			if inBound(x+i*dx[j], y+i*dy[j]) {
				tabla[x+i*dx[j]][y+i*dy[j]].Mutabil = true
			}
		}
	}
}

func (p *Piesa) miscareCal(tabla *[8][8]Piesa, x, y int) {
	var dx = []int{-2, -1, -1, -2, 2, 1, 1, 2}
	var dy = []int{-1, -2, 2, 1, -1, -2, 2, 1}
	for i := 0; i < 8; i++ {
		if inBound(x+dx[i], y+dy[i]) {
			tabla[x+dx[i]][y+dy[i]].Mutabil = true
		}
	}
}

func (p *Piesa) miscareTura(tabla *[8][8]Piesa, x, y int) {
	var dx = []int{-1, 0, 1, 0}
	var dy = []int{0, 1, 0, -1}
	for i := 1; i < 8; i++ {
		for j := 0; j < 4; j++ {
			if inBound(x+i*dx[j], y+i*dy[j]) {
				tabla[x+i*dx[j]][y+i*dy[j]].Mutabil = true
			}
		}
	}
}

func (p *Piesa) miscareRegina(tabla *[8][8]Piesa, x, y int) {
	p.miscareNebun(tabla, x, y)
	p.miscareTura(tabla, x, y)
}
