package main

import (
	"fmt"
	"strconv"
)

func isRight(s string) bool {
	if s[0] == '0' {
		return false
	}
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		if s[i] != s[j] {
			return false
		}
	}
	return true
}

func reverse(s string) string {
	var r []byte
	for i := len(s) - 1; i >= 0; i-- {
		r = append(r, s[i])
	}
	return string(r)
}

func main() {
	var s string
	fmt.Scan(&s)
	if isRight(s) {
		fmt.Println(s, "is a palindromic number.")
		return
	}
	if s == "0" {
		fmt.Println(s, "is a palindromic number.")
		return
	}
	for i := 0; i < 10; i++ {
		n1, _ := strconv.Atoi(s)
		n2, _ := strconv.Atoi(reverse(s))
		n := n1 + n2
		fmt.Printf("%s + %s = %d\n", s, reverse(s), n)
		s = strconv.Itoa(n)
		if isRight(s) {
			fmt.Printf("%s is a palindromic number.\n", s)
			return
		}
	}
	fmt.Println("Not found in 10 iterations.")
}
