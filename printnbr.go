package main

import "fmt"

func PrintNbr(n int) {
	if n < 0 {
		fmt.Printf("%c", '-')
		if n/10 < 0 {
			PrintNbr(n / 10)
		}
		fmt.Printf("%c", '0'+(-(n % 10)))
	} else {
		if n/10 > 0 {
			PrintNbr(n / 10)
		}
		fmt.Printf("%c", '0'+(n%10))
	}
}

func main() {
	PrintNbr(-123)
	PrintNbr(0)
	PrintNbr(123)
	fmt.Printf("\n")
}
