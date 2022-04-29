package main

import (
	"fmt"
	"piscine"
)

func main() {
	a := 13
	b := 2
	if b == 0{
		fmt.Println(0)
	}else{
	piscine.UltimateDivMod(&a, &b)
	fmt.Println(a)
	fmt.Println(b)
	}
}