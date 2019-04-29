package main

import (
	"fmt"
	"strconv"
)

func sort(s string) (t string) {
	l := []byte(s)
	for i := len(l) - 1; i > 0; i-- {
		for j := 1; j <= i; j++ {
			if l[j-1] < l[j] {
				l[j-1], l[j] = l[j], l[j-1]
			}
		}
	}
	t = string(l)
	if len(s) == 1 {
		t += "000"
	} else if len(s) == 2 {
		t += "00"
	} else if len(s) == 3 {
		t += "0"
	}
	return
}

func reverse(s string) (t string) {
	l := []byte(s)
	for i, j := 0, len(s)-1; i < len(s); i, j = i+1, j-1 {
		l[i] = s[j]
	}
	t = string(l)
	return
}
func main() {
	var s string = "1"
	var mask bool
	fmt.Scanf("%s", &s)
	for {
		n, _ := strconv.Atoi(sort(s))
		m, _ := strconv.Atoi(reverse(sort(s)))
		if mask == false {
			mask = true
		} else {
			fmt.Printf("\n")
		}
		fmt.Printf("%04d - %04d = %04d", n, m, n-m)
		if n-m == 0 || n-m == 6174 {
			break
		} else {
			s = strconv.Itoa(n - m)
		}
	}
}
