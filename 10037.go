package main

import "fmt"

func main() {
	var Galleon1, Galleon2, Sickle1, Sickle2, Knut1, Knut2 int
	fmt.Scanf("%d.%d.%d %d.%d.%d.%d", &Galleon1, &Sickle1, &Knut1, &Galleon2, &Sickle2, &Knut2)
	sum2 := Galleon2*17*29 + Sickle2*29 + Knut2
	sum1 := Galleon1*17*29 + Sickle1*29 + Knut1
	if sum1 == sum2 {
		fmt.Printf("0.0.0")
		return
	}
	var diff int
	if sum2 > sum1 {
		diff = sum2 - sum1
	} else {
		fmt.Printf("-")
		diff = sum1 - sum2
	}
	fmt.Printf("%d", diff/(29*17))
	diff %= 29 * 17
	fmt.Printf(".%d", diff/29)
	diff %= 29
	fmt.Printf(".%d", diff)
}
