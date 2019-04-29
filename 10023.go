package main

import "fmt"

func main() {
	var p [10]int
	for i := 0; i < 10; i++ {
		fmt.Scanf("%d", &p[i])
	}
	//fmt.Println(p)
	for i := 1; i < 10; i++ {
		if p[i] != 0 {
			fmt.Printf("%d", i)
			p[i]--
			break
		}
	}
	//fmt.Println(p)
	for i := 0; i < 10; i++ {
		for j := p[i]; j > 0; j-- {
			fmt.Printf("%d", i)
		}
	}
}
