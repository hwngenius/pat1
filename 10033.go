package main

import (
	"fmt"
	"unicode"
)

func main(){
	var s1,s2 string
	fmt.Scan(&s1)
	var mask int
	var set[512]int
	for i:=0;i<len(s1);i++{
		if s1[i]=='+'{
			mask=1
			continue
		}
		t:=unicode.ToLower(rune(s1[i]))
		set[int(t)]=1
	}
	fmt.Scan(&s2)
	for i:=0;i<len(s2);i++{
		if mask==1{
			if unicode.IsUpper(rune(s2[i])){
				continue
			}
		}
		t:=unicode.ToLower(rune(s2[i]))
		if set[int(t)]!=1{
			fmt.Printf("%c",s2[i])
		}
	}
	fmt.Println()
}
