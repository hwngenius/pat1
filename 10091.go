package main

import "fmt"

func main() {
	var M int
	fmt.Scan(&M)
	var t int
	for i := 0; i < M; i++ {
		fmt.Scan(&t)
		t1 := t * t
		t2 := 0
		if t >= 1000 {
			t2 = 10000
		} else if t >= 100 {
			t2 = 1000
		} else if t >= 10 {
			t2 = 100
		} else {
			t2 = 10
		}
		mask := 0
		j := 0
		for j = 1; j < 10; j++ {
			if (t1*j)%t2 == t {
				mask = 1
				break
			}
		}
		if mask == 0 {
			fmt.Println("No")
		} else {
			fmt.Println(j, t1*j)
			mask = 0
		}
	}
}
