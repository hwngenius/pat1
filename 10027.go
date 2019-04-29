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
	for i := 1; i <= m; i++ {
		for j := 1; j < i; j++ {
			fmt.Printf(" ")
		}
		for j := 0; j < 2*m-1-2*(i-1); j++ {
			fmt.Printf("%c", s)
		}
		fmt.Printf("\n")
	}
	t := 3
	for i := 1; i < m; i++ {
		for j := 0; j < (2*m-1-t)/2; j++ {
			fmt.Printf(" ")
		}
		for j := 0; j < t; j++ {
			fmt.Printf("%c", s)
		}
		fmt.Println()
		t += 2
	}

	fmt.Println(n)
}
