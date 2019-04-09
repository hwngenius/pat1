package main

import "fmt"

type Stu struct {
	id   string
	de   int
	cai  int
	sum int
	flag int
}

bool cmp(s1 Stu,s2 Stu)(result bool){
	if s1.flag!=s2.flag{
		result=s1.flag<s2.flag
	}
	return
}

func main() {
	var N, L, H int
	var sample [100000]Stu
	fmt.Scanf("%d", &N)
	fmt.Scanf("%d", &L)
	fmt.Scanf("%d", &H)
	for i := 0; i < N; i++ {
		fmt.Scanf("%s%d%d", &sample[i].id, &sample[i].de, &sample[i].cai)
		sample[i].sum=sample[i].de+sample[i].cai
		if sample[i].de >= H && sample[i].cai >= H {
			sample[i].flag = 1
		} else if sample[i].de >= H && sample[i].cai >= L {
			sample[i].flag = 2
		} else if sample[i].de >= L && sample[i].cai >= L {
			sample[i].flag = 3
		} else if sample[i].de >= L && sample[i].cai < L {
			sample[i].flag = 4
		} else {
			sample[i].flag = 5
		}
	}
	sort(sample, N)
	for i := 0; i < N; i++ {
		if sample[i].flag != 5 {
			fmt.Println(sample[i].id, sample[i].de, sample[i].cai)
		}
	}
}
