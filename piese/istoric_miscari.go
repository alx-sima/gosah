package piese

import "fmt"

// transforma x din int in litera corespunzatoare coloanei
func transform(x int) string {
	return fmt.Sprintf("%c", 'a'+x)
}

// verifica daca exista mai multe piese de acelasi tip care pot ajunge in acelasi patrat
func verifPieseAtacaAcelasiPatrat(x, y int, piesa rune) {
	pieseAtacaPatrat.nr = 0
	pieseAtacaPatrat.lin = []int{}
	pieseAtacaPatrat.col = []int{}

	p := Board[x][y]

	switch piesa {
	case 'P':
		p.miscarePion(&Board, x, y, false, false, true)
	case 'B':
		p.miscareNebun(&Board, x, y, false, false, true, false)
	case 'N':
		p.miscareCal(&Board, x, y, false, false, true)
	case 'R':
		p.miscareTura(&Board, x, y, false, false, true, false)
	case 'Q':
		p.miscareRegina(&Board, x, y, false, false, true)
	default:
		return
	}
}

func numire(capturare bool, x, y, m, n int, piesa, promotion rune) (rezultat string) {
	// In sah randurile incep de la 1, y in cod incepe de la 0 => incrementam cu 1
	m++
	x++

	// mutarile pionilor nu au prefix cu piesa mutata
	if piesa != 'P' {
		rezultat += string(piesa)
	} else if capturare {
		rezultat += transform(y)
	}

	// daca mai multe piese de acelasi tip pot ajunge in patratul selectat, adaugam linia, coloana sau ambele a piesei mutate
	verifPieseAtacaAcelasiPatrat(m - 1, n, piesa)
	if pieseAtacaPatrat.nr > 1 {
		var linSame, colSame bool
		for i := 0; i < pieseAtacaPatrat.nr - 1; i++ {
			for j := i + 1; j < pieseAtacaPatrat.nr; j++ {
				if pieseAtacaPatrat.lin[i] == pieseAtacaPatrat.lin[j] {
					linSame = true
				}
				if pieseAtacaPatrat.col[i] == pieseAtacaPatrat.col[j] {
					colSame = true
				}
			}
		}
		if colSame || !(colSame || linSame) {
			rezultat += transform(y)
		}
		if linSame {
			rezultat += fmt.Sprintf("%d", x)
		}
	}

	// capturarea este reprezentata prin x
	if capturare {
		rezultat += "x"
	}

	// adaugam pozitia in care ajunge piesa
	rezultat += transform(n)
	rezultat += fmt.Sprintf("%d", m)

	// daca piesa este pion si ajunge pe ultima coloana adauga in ce a promovat
	if promotion != 0 {
		rezultat += "="
		rezultat += string(promotion)
	}
	return
}
