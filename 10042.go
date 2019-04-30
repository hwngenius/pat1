package main

import (
	"fmt"
	"unicode"
)

func main() {
	m := make(map[byte]int)
	var t, b byte
	var max int
	for i, _ := fmt.Scanf("%c", &b); i == 1; i, _ = fmt.Scanf("%c", &b) {
		if unicode.IsLetter(rune(b)) {
			b = byte(unicode.ToLower(rune(b)))
			m[b]++
			if m[b] > max {
				max = m[b]
				t = b
			} else if m[b] == max {
				if t > b {
					t = b
				}
			}
		}
	}
	fmt.Printf("%c %d", t, max)
}
