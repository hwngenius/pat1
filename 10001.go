package main

import "fmt"

func main() {
	var n, count int = 3, 0
	fmt.Scan(&n)
	for {
		if n == 1 {
			break
		} else if n%2 == 1 {
			n = (3*n + 1) / 2
			count += 1
		} else {
			n /= 2
			count++
		}
	}
	fmt.Println(count)
}
