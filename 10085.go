package main

import (
	"fmt"
	"sort"
	"strings"
)

type rank struct {
	name  string
	score int
	num   int
}

type mySlice []rank

func (m mySlice) Len() int {
	return len(m)
}

func (m mySlice) Less(i, j int) bool {
	if m[i].score != m[j].score {
		return m[i].score > m[j].score
	} else if m[i].num != m[j].num {
		return m[i].num < m[j].num
	} else {
		return m[i].name < m[j].name
	}
}

func (m mySlice) Swap(i, j int) {
	m[i], m[j] = m[j], m[i]
}

func main() {
	var N int
	m1 := make(map[string]int)
	m2 := make(map[string]int)
	var school, id string
	var score int
	fmt.Scan(&N)
	for i := 0; i < N; i++ {
		fmt.Scanf("%s %d %s", &id, &score, &school)
		if id[0] == 'B' {
			m1[strings.ToLower(school)] += int(float64(score) / 1.5)
		} else if id[0] == 'A' {
			m1[strings.ToLower(school)] += score
		} else {
			m1[strings.ToLower(school)] += int(float64(score) * 1.5)
		}
		m2[strings.ToLower(school)]++
	}
	var result []rank
	var t rank
	var sum int
	for k, v := range m1 {
		t.name = k
		t.num = m2[k]
		t.score = v
		result = append(result, t)
		sum++
	}
	sort.Sort(mySlice(result))
	fmt.Println(sum)
	for i, r, equal := 0, 1, 0; i < len(result); i++ {
		if i > 0 {
			if result[i].score == result[i-1].score {
				equal++
			} else {
				r++
				r += equal
				equal = 0
			}
		}
		fmt.Println(r, result[i].name, result[i].score, result[i].num)
	}
}
