package main

import "fmt"

func main() {
	var N, M int
	var G, G1 int
	fmt.Scanf("%d %d", &N, &M)
	for i := 0; i < N; i++ {
		var sum int
		max := 0
		min := M
		n := 0
		fmt.Scanf("%d", &G1)
		//fmt.Println(G1)
		for j := 0; j < N-1; j++ {
			fmt.Scan(&G)
			if G <= M && G >= 0 {
				sum += G
				n++
				//fmt.Println(G)
				if G > max {
					max = G
				}
				if G < min {
					min = G
				}
			}
		}
		sum = sum - max - min
		n -= 2
		sum1 := float64(sum) / float64(n)
		//fmt.Println(sum1)
		result := int((sum1+float64(G1))/float64(2) + 0.5)
		fmt.Println(result)
	}
}
