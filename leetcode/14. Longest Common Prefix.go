package main

func longestCommonPrefix(strs []string) string {
	answer := ""

	for i := 0; i < len(strs[0]); i++ {
		for _, r := range strs {
			if i >= len(r) {
				return answer
			}
			if strs[0][i] != r[i] {
				return answer
			}
		}

		answer += string(strs[0][i])
	}

	return answer
}

func main() {
	arr := []string{"flower", "flow", "flop"}
	println(longestCommonPrefix(arr))
}
