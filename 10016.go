package main

import (
	"fmt"
)

func main() {
	var s1, s2 string
	var i1, i2 int
	var sum1, sum2 int
	fmt.Scanf("%s%d%s%d", &s1, &i1, &s2, &i2)
	s := rune('0') + rune(i1)
	for i := 0; i < len(s1); i++ {
		if rune(s1[i]) == s {
			sum1 = sum1*10 + i1
		}
	}
	s = rune('0') + rune(i2)
	for i := 0; i < len(s2); i++ {
		if rune(s2[i]) == s {
			sum2 = sum2*10 + i2
		}
	}
	fmt.Printf("%d", sum1+sum2)
}
