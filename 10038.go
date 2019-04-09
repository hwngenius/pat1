package main

import "fmt"

func main() {
	var N int
	var sample [100]int
	var n int
	fmt.Scanf("%d", &N)
	for i := 0; i < N; i++ {
		fmt.Scanf("%d", &n)
		sample[n]++
	}
	fmt.Scanf("%d", &N)
	for i := 0; i < N; i++ {
		fmt.Scan(&n)
		if i > 0 {
			fmt.Printf(" ")
		}
		fmt.Printf("%d", sample[n])

	}
}
