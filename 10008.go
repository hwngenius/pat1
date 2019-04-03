package main

import "fmt"

func main() {
	var N, M int
	var p [100]int
	fmt.Scanf("%d%d", &N, &M)

	for i := 0; i < N; i++ {
		fmt.Scanf("%d", &p[i])
	}
	M %= N
	j := 0
	for i := N - M; i < N; i++ {
		fmt.Printf("%d", p[i])
		j++
		if j != N {
			fmt.Printf(" ")
		}
	}
	for i := 0; i < N-M; i++ {
		fmt.Printf("%d", p[i])
		j++
		if j != N {
			fmt.Printf(" ")
		}
	}
}
