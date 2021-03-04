package piese

import "fmt"

func transform(x int) string {
	return fmt.Sprintf("%c", 'a'+x)
}

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

	if piesa != 'P' {
		rezultat += string(piesa)
	}

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
		if linSame || !(colSame || linSame) {
			rezultat += transform(y)
		}
		if colSame {
			rezultat += fmt.Sprintf("%d", x)
		}
	}

	if capturare {
		rezultat += "x"
	}

	rezultat += transform(n)
	rezultat += fmt.Sprintf("%d", m)

	if promotion != 0 {
		rezultat += "="
		rezultat += string(promotion)
	}
	return
}
