package main

import (
	"fmt"
	"math"
)

func main() {
	var N int
	var ID string
	var x, y int
	var nMax, nMin float64
	var sMax, sMin string
	nMax = 0
	nMin = 200
	fmt.Scan(&N)
	for i := 0; i < N; i++ {
		fmt.Scanf("%s %d %d", &ID, &x, &y)
		t := math.Sqrt(float64(x*x + y*y))
		if nMin > t {
			nMin = t
			sMin = ID
		}
		if nMax < t {
			nMax = t
			sMax = ID
		}
	}
	fmt.Println(sMin, sMax)
}
