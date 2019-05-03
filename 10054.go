package main

import (
	"fmt"
	"math"
	"strconv"
	"strings"
	"unicode"
)

func IsLegel(s string) bool {
	for i := 0; i < len(s); i++ {
		if unicode.IsLetter(rune(s[i])) {
			return false
		}
	}
	return true
}

func main() {
	var N int
	var s string
	var number int
	var sum float64
	fmt.Scanf("%d", &N)
	for i := 0; i < N; i++ {
		fmt.Scanf("%s", &s)
		if !IsLegel(s) {
			fmt.Printf("ERROR: %s is not a legal number\n", s)
		} else {
			lis := strings.Split(s, ".")
			if len(lis) > 2 {
				fmt.Printf("ERROR: %s is not a legal number\n", s)
			} else {
				if len(lis) == 1 {
					t, _ := strconv.Atoi(lis[0])
					if t <= 1000 && t >= -1000 {
						number++
						sum += float64(t)
					} else {
						fmt.Printf("ERROR: %s is not a legal number\n", s)
					}
				} else {
					t := len(lis[1])
					if t >= 3 {
						fmt.Printf("ERROR: %s is not a legal number\n", s)
					} else {
						t1, _ := strconv.Atoi(lis[0])
						if t1 < -1000 || t1 > 1000 {
							fmt.Printf("ERROR: %s is not a legal number\n", s)
						} else {
							t2, _ := strconv.Atoi(lis[0])
							t3, _ := strconv.Atoi(lis[1])
							number++
							sum += float64(t2)
							if lis[0][0] == '-' {
								sum -= math.Pow(0.1, float64(t)) * float64(t3)
							} else {
								sum += math.Pow(0.1, float64(t)) * float64(t3)
							}
						}
					}
				}
			}
		}
	}
	if number == 0 {
		fmt.Printf("The average of 0 numbers is Undefined\n")
	} else {
		fmt.Printf("The average of %d numbers is %.2f\n", number, sum/float64(number))
	}
}
