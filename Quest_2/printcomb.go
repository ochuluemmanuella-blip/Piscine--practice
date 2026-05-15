package main

import "fmt"

func PrintComb() {
	for i := '0'; i <= '9'; i++ {
		for j := '1'; j <= '9'; j++ {
			for k := j + 1; k <= '9'; k++ {
				fmt.Printf("%c%c%c", i, j, k)

				if !(i == '7' && j == '8' && k == '9') {
					fmt.Print(", ")
				}
			}
		}
	}
	fmt.Println()
}
// func main() {
// 	PrintComb()
// }
