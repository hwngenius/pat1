package main

import (
	"fmt"
	"unicode"
)

func main(){
	var weight= [...]int{7,9,10,5,8,4,2,1,6,3,7,9,10,5,8,4,2}
	var change= [...]byte{'1','0','X','9','8','7','6','5','4','3','2'}
	var N int
	var s string
	var sum int
	fmt.Scan(&N)
	for i:=0;i<N;i++{
		fmt.Scan(&s)
		w:=0
		mask:=0
		for j:=0;j<len(weight);j++{
			if !unicode.IsDigit(rune(s[j])){
				sum++
				fmt.Println(s)
				mask=1
				break
			}
			w+=weight[j]*int(s[j]-'0')
		}
		if mask==1{
			mask=0
			continue
		}
		w%=11
		if change[w]!=s[17]{
			fmt.Println(s)
			sum++
		}
	}
	if sum==0{
		fmt.Println("All passed")
	}
}