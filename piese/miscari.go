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

func (p *Piesa) miscarePion(tabla *[8][8]Piesa, x, y int) {}

func (p *Piesa) miscareNebun(tabla *[8][8]Piesa, x, y int) {}

func (p *Piesa) miscareCal(tabla *[8][8]Piesa, x, y int) {}

func (p *Piesa) miscareTura(tabla *[8][8]Piesa, x, y int) {}

func (p *Piesa) miscareRegina(tabla *[8][8]Piesa, x, y int) {}
