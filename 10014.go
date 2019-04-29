package main

import (
	"fmt"
	"unicode"
)

func main() {
	var WEEK = [...]string{"MON", "TUE", "WED", "THU", "FRI", "SAT", "SUN"}
	var s1, s2, s3, s4 string
	var mask int
	var week, hour, min int
	fmt.Scanf("%s\n%s\n%s\n%s", &s1, &s2, &s3, &s4)
	for i, j := 0, 0; i < len(s1) && j < len(s2); i, j = i+1, j+1 {
		if s1[i] == s2[j] && mask <= 1 {
			if mask == 0 {
				if s1[i] >= 'A' && s1[i] <= 'G' {
					week = int(s1[i] - 'A' + 1)
					mask = 1
				}
			} else {
				if unicode.IsDigit(rune(s1[i])) {
					hour = int(s1[i] - '0')
					mask++
				} else if s1[i] >= 'A' && s1[i] <= 'N' {
					hour = int(s1[i]-'A') + 10
					mask++
				}
			}
		}
	}
	for i, j := 0, 0; i < len(s3) && j < len(s4); i, j = i+1, j+1 {
		if s3[i] == s4[j] && unicode.IsLetter(rune(s3[i])) {
			min = i
			break
		}
	}
	//fmt.Println(week, hour, min)
	fmt.Printf("%s %02d:%02d", WEEK[week-1], hour, min)
}
