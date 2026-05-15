package main

func BasicAtoi(s string) int {
	result := 0
	for _, char := range s {
		num := int(char - '0')
		result = result*10 + num
	}
	return result
}
