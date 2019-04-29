package main

import "fmt"

func main() {
	var N, M, t int
	fmt.Scanf("%d %d", &N, &M)
	var score, is []int
	for i := 0; i < M; i++ {
		fmt.Scanf("%d", &t)
		score = append(score, t)
	}
	for i := 0; i < M; i++ {
		fmt.Scanf("%d", &t)
		is = append(is, t)
	}
	sum := 0
	for i := 0; i < N; i++ {
		for j := 0; j < M; j++ {
			fmt.Scanf("%d", &t)
			if t == is[j] {
				sum += score[j]
			}
		}
		fmt.Println(sum)
		sum = 0
	}
}
