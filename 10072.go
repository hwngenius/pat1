package main

import "fmt"

func isThere(L []string, s string) bool {
	for i := 0; i < len(L); i++ {
		if s == L[i] {
			return true
		}
	}
	return false
}

func main() {
	var N, M int
	var name, tmp string
	var goodsNum int
	var first bool
	var person, good int
	fmt.Scanf("%d %d", &N, &M)
	var forbidden []string
	for i := 0; i < M; i++ {
		s := ""
		fmt.Scanf("%s", &s)
		forbidden = append(forbidden, s)
	}
	for i := 0; i < N; i++ {
		mask := 0
		fmt.Scanf("%s %d", &name, &goodsNum)
		for j := 0; j < goodsNum; j++ {
			fmt.Scan(&tmp)
			if isThere(forbidden, tmp) {
				if mask == 0 {
					if first == false {
						first = true
					} else {
						fmt.Println()
					}
					person++
					fmt.Printf("%s:", name)
					mask = 1
				}
				fmt.Printf(" %s", tmp)
				good++
			}
		}
	}
	if person != 0 {
		fmt.Printf("\n%d %d", person, good)
	} else {
		fmt.Println(person, good)
	}
}
