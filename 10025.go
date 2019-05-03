package main

import "fmt"

var N, K int
var sample [100010]node

type node struct {
	address string
	data    int
	next    string
}

func findNode(address string) node {
	for i := 0; i < N; i++ {
		if sample[i].address == address {
			return sample[i]
		}
	}
	return node{}
}

func main() {
	var firstAddress string
	fmt.Scanf("%s%d%d", &firstAddress, &N, &K)
	for i := 0; i < N; i++ {
		fmt.Scanf("%s %d %s", &sample[i].address, &sample[i].data, &sample[i].next)
	}
	var result []node
	var tmp node
	for tmp = findNode(firstAddress); tmp.next != "-1"; tmp = findNode(tmp.next) {
		result = append(result, tmp)
	}
	result = append(result, tmp)
	var ouput []node
	t := len(result) / K
	for i := 0; i < t; i++ {
		for j := K - 1; j >= 0; j-- {
			ouput = append(ouput, result[j+i*K])
		}
	}
	for i := t * K; i < len(result); i++ {
		ouput = append(ouput, result[i])
	}
	for i := 0; i < len(ouput)-1; i++ {
		ouput[i].next = ouput[i+1].address
	}
	ouput[len(ouput)-1].next = "-1"
	for i := 0; i < len(ouput); i++ {
		fmt.Println(ouput[i].address, ouput[i].data, ouput[i].next)
	}
}
