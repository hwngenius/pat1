package main

import "fmt"

func main() {
	var mask [100000]bool
	var n int
	fmt.Scan(&n)
	for i := 2; i <= n; i++ {
		if !mask[i] {
			for j := i * 2; j <= n; j += i {
				mask[j] = true
			}
		}
	}
	sum := 0
	for i := 2; i <= n; i++ {
		if !mask[i] && !mask[i+2] && i+2 <= n {
			sum++
		}
	}
	fmt.Printf("%d", sum)
}
