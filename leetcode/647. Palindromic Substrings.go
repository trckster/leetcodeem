package main

func countSubstrings(s string) int {
	answer := 0

	for i := 0; i < len(s); i++ {
		indent := 0
		for i-indent >= 0 && i+indent < len(s) {
			if s[i-indent] != s[i+indent] {
				break
			}

			indent++
			answer++
		}
	}

	for i := 0; i < len(s)-1; i++ {
		indent := 0
		for i-indent >= 0 && i+1+indent < len(s) {
			if s[i-indent] != s[i+1+indent] {
				break
			}

			indent++
			answer++
		}
	}

	return answer
}

func main() {
	println(countSubstrings("abc"))
	println(countSubstrings("aaa"))
	println(countSubstrings("aaaaa"))
	println(countSubstrings("aboba"))
	println(countSubstrings("ffggff"))
}
