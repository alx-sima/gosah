package piese

const (
	// Width retine lungimea ecranului
	Width = 1920
	// Height retine inaltimea ecranului
	Height = 1080
	// L retine latura unui patrat
	L = Height / 8
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
	// Changed retine daca trebuie (re)desenat ecranul
	Changed bool
	// existaMutare retine daca exista miscari legale
	existaMutare bool
	// SahAlb retine daca regele alb e in sah
	SahAlb bool
	// SahNegru retine daca regele negru e in sah
	SahNegru bool
	// Mat retine daca exista miscari care sa te scoata din sah sau nu
	Mat bool
	// Pat retine daca jocul se afla in stadiul de pat/egalitate (pentru a termina jocul)
	Pat bool
	// MutariUltimaCapturare retine numarul de mutari de la ultima capturare
	MutariUltimaCapturare int
	// Turn retine 'W' daca e randul albului, sau 'B' daca e randul negrului
	Turn rune
	// Editing retine true daca este in modul editare, false daca e joc normal
	Editing bool
)
