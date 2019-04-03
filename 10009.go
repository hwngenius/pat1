package main

import "fmt"

func main() {
	var s string
	t, _ := fmt.Scanf("%s", &s)
	defer fmt.Printf("%s", s)
	for t, _ = fmt.Scanf("%s", &s); t != 0; t, _ = fmt.Scanf("%s", &s) {
		defer fmt.Printf("%s ", s)
	}
	fmt.Printf("\b")
}
