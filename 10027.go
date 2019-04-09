package main

import "fmt"

func fun(n int) (i, r int) {
	if n == 1 {
		i = 1
		r = 1
	} else {
		for i, r = 2, 1; r <= n; i++ {
			r += 2 * (2*i - 1)
		}
		i--
		r -= 2 * (2*i - 1)
		i--
	}
	return
}

func main() {
	var n int = 7
	var s byte = '*'
	fmt.Scanf("%d %c", &n, &s)
	m, r := fun(n)
	n -= r
	if m == 1 {
		fmt.Printf("%c\n%d", s, n)
		return
	}
	for i := 1; i <= 2*m-1; i++ {
		for j := 1; j <= 2*m-1; j++ {
			if (i <= j && j+i <= 2*m) || (j <= i && i+j >= 2*m) {
				fmt.Printf("%c", s)
			} else {
				fmt.Printf(" ")
			}
		}
		fmt.Printf("\n")
	}
	fmt.Printf("%d", n)
}
