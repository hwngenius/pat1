package main

import "fmt"

func cmp(x, y byte) (result int) {
	switch x {
	case 'C':
		switch y {
		case 'C':
			result = 1
		case 'J':
			result = 0
		case 'B':
			result = 2
		}
	case 'J':
		switch y {
		case 'C':
			result = 2
		case 'J':
			result = 1
		case 'B':
			result = 0
		}
	case 'B':
		switch y {
		case 'C':
			result = 0
		case 'J':
			result = 2
		case 'B':
			result = 1
		}
	}
	return
}

func Char2Num(c byte) (result int) {
	switch c {
	case 'C':
		result = 0
	case 'J':
		result = 1
	case 'B':
		result = 2
	}
	return
}

func Num2Char(n int) (c byte) {
	switch n {
	case 0:
		c = 'C'
	case 1:
		c = 'J'
	case 2:
		c = 'B'
	}
	return
}

func main() {
	var n int
	var x, y byte
	var xN, yN [3]int
	var win, loss, no int
	fmt.Scanf("%d", &n)
	for i := 0; i < n; i++ {
		fmt.Scanf("%c%c", &x, &y)
		xN[Char2Num(x)]++
		yN[Char2Num(y)]++
		if cmp(x, y) == 0 {
			win++
		} else if cmp(x, y) == 1 {
			loss++
		} else {
			no++
		}
	}
	fmt.Printf("%d %d %d\n", win, loss, no)
	index1, max := 0, 0
	for i := 0; i < 3; i++ {
		if xN[i] > max {
			max = xN[i]
			index1 = i
		}
	}
	index2, max := 0, 0
	for i := 0; i < 3; i++ {
		if yN[i] > max {
			max = yN[i]
			index2 = i
		}
	}
	fmt.Printf("%c %c", Num2Char(index1), Num2Char(index2))
}
