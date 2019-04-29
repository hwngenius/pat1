package main

import (
	"fmt"
	"sort"
)

func main() {
	var hash [10000]bool
	var num [110]int
	var K, input int
	fmt.Scan(&K)
	for i := 0; i < K; i++ {
		fmt.Scan(&input)
		num[i] = input
		for input != 1 {
			if input%2 == 1 {
				input = (3*input + 1) / 2
			} else {
				input /= 2
			}
			hash[input] = true
		}
	}
	var ouput []int
	for i := 0; i < K; i++ {
		if !hash[num[i]] {
			ouput = append(ouput, num[i])
		}
	}
	sort.Ints(ouput)
	fmt.Printf("%d", ouput[len(ouput)-1])
	for i := len(ouput) - 2; i >= 0; i-- {
		fmt.Printf(" %d", ouput[i])
	}
}
