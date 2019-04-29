package main

import "fmt"

type person struct {
	yy int
	mm int
	dd int
}

func older(a, b person) bool {
	if a.yy != b.yy {
		return (a.yy) < (b.yy)
	} else if a.mm != b.mm {
		return (a.mm) < (b.mm)
	} else {
		return (a.dd) <= (b.dd)
	}
}

func main() {
	var input person
	var oldest = person{2014, 9, 6}
	var youngest = person{1814, 9, 6}
	var left = youngest
	var right = oldest
	var N, sum int
	var name, oldname, youngname string

	fmt.Scan(&N)
	for i := 0; i < N; i++ {
		fmt.Scanf("%s %d/%d/%d", &name, &input.yy, &input.mm, &input.dd)
		if older(input, right) && older(left, input) {
			sum++
			if older(input, oldest) {
				oldname = name
				oldest = input
			}
			if older(youngest, input) {
				youngname = name
				youngest = input
			}
		}
	}
	if sum != 0 {
		fmt.Println(sum, oldname, youngname)
	} else {
		fmt.Println(0)
	}
}
