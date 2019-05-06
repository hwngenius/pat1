package main

import (
	"fmt"
	"sort"
)

type grade struct {
	online int
	mid    int
	final  int
	g      int
	pass   bool
}

type stu struct {
	G    grade
	name string
}

type mySlice []stu

func (m mySlice) Less(i, j int) bool {
	if m[i].G.g != m[j].G.g {
		return m[i].G.g > m[j].G.g
	} else {
		return m[i].name < m[j].name
	}
}

func (m mySlice) Len() int {
	return len(m)
}

func (m mySlice) Swap(i, j int) {
	m[i], m[j] = m[j], m[i]
}
func main() {
	var P, M, N int = 1, 1, 1
	var result []stu
	fmt.Scanf("%d %d %d", &P, &M, &N)
	m := make(map[string]*grade)
	var name string = "1"
	var g int
	for i := 0; i < P; i++ {
		fmt.Scanf("%s %d", &name, &g)
		m[name] = &grade{}
		m[name].online = g
	}
	for i := 0; i < M; i++ {
		fmt.Scanf("%s %d", &name, &g)
		_, e := m[name]
		if e == false {
			m[name] = &grade{}
		}
		m[name].mid = g
	}
	for i := 0; i < N; i++ {
		fmt.Scanf("%s %d", &name, &g)
		_, e := m[name]
		if e == false {
			m[name] = &grade{}
		}
		m[name].final = g
	}
	for k, v := range m {
		if v.online < 200 {
			v.pass = false
			continue
		}
		if v.final >= v.mid {
			v.g = v.final
		} else {
			v.g = int(float64(v.final)*0.6 + float64(v.mid)*0.4 + 0.5)
		}
		if v.mid == 0 {
			v.mid = -1
		}
		if v.g >= 60 {
			v.pass = true
		}
		if v.pass == true {
			var t stu
			t.name = k
			t.G = *v
			result = append(result, t)
		}
	}
	sort.Sort(mySlice(result))
	for i := 0; i < len(result); i++ {
		fmt.Println(result[i].name, result[i].G.online, result[i].G.mid, result[i].G.final, result[i].G.g)
	}
}
