package main

import "fmt"

func helper(n int, pos int, start rune, combo []rune) {
	if pos == n {
		for _, c := range combo {
			fmt.Printf("%c", c)
		}
		isLast := true
		for i, c := range combo {
			if c != '0'+rune(10-n+i) {
				isLast = false
				break
			}
		}
		if !isLast {
			fmt.Printf(", ")
		}
		return
	}
	for d := start; d <= '9'-rune(n-1-pos); d++ {
		combo[pos] = d
		helper(n, pos+1, d+1, combo)
	}
}

func PrintCombN(n int) {
	combo := make([]rune, n)
	helper(n, 0, '0', combo)
	fmt.Printf("\n")
}

func main() {
	PrintCombN(1)
	PrintCombN(3)
	PrintCombN(9)
}
