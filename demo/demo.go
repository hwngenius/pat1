package main

import "fmt"

func reserve(set []int) {
	for i, j := 0, len(set)-1; i < j; i, j = i+1, j-1 {
		set[i], set[j] = set[j], set[i]
	}
}

func main() {
	var N, K, input int
	fmt.Scan(&N)
	fmt.Scan(&K)
	var set []int
	for i := 0; i < N; i++ {
		fmt.Scan(&input)
		set = append(set, input)
	}
	K %= N
	reserve(set[N-K:])
	reserve(set[:N-K])
	reserve(set)
	for i := 0; i < len(set); i++ {
		if i > 0 {
			fmt.Printf(" ")
		}
		fmt.Printf("%d", set[i])
	}
}
