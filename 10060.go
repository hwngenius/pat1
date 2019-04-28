package main

import (
	"fmt"
	"sort"
)

func main() {
	var N int
	var sample []int
	fmt.Scan(&N)
	var m int
	for i := 0; i < N; i++ {
		fmt.Scan(&m)
		sample = append(sample, m)
	}
	sort.Ints(sample)
	var sum int
	for i := 0; i < N; i++ {
		if N-i+1 <= sample[i] {
			sum = N - i
			break
		}
	}
	fmt.Println(sum)
}
