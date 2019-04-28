package main

import "fmt"

func main() {
	var aSound, aPost, bSound, bPost int
	var N int
	var aWin, bWin int
	fmt.Scanf("%d", &N)
	for i := 0; i < N; i++ {
		fmt.Scanf("%d %d %d %d", &aSound, &aPost, &bSound, &bPost)
		if aPost == aSound+bSound && bPost != aPost {
			aWin++
		} else if bPost == aSound+bSound && aPost != bPost {
			bWin++
		}
	}
	fmt.Println(bWin, aWin)
}
