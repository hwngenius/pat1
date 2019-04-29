package main

import (
	"fmt"
	"sort"
)

func main() {
	var N, p int
	var input int
	var sample []int
	fmt.Scan(&N)
	fmt.Scan(&p)
	for i := 0; i < N; i++ {
		fmt.Scan(&input)
		sample = append(sample, input)
	}
	sort.Ints(sample)
	i := 0
	for i = 0; i < len(sample); i++ {
		if sample[i]*p >= sample[len(sample)-1] {
			fmt.Println(len(sample) - i)
			break
		}
	}
	if i == len(sample) {
		fmt.Println(0)
	}
}
