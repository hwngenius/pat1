package main

import "fmt"

func main() {
	var n, m, mask int
	var input [1010]int
	for t, _ := fmt.Scanf("%d%d", &n, &m); t == 2; t, _ = fmt.Scanf("%d%d", &n, &m) {
		if m != 0 {
			input[m] += n
		}
	}
	for i := 1000; i >= 0; i-- {
		if input[i] != 0 {
			if mask == 0 {
				fmt.Printf("%d %d", i*input[i], i-1)
				mask = 1
			} else {
				fmt.Printf(" %d %d", i*input[i], i-1)
			}
		}
	}
	if mask == 0 {
		fmt.Printf("0 0")
	}
}
