package main

import "fmt"

func main() {
	var A, B, S, D, r int
	fmt.Scanf("%d %d %d", &A, &B, &D)
	S = A + B
	r = D
	for {
		if r <= S {
			r *= D
		} else {
			r /= D
			break
		}
	}
	if S == 0 {
		fmt.Printf("0")
	}
	for S != 0 {
		fmt.Printf("%d", S/r)
		S %= r
		r /= D
	}
	for ; r != 0; r /= D {
		fmt.Printf("0")
	}
}
