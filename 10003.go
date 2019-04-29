package main

import "fmt"

func main() {
	var n int
	var s string
	fmt.Scan(&n)
	for i := 0; i < n; i++ {
		fmt.Scan(&s)
		var loc_p, loc_t, other, num_p, num_t int
		for j := 0; j < len(s); j++ {
			if s[j] == 'P' {
				loc_p = j
				num_p++
			} else if s[j] == 'T' {
				loc_t = j
				num_t++
			} else if s[j] != 'A' {
				other++
			}
		}
		if other > 0 || num_p != 1 || num_t != 1 || loc_t-loc_p <= 1 {
			fmt.Printf("NO\n")
			continue
		}
		x := loc_p
		y := loc_t - loc_p - 1
		z := len(s) - loc_t - 1
		if z-x*(y-1) == x {
			fmt.Println("YES")
		} else {
			fmt.Println("NO")
		}
	}
}
