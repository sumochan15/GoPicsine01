package piscine

func StrLen(s string)int{
	len := 0
	for range s{
		len++
	}
	return len
}

func StrRev(s string) string {
	
	rune_str := []rune(s)
	rev_str := ""
	// var rev_str string

	for i := StrLen(s) - 1; i >= 0; i--{
		rev_str += string(rune_str[i])
	}
	return rev_str
}
