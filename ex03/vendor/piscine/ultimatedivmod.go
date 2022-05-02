package piscine

func UltimateDivMod(a *int, b *int) {
	if *b != 0 {
	*a, *b= *a / *b, *a % *b 
	}
}