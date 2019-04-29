package main

import "fmt"

func reserve(n []int) {
	for i, j := 0, len(n)-1; i < j; i, j = i+1, j-1 {
		n[i], n[j] = n[j], n[i]
	}
}

func main() {
	var N, M, input int
	var sample []int
	fmt.Scanf("%d %d", &N, &M)
	for i := 0; i < N; i++ {
		fmt.Scan(&input)
		sample = append(sample, input)
	}
	reserve(sample[N-M%N:])
	reserve(sample[:N-M%N])
	reserve(sample[:])
	for i := 0; i < len(sample); i++ {
		if i > 0 {
			fmt.Printf(" ")
		}
		fmt.Printf("%d", sample[i])
	}
}
