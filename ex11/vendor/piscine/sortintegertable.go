package piscine

func SortIntegerTable(table []int) {
	len := 0

	for range table {
		len++
	}
	for i := 0 ; i < len-1; i++ {
		for j := i+1 ; j < len; j++{
			if table[i] > table[j]{
				table[i], table[j] = table[j], table[i]
			}
		}
	}
}