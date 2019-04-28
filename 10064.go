package main

import (
	"fmt"
	"sort"
)

func friendNum(s string) (r int) {
	for i := 0; i < len(s); i++ {
		r += int(s[i]) - int('1') + 1
	}
	return
}

func main() {
	var set [10000]int
	var N, sum int
	var sample string
	var result []int
	fmt.Scanf("%d", &N)
	for i := 0; i < N; i++ {
		fmt.Scanf("%s", &sample)
		t := friendNum(sample)
		set[t]++
		if set[t] == 1 {
			sum++
			result = append(result, t)
		}
	}
	sort.Ints(result)
	fmt.Println(sum)
	fmt.Printf("%d", result[0])
	for i := 1; i < len(result); i++ {
		fmt.Printf(" %d", result[i])
	}
}
