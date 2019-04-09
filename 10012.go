package main

import "fmt"

func main() {
	var A [5]int
	var flag [5]bool
	var m, sum, n, demo int
	fmt.Scan(&n)
	for i := 0; i < n; i++ {
		fmt.Scan(&demo)
		switch demo % 5 {
		case 0:
			if demo%2 == 0 {
				flag[0] = true
				A[0] += demo
			}
		case 1:
			flag[1] = true
			if m%2 == 0 {
				A[1] += demo
			} else {
				A[1] -= demo
			}
			m++
		case 2:
			flag[2] = true
			A[2] += 1
		case 3:
			flag[3] = true
			A[3] += demo
			sum += 1
		case 4:
			flag[4] = true
			if A[4] < demo {
				A[4] = demo
			}
		}
	}
	mask := 0
	for i := 0; i < 5; i++ {
		if mask == 0 {
			mask = 1
		} else {
			fmt.Printf(" ")
		}
		if flag[i] == false {
			fmt.Printf("N")
		} else if i == 3 {
			t := float32(A[3]) / float32(sum)
			fmt.Printf("%.1f", t)
		} else {
			fmt.Printf("%d", A[i])
		}
	}
}
