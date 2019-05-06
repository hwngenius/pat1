package main

import (
	"fmt"
	"math"
)

func main() {
	var N int
	var sum float64
	var a, b int
	fmt.Scanf("%d", &N)
	for i := 0; i < N; i++ {
		fmt.Scanf("%d %d", &a, &b)
		s := math.Sqrt(float64(a*a + b*b))
		if s > sum {
			sum = s
		}
	}
	fmt.Printf("%.2f", sum)
}
