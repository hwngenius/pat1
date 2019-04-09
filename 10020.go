package main

import "fmt"

type mooncake struct {
	num   int
	sum   int
	price float32
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
	var N, need int
	var bonus float32
	fmt.Scanf("%d %d", &N, &need)
	//fmt.Println(need)
	for i := 0; i < N; i++ {
		fmt.Scanf("%d", &sample[i].num)
	}
	for i := 0; i < N; i++ {
		fmt.Scanf("%d", &sample[i].sum)
		sample[i].price = float32(sample[i].sum) / float32(sample[i].num)
	}
	sort(&sample, N)
	for i := 0; i < N; i++ {
		if need > sample[i].num {
			need -= sample[i].num
			bonus += float32(sample[i].sum)
		} else if need == sample[i].num {
			bonus += float32(sample[i].sum)
			break
		} else {
			bonus += float32(sample[i].price * float32(need))
			break
		}
		//fmt.Println(sample[i].price, bonus, need)
	}
	fmt.Printf("%.2f", bonus)
}
