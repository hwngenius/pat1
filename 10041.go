package main

import "fmt"

type Student struct {
	id   string
	test int
}

func main() {
	var N, M, c int
	var test, machie int
	var id string
	var stu [1001]Student
	fmt.Scanf("%d", &N)
	for i := 0; i < N; i++ {
		fmt.Scanf("%s %d %d", &id, &machie, &test)
		stu[machie].test = test
		stu[machie].id = id
	}
	fmt.Scanf("%d", &M)
	for i := 0; i < M; i++ {
		fmt.Scanf("%d", &c)
		fmt.Printf("%s %d\n", stu[c].id, stu[c].test)
	}
}
