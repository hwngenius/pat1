package main

import (
	"fmt"
	"strconv"
)

func f(s string) string {
	var t []byte
	for i := len(s) - 1; i >= 0; i-- {
		t = append(t, s[i])
	}
	return string(t)
}
func reverse(n int) int {
	s := f(strconv.Itoa(n))
	t, _ := strconv.Atoi(s)
	return t
}

func main() {
	var M, x, y int = 48, 3, 7
	fmt.Scanf("%d %d %d", &M, &x, &y)
	var i, t, t1 int
	for i = 99; i > 9; i-- { //i甲
		t = reverse(i)
		t1 = 0 //t乙
		if t > i {
			t1 = t - i
		} else {
			t1 = i - t
		}
		if t/y == t1/x && t%y == 0 && t1%x == 0 {
			break
		}
	}
	if i < 10 {
		fmt.Println("No Solution")
	} else {
		fmt.Printf("%d", i)
		if i == M {
			fmt.Printf(" Ping")
		} else if i < M {
			fmt.Printf(" Gai")
		} else {
			fmt.Printf(" Cong")
		}
		if t == M {
			fmt.Printf(" Ping")
		} else if t < M {
			fmt.Printf(" Gai")
		} else {
			fmt.Printf(" Cong")
		}
		if t/y == M {
			fmt.Printf(" Ping")
		} else if t/y < M {
			fmt.Printf(" Gai")
		} else {
			fmt.Printf(" Cong")
		}
	}

}
