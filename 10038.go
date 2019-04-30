package main

import "fmt"

func main() {
	var N,k int
	var sample [110]int
	var n int
	fmt.Scanf("%d", &N)
	for i := 0; i < N; i++ {
		fmt.Scanf("%d", &n)
		sample[n]++
	}
	fmt.Scanf("%d", &k)
	for i := 0; i < k; i++ {
		fmt.Scan(&n)
		if i > 0 {
			fmt.Printf(" ")
		}
		fmt.Printf("%d", sample[n])

	}
}
