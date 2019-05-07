package main

import "fmt"

func main() {
	var N, M, t int
	fmt.Scanf("%d %d", &N, &M)
	var set [1010]int
	for i := 0; i < M; i++ {
		for j := 0; j < N; j++ {
			fmt.Scan(&t)
			set[j] += t
		}
	}
	var max int
	var result []int
	for i := 0; i < N; i++ {
		if set[i] > max {
			max = set[i]
			result = nil
			result = append(result, i+1)
		} else if max == set[i] {
			result = append(result, i+1)
		}
	}
	fmt.Println(max)
	for i := 0; i < len(result); i++ {
		if i > 0 {
			fmt.Printf(" ")
		}
		fmt.Printf("%d", result[i])
	}
}
