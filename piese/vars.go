package piese

const (
	Width  = 1920       // Width retine lungimea ecranului
	Height = 1080       // Height retine inaltimea ecranului
	L      = Height / 8 // L retine latura unui patrat
)

var (
	Board      [8][8]Piesa  // Board retine tabla de joc
	PieseAlbe  []rune       // PieseAlbe retine piesele albe ramase. 'B' reprezinta nebunul de pe patratele albe, iar 'b' reprezinta nebunul de pe patratele negre
	PieseNegre []rune       // PieseNegre retine piesele negre ramase. 'B' reprezinta nebunul de pe patratele albe, iar 'b' reprezinta nebunul de pe patratele negre
	Selected   PozitiePiesa // Selected retine piesa pe care s-a dat click
	Clicked    bool         // Clicked retine daca fost dat click pe o piesa
	Changed    bool         // Changed retine daca trebuie (re)desenat ecranul
	SahAlb     bool         // SahAlb retine daca regele alb e in sah
	SahNegru   bool         // SahNegru retine daca regele negru e in sah
	Mat        bool         // Mat retine daca exista miscari care sa te scoata din sah sau nu
	Pat        bool         // Pat retine daca jocul se afla in stadiul de pat/egalitate (pentru a termina jocul)
	Turn       rune         // Turn retine 'W' daca e randul albului, sau 'B' daca e randul negrului
)
