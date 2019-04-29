package main

import (
	"fmt"
	"unicode"
)

func main() {
	var set [300]int
	var s1, s2 string
	fmt.Scan(&s1)
	for i := 0; i < len(s1); i++ {
		t := unicode.ToUpper(rune(s1[i]))
		set[int(t)] = 1
	}
	fmt.Scan(&s2)
	for i := 0; i < len(s2); i++ {
		t := unicode.ToUpper(rune(s2[i]))
		set[int(t)] = 0
	}
	for i := 0; i < len(s1); i++ {
		t := unicode.ToUpper(rune(s1[i]))
		if set[int(t)] == 1 {
			set[int(t)] = 0
			fmt.Printf("%c", t)
		}
	}
}
