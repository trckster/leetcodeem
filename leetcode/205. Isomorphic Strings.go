package main

func isIsomorphic(s string, t string) bool {
	straight := make(map[uint8]uint8)
	backward := make(map[uint8]uint8)

	for i := 0; i < len(s); i++ {
		if v, ok := straight[s[i]]; ok {
			if t[i] != v {
				return false
			}
		} else {
			straight[s[i]] = t[i]
		}

		if v, ok := backward[t[i]]; ok {
			if s[i] != v {
				return false
			}
		} else {
			backward[t[i]] = s[i]
		}
	}

	return true
}
