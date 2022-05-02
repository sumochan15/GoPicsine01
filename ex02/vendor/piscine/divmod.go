package piscine

func DivMod(a int, b int, div *int, mod *int) {
	if b != 0 {
		*div, *mod = a/b, a%b
	}
}