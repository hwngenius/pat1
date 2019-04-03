package main

import "fmt"

func byte2Num(c byte) (r int) {
	switch c {
	case '0':
		r = 0
	case '1':
		r = 1
	case '2':
		r = 2
	case '3':
		r = 3
	case '4':
		r = 4
	case '5':
		r = 5
	case '6':
		r = 6
	case '7':
		r = 7
	case '8':
		r = 8
	case '9':
		r = 9
	}
	return
}

func main() {
	var s string
	fmt.Scan(&s)
	for i := len(s) - 1; i >= 0; i-- {
		for j := 1; j <= byte2Num(s[i]); j++ {
			if i == 2 {
				fmt.Printf("B")
			} else if i == 1 {
				fmt.Printf("S")
			} else {
				fmt.Printf("%d", j)
			}
		}
	}
}
