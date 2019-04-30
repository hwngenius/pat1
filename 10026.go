package main

import "fmt"

func main() {
	var start, end, value int
	var h, m, s int
	fmt.Scanf("%d%d", &start, &end)
	value = end - start + 50
	value = int(float32(value) / float32(100))
	h = value / 3600
	value %= 3600
	m = value / 60
	value %= 60
	s = value
	fmt.Printf("%02d:%02d:%02d", h, m, s)
}
