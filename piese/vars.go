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

var (
	// Board retine tabla de joc
	Board [8][8]Piesa
	// PieseAlbe retine piesele albe ramase. 'B' reprezinta nebunul de pe patratele albe, iar 'b' reprezinta nebunul de pe patratele negre
	PieseAlbe []rune
	// PieseNegre retine piesele negre ramase. 'B' reprezinta nebunul de pe patratele albe, iar 'b' reprezinta nebunul de pe patratele negre
	PieseNegre []rune
	// Selected retine piesa pe care s-a dat click
	Selected PozitiePiesa
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
	// MutariUltimaCapturare retine numarul de mutari de la ultima capturare
	MutariUltimaCapturare int
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
	// Stare retine daca meciul e in sah sau pat
	Stare int
	//
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
)
