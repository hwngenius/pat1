package main

import (
	"fmt"
	"unicode"
)

func fun(s string) int {
	if len(s) < 6 {
		return 1
	}
	var zimu, shuzi bool
	for i := 0; i < len(s); i++ {
		if unicode.IsPunct(rune(s[i])) {
			if s[i] == '.' {
				continue
			} else {
				return 2
			}
		} else {
			if unicode.IsDigit(rune(s[i])) {
				shuzi = true
			} else if unicode.IsLetter(rune(s[i])) {
				zimu = true
			}
		}
	}
	if 
	if zimu == false {
		return 4
	}
	if shuzi == false {
		return 3
	}
	return 0
}

func main() {
	var N int = 1
	var s string = "123456"
	fmt.Scan(&N)
	for i := 0; i < N; i++ {
		fmt.Scan(&s)
		switch fun(s) {
		case 0:
			fmt.Println("Your password is wan mei.")
		case 1:
			fmt.Println("Your password is tai duan le.")
		case 2:
			fmt.Println("Your password is tai luan le.")
		case 3:
			fmt.Println("Your password needs shu zi.")
		case 4:
			fmt.Println("Your password needs zi mu.")
		}
	}
}
