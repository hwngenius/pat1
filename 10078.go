package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"unicode"
)

func C(sample string) (r string) {
	mask := 0
	for i := 0; i < len(sample); i++ {
		if mask == 0 {
			mask++
		} else if mask >= 1 {
			if sample[i] == sample[i-1] {
				mask++
			} else {
				if mask > 1 {
					r += strconv.Itoa(mask)
				}
				r += string(sample[i-1])
				mask = 1
			}
		}
	}
	if mask == 1 {
		r += string(sample[len(sample)-1])
	} else {
		r += strconv.Itoa(mask)
		r += string(sample[len(sample)-1])
	}
	return
}
func D(sample string) {
	for i := 0; i < len(sample); i++ {
		if unicode.IsDigit(rune(sample[i])) {
			mask := 0
			for unicode.IsDigit(rune(sample[i+mask])) {
				mask++
			}
			for t, _ := strconv.Atoi(string(sample[i : i+mask])); t > 0; t-- {
				fmt.Printf("%c", sample[i+mask])
			}
			i += mask
		} else {
			fmt.Printf("%c", sample[i])
		}
	}
}
func main() {
	var s, sample string
	// s = "D"
	// sample = "5T2h4is i5s a3 te4st CA3a a2s"
	fmt.Scan(&s)
	input := bufio.NewScanner(os.Stdin)
	for input.Scan() {
		sample += input.Text()
	}
	if s == "C" {
		fmt.Println(C(sample))
	} else {
		D(sample)
	}
}
