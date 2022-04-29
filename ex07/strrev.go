

func StrLen(s string) int {
	len := 0
	for range s {
		len++
	}
	return len
}

func StrRev(s string) string{
	len := Strlen(s)
	runes := []rune(s)
	for i, j := 0, len-1; i < j; i, j = i+1, j-1{
		runes[i], runes[j] = runes[j], runes[i]
	}
	return string(runes)
}