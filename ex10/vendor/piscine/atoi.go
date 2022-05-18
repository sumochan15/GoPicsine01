package piscine

func Atoi(s string) int {
	sign := 1
	sum := 0

	for i, c := range s{
		if i == 0 && (c == '-' || c == '+'){
			if c =='-'{
				sign = -1
			}
		} else if c < '0' || '9' < c {
				return 0
		} else {
			sum = sum*10 + int(c) - '0'
		}
	}
	return sum * sign
}

//int minの時にオーバーフローの可能性あり。
//符号を逆にするかマイナスはマイナスのまま処理する必要がある