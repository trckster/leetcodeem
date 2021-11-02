package main

func myAtoi(s string) int {
	answer := 0
	var sign bool

	i := 0
	for ; i < len(s) && s[i] == ' '; i++ {
	}

	if i >= len(s) {
		return answer
	}

	if s[i] == '-' {
		sign = true
		i += 1
	} else {
		sign = false
		if s[i] == '+' {
			i += 1
		}
	}

	for ; i < len(s) && s[i] >= 48 && s[i] <= 57; i += 1 {
		answer = 10*answer + int(s[i]) - 48
		if answer > 2147483648 {
			break
		}
	}

	if sign {
		answer *= -1
		if answer < -2147483648 {
			answer = -2147483648
		}
	} else if answer > 2147483647 {
		answer = 2147483647
	}

	return answer
}

func main() {
	println(myAtoi("9223372036854775808"))
}
