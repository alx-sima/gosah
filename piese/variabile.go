package piese

const (
	// Width retine lungimea ecranului
	Width = 1920
	// Height retine inaltimea ecranului
	Height = 1080
	// L retine latura unui patrat
	L = Height / 8
	// Offset retine y-ul care trebuie adaugat pt padding
	Offset = (Width - Height) / 2
	// Status == PatVal inseamna ca meciul este in pat
	PatVal = 1
	// Status == MatVal inseamna ca meciul este in mat
	MatVal = 2
)

type pieseCounter struct {
	piese map[rune]int
	nr    int
}

func (p *pieseCounter) edit(tip rune, i int) {
	p.piese[tip] += i
	p.nr += i
	if p.piese[tip] < 0 {
		p.piese[tip] = 0
	}
	if p.nr < 0 {
		p.nr = 0
	}
}

var (
	// Board retine tabla de joc
	Board [8][8]Piesa
	// ramaseAlbe, ramaseNegre retin piesele ramase pt. fiecare jucator.
	ramaseAlbe, ramaseNegre pieseCounter

	// selected retine piesa pe care s-a dat click
	selected PozitiePiesa
	// Clicked retine daca fost dat click pe o piesa
	Clicked bool
	// existaMutare retine daca exista miscari legale
	existaMutare bool
	// SahAlb retine daca regele alb e in sah
	SahAlb bool
	// SahNegru retine daca regele negru e in sah
	SahNegru bool
	// DEPRECATED
	// Mat retine daca exista miscari care sa te scoata din sah sau nu
	Mat bool
	// DEPRECATED
	// Pat retine daca jocul se afla in stadiul de pat/egalitate (pentru a termina jocul)
	Pat bool
	// mutariUltimaCapturare retine numarul de mutari de la ultima capturare
	mutariUltimaCapturare int
	// Turn retine 'W' daca e randul albului, sau 'B' daca e randul negrului
	Turn rune
	// Editing retine true daca este in modul editare, false daca e joc normal
	Editing bool
	// RegeNegru retine pozitia regelui negru
	RegeNegru PozitiePiesa
	// RegeAlb retine pozitia regelui alb
	RegeAlb PozitiePiesa
	// Started retine daca jocul a inceput
	Started bool
	// Nivele retine numele nivelelor din fisiere (+random si editor)
	Nivele []string
	// Castigator retine cine a castigat
	Castigator string
	// stare retine daca meciul e in sah sau pat
	stare int
	// TimpRamas retine timpul ramas pentru fiecare jucator
	TimpRamas struct {
		Alb struct {
			Min int
			Sec int
		}
		Negru struct {
			Min int
			Sec int
		}
	}
	// Miscari retine toate miscarile efectuate de jucatori
	Miscari struct {
		Negru []string
		Alb   []string
	}
	// Miscari.Negru = append(Miscari.Negru, miscare)
	pieseAtacaPatrat struct {
		nr int
		lin []int
		col []int
	}

	// DEPRECATED
	// TipSelectat retine ce piesa e selectata in editor
	TipSelectat rune
)
