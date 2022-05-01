package piscine

func BasicAtoi2(s string) int {
	sum := 0

	for _, c := range s {
		if c < '0' || '9' < c {
			return 0
		}
		sum = sum*10 + int(c) - '0'
	}
	return sum
}
