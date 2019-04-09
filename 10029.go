package main

import (
	"fmt"
	"strings"
)

func main() {
	var s1, s2 string
	fmt.Scanf("%s\n%s", &s1, &s2)
	var t []byte
	for i, j := 0, 0; i < len(s1) && j < len(s2); i++ {
		if s1[i] != s2[j] {
			mask := 0
			for k := range t {
				if strings.ToUpper(string(s1[i])) == strings.ToUpper(string(t[k])) {
					mask = 1
				}
			}
			if mask == 0 {
				t = append(t, s1[i])
			}
		} else {
			j++
		}
	}
	s := string(t)
	s = strings.ToUpper(s)
	fmt.Println(s)
}
