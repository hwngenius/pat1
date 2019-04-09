package main

import (
	"fmt"
	"strconv"
)

var weight = []int{7, 9, 10, 5, 8, 4, 2, 1, 6, 3, 7, 9, 10, 5, 8, 4, 2}
var change = []byte{'1', '0', 'X', '9', '8', '7', '6', '5', '4', '3', '2'}

func isRight(s string) (b bool) {
	_, t := strconv.Atoi(s)
	if t != nil {
		return false
	} else {
		sum := 0
		for i := 0; i < 17; i++ {
			n, _ := strconv.Atoi(string(s[i]))
			sum += n * weight[i]
		}
		sum %= 11
		if change[sum] == s[17] {
			return true
		} else {
			return false
		}
	}
}

func main() {
	var N int = 1
	var s string = "320124198808240056"
	var mask bool
	fmt.Scanf("%d", &N)
	for i := 0; i < N; i++ {
		fmt.Scanf("%s", &s)
		if !isRight(s) {
			mask = true
			fmt.Println(s)
		}
	}
	if mask == false {
		fmt.Printf("All passed")
	}
}
