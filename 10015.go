package main

import "fmt"

type Stu struct {
	id   string
	de   int
	cai  int
	zong int
	flag int
}

func cmp(s1, s2 Stu) (r bool) {
	if s1.flag != s2.flag {
		r = s1.flag < s2.flag
	} else if s1.zong == s2.zong {
		if s1.de == s2.de {
			r = s1.id < s2.id
		} else {
			r = s1.de > s2.de
		}
	} else {
		r = s1.zong > s2.zong
	}
	return
}
func sort(sample []Stu) {
	for i := 1; i < len(sample); i++ {
		for j := len(sample) - 1; j >= i; j-- {
			if cmp(sample[j], sample[j-1]) {
				sample[j], sample[j-1] = sample[j-1], sample[j]
			}
		}
	}
}

func main() {
	var N, H, L, sum int
	fmt.Scanf("%d%d%d", &N, &L, &H)
	var sample []Stu
	var demo Stu
	for i := 0; i < N; i++ {
		fmt.Scanf("%s%d%d", &demo.id, &demo.de, &demo.cai)
		demo.zong = demo.de + demo.cai
		if demo.de >= H && demo.cai >= H {
			demo.flag = 1
			sum++
		} else if demo.de >= H && demo.cai >= L {
			demo.flag = 2
			sum++
		} else if demo.de >= L && demo.cai >= L && demo.de > demo.cai {
			demo.flag = 3
			sum++
		} else if demo.de >= L && demo.cai >= L {
			demo.flag = 4
			sum++
		} else {
			demo.flag = 5
		}
		sample = append(sample, demo)
	}
	sort(sample)
	fmt.Println(sum)
	for i := 0; i < sum; i++ {
		fmt.Println(sample[i].id, sample[i].de, sample[i].cai)
	}
}
