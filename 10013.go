package main

import "fmt"

func main() {
	var M, N int
	var set [1000000]bool
	for i := 2; i < 1000000; i++ {
		if set[i] == false {
			for j := i + i; j < 1000000; j += i {
				set[j] = true
			}
		}
	}
	fmt.Scanf("%d %d", &M, &N)
	var count, pt int
	for i := 2; i < 1000000; i++ {
		if set[i] == false {
			count++
			if count >= M && count <= N {
				pt++
				if pt%10 != 1 {
					fmt.Printf(" ")
				}
				fmt.Printf("%d", i)
				if pt%10 == 0 {
					fmt.Println()
				}
			}
		}
	}
}
