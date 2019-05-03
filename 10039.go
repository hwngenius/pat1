package main

import "fmt"

func main() {
	m := make(map[byte]int)
	var buy, want string
	var loss int
	fmt.Scanf("%s\n%s", &buy, &want)
	for i := 0; i < len(buy); i++ {
		m[buy[i]]++
	}
	for i := 0; i < len(want); i++ {
		m[want[i]]--
		if m[want[i]] < 0 {
			loss++
		}
	}
	if loss != 0 {
		fmt.Printf("No %d", loss)
	} else {
		fmt.Printf("Yes %d", len(buy)-len(want))
	}
}
