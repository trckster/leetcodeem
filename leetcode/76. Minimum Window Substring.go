package main

func minWindow(s string, t string) string {
	tmap := make(map[rune]int)
	for _, c := range t {
		if _, ok := tmap[c]; ok {
			tmap[c] += 1
		} else {
			tmap[c] = 1
		}
	}

	substringLength := 0
	start := 0

	answer := ""

	for end, c := range s {
		if cnt, ok := tmap[c]; ok {
			tmap[c] -= 1
			if cnt > 0 {
				substringLength += 1
			}

			if substringLength == len(t) {
				for {
					startRune := rune(s[start])
					cnt2, ok2 := tmap[startRune]
					if cnt2 < 0 {
						tmap[startRune] += 1
					} else if ok2 {
						break
					}

					start += 1
				}

				newAnswer := s[start : end+1]
				if len(newAnswer) < len(answer) || answer == "" {
					answer = newAnswer
				}
			}
		}
	}

	return answer
}
