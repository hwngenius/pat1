package main

import "fmt"

func main(){
	var N int
	var set [100010]int
	var score,school int
	var Mschool,Mscore int
	fmt.Scan(&N)
	for i:=0;i<N;i++{
		fmt.Scan(&school)
		fmt.Scan(&score)
		set[school]+=score
	}
	Mschool=-1
	Mscore=-1
	for i:=0;i<=N;i++{
		if set[i]>Mscore{
			Mschool=i
			Mscore=set[i]
		}
	}
	fmt.Println(Mschool,Mscore)
}