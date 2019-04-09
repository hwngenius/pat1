package main

import "fmt"

func main() {
	var n, a, b, c int
	var mask bool = false
	fmt.Scan(&n)
	for i := 1; i <= n; i++ {
		fmt.Scanf("%d%d%d", &a, &b, &c)
		if a+b > c {
			if mask {
				fmt.Printf("\n")
			}
			fmt.Printf("Case #%d: true", i)
		} else {
			if mask {
				fmt.Printf("\n")
			}
			fmt.Printf("Case #%d: false", i)
		}
		if mask == false {
			mask = true
		}
	}
}
