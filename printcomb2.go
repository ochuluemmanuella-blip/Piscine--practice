package main

import "fmt"

func Printcomb2() {
	for i := '0'; i <= '9'; i++ {
		for j := '0'; j <= '9'; j++ {
			for k := '0'; k <= '9'; k++ {
				for l := '0'; l <= '9'; l++ {
					if i > k || (i == k && j >= l) {
						continue

					}
					fmt.Printf("%c%c %c%c", i, j, k, l)
					if !(i == '9' && j == '8' && k == '9' && l == '9') {
						fmt.Print(", ")

					}
				}
			}
		}
	}
	fmt.Printf("\n")
}
func main() {
	Printcomb2()
}
