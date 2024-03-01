package main

func maximumOddBinaryNumber(s string) string {
	zeros, ones := 0, 0

	for _, c := range s {
		if c == '1' {
			ones++
		} else {
			zeros++
		}
	}

	answer := ""
	for ; ones > 1; ones-- {
		answer += "1"
	}
	for ; zeros > 0; zeros-- {
		answer += "0"
	}

	return answer + "1"
}
