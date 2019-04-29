package main

import "fmt"

type mooncake struct {
	num   float64
	sum   float64
	price float64
}

func sort(p *[1000]mooncake, N int) {
	for i := 1; i < N; i++ {
		for j := N - 1; j >= i; j-- {
			if p[j].price > p[j-1].price {
				p[j], p[j-1] = p[j-1], p[j]
			}
		}
	}
}

func main() {
	var sample [1000]mooncake
	var N int
	var bonus, need float64
	fmt.Scanf("%d %f", &N, &need)
	//fmt.Println(need)
	for i := 0; i < N; i++ {
		fmt.Scanf("%f", &sample[i].num)
	}
	for i := 0; i < N; i++ {
		fmt.Scanf("%f", &sample[i].sum)
		sample[i].price = float64(sample[i].sum) / float64(sample[i].num)
	}
	sort(&sample, N)
	for i := 0; i < N; i++ {
		if need > sample[i].num {
			need -= sample[i].num
			bonus += float64(sample[i].sum)
		} else if need == sample[i].num {
			bonus += float64(sample[i].sum)
			break
		} else {
			bonus += float64(sample[i].price * float64(need))
			break
		}
		//fmt.Println(sample[i].price, bonus, need)
	}
	fmt.Printf("%.2f", bonus)
}
