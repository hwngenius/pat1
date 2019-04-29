package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func int2byte(n int) (r byte) {
	switch n {
	case 0:
		r = 'B'
	case 1:
		r = 'C'
	case 2:
		r = 'J'
	default:
		r = ','
	}
	return
}
func byte2int(b byte) (n int) {
	switch b {
	case 'B':
		n = 0
	case 'C':
		n = 1
	case 'J':
		n = 2
	default:
		n = 0
	}
	return
}
func play(a1, a2 byte) (n int) { //-1,a1 win;0,none win;1,a2 win.
	switch a1 {
	case 'B':
		switch a2 {
		case 'B':
			n = 0
		case 'C':
			n = -1
		case 'J':
			n = 1
		}
	case 'J':
		switch a2 {
		case 'B':
			n = -1
		case 'C':
			n = 1
		case 'J':
			n = 0
		}
	case 'C':
		switch a2 {
		case 'B':
			n = 1
		case 'C':
			n = 0
		case 'J':
			n = -1
		}
	}
	return
}
func main() {
	var N int
	var a1, a2 byte
	var a1Win, a2Win [3]int
	var win, loss, nowin int
	input := bufio.NewScanner(os.Stdin)
	fmt.Scan(&N)
	for input.Scan() {
		l := strings.Split(input.Text(), " ")
		a1 = l[0][0]
		a2 = l[1][0]
		t := play(a1, a2)
		if t == 0 {
			nowin++
		} else if t == 1 {
			loss++
			a2Win[byte2int(a2)]++
		} else {
			win++
			a1Win[byte2int(a1)]++
		}
	}
	fmt.Println(win, nowin, loss)
	fmt.Println(loss, nowin, win)
	var max1, max2 int
	var i1, i2 int
	for i := 0; i < 3; i++ {
		if max1 < a1Win[i] {
			max1 = a1Win[i]
			i1 = i
		}
	}
	for i := 0; i < 3; i++ {
		if max2 < a2Win[i] {
			max2 = a2Win[i]
			i2 = i
		}
	}
	// fmt.Println(a1Win)
	// fmt.Println(a2Win)
	fmt.Printf("%c %c", int2byte(i1), int2byte(i2))
}
