package main

import "fmt"

func byte2int(b byte)(int){
	if b=='P'{return 0}
	if b=='A'{return 1}
	if b=='T'{return 2}
	if b=='e'{return 3}
	if b=='s'{return 4}
	if b=='t'{return 5}
	return -1
}

func main(){
	var demo="PATest"
	var s string
	var set [6]int
	var sum int
	fmt.Scan(&s)
	for i:=0;i<len(s);i++{
		t:=byte2int(s[i])
		if t!=-1{
			set[t]++
			sum++
		}
	}
	for i:=0;i<6;i=(i+1)%6{
		if set[i]!=0{
			fmt.Printf("%c",demo[i])
			set[i]--
			sum--
			if sum==0{
				return
			}
		}
	}
}