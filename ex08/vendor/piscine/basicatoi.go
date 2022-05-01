package piscine

func BasicAtoi(s string) int {
	sum := 0
	for _, c := range s {
		sum = sum*10 + int(c) - '0'
	}
	return sum
}
