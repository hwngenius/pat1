package main

import "fmt"

//0 no,1 win 2 loss
func cmp(x, y byte) (r int) {
	switch x {
	case 'C':
		switch y {
		case 'C':
			r = 0
		case 'J':
			r = 1
		case 'B':
			r = 2
		}
	case 'J':
		switch y {
		case 'C':
			r = 2
		case 'J':
			r = 0
		case 'B':
			r = 1
		}
	case 'B':
		switch y {
		case 'C':
			r = 1
		case 'J':
			r = 2
		case 'B':
			r = 0
		}
	}
	return
}

func main() {
	var n, no, win, loss int
	xN := make(map[byte]int)
	yN := make(map[byte]int)
	xN['B'] = 0
	xN['C'] = 0
	xN['J'] = 0
	yN['B'] = 0
	yN['C'] = 0
	yN['J'] = 0
	var x, y byte
	fmt.Scanf("%d", &n)
	for i := 0; i < n; i++ {
		fmt.Scanf("%c %c\n", &x, &y)
		t := cmp(x, y)
		if t == 0 {
			no++
		} else if t == 1 {
			xN[x]++
			win++
		} else if t == 2 {
			loss++
			yN[y]++
		}
	}
	fmt.Println(win, no, loss)
	fmt.Println(loss, no, win)
	var l1, l2 byte
	var max int
	for i1 := range xN {
		//fmt.Printf("%c %d\n", i1, xN[i1])
		if xN[i1] > max {
			max = xN[i1]
			l1 = i1
		}
	}
	max = 0
	for i1 := range yN {
		//fmt.Printf("%c %d\n", i1, yN[i1])
		if yN[i1] > max {
			max = yN[i1]
			l2 = i1
		}
	}
	fmt.Printf("%c %c", l1, l2)
}
