package main

import "fmt"

func main() {
	var N, t int
	var set [10010]int
	fmt.Scan(&N)
	for i := 0; i < N; i++ {
		fmt.Scan(&t)
		t1 := t - (i + 1)
		if t1 < 0 {
			t1 = -t1
		}
		set[t1]++
	}
	for i := N; i >= 0; i-- {
		if set[i] > 1 {
			fmt.Println(i, set[i])
		}
	}
}
