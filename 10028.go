package main

import "fmt"

func isright(y, m, d int) (i bool) {
	if y > 2014 {
		i = false
	} else if y == 2014 {
		if m > 9 {
			i = false
		} else if m == 9 {
			if d > 6 {
				i = false
			} else {
				i = true
			}
		}
	} else if y < 1814 {
		i = false
	} else if y == 1814 {
		if m < 9 {
			i = false
		} else if m == 9 {
			if d < 6 {
				i = false
			} else {
				i = true
			}
		}
	} else {
		i = true
	}

	return
}

func young(y1, m1, d1, y2, m2, d2 int) (i bool) {
	if y1 != y2 {
		return y1 > y2
	} else if m1 != m2 {
		return m1 > m2
	} else {
		return d1 > d2
	}
}

type birth struct {
	name string
	y    int
	m    int
	d    int
}

func main() {
	var sample, y, o birth
	sample = birth{"Tom", 1814, 9, 6}
	var n, sum, mask int = 5, 0, 0
	fmt.Scanf("%d", &n)
	for i := 0; i < n; i++ {
		fmt.Scanf("%s %d/%d/%d", &sample.name, &sample.y, &sample.m, &sample.d)
		if isright(sample.y, sample.m, sample.d) {
			if mask == 0 {
				y = sample
				o = sample
				mask = 1
			}
			sum++
			// fmt.Println(sum, sample)
			if !young(sample.y, sample.m, sample.d, o.y, o.m, o.d) {
				o = sample
				//fmt.Println(o)
			} else if young(sample.y, sample.m, sample.d, y.y, y.m, y.d) {
				y = sample
				//fmt.Println(y)
			}
		}
	}
	fmt.Printf("%d %s %s", sum, o.name, y.name)
}
