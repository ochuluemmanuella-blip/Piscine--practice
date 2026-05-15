package main

func Atoi(s string) int {
	if s == "" {
		return 0
	}

	result := 0
	sign := 1
	i := 0

	// Handle sign
	if s[0] == '+' {
		i = 1
	} else if s[0] == '-' {
		sign = -1
		i = 1
	}

	for ; i < len(s); i++ {
		char := s[i]
		if char < '0' || char > '9' {
			return 0
		}
		num := int(char - '0')
		result = result*10 + num
	}
	return result * sign
}
