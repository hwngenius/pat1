package main

import "fmt"

type Student struct {
	name  string
	id    string
	score int
}

func main() {
	var n int
	var top, low, sample Student
	fmt.Scan(&n)
	fmt.Scan(&sample)
	low = sample
	top = sample
	for ; n > 1; n-- {
		fmt.Scan(&sample)
		if sample.score > top.score {
			top = sample
		} else if sample.score < low.score {
			low = sample
		}
	}
	fmt.Printf("%s %s\n%s %s", top.name, top.id, low.name, low.id)
}
