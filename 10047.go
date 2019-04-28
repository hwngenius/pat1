package main

import "fmt"

func main() {
	var team [1001]int
	var N int
	var t, n, g int
	var maxScore, winTeam int
	fmt.Scanf("%d", &N)
	for i := 0; i < N; i++ {
		fmt.Scanf("%d-%d %d", &t, &n, &g)
		team[t] += g
		if maxScore < team[t] {
			maxScore = team[t]
			winTeam = t
		}
	}
	fmt.Println(winTeam, maxScore)
}
