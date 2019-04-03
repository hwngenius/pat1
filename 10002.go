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

func num2Spell(n int) (s string) {
	switch n {
	case 0:
		s = "ling"
	case 1:
		s = "yi"
	case 2:
		s = "er"
	case 3:
		s = "san"
	case 4:
		s = "si"
	case 5:
		s = "wu"
	case 6:
		s = "liu"
	case 7:
		s = "qi"
	case 8:
		s = "ba"
	case 9:
		s = "jiu"
	}
	return
}
func main() {
	var s string
	fmt.Scan(&s)
	sum := 0
	for i := 0; i < len(s); i++ {
		sum += byte2Num(s[i])
	}
	mask := 0
	for sum > 0 {
		s = num2Spell(sum % 10)
		sum /= 10
		if mask == 0 {
			mask = 1
			defer fmt.Printf("%s", s)
		} else {
			defer fmt.Printf("%s ", s)
		}
	}
}
