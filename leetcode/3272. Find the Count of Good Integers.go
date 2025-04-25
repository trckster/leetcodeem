package main

import (
	"crypto/md5"
	"encoding/binary"
	"encoding/json"
	"math"
	"strconv"
)

func Reverse(s string) string {
	runes := []rune(s)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	return string(runes)
}

func fac(n int) int64 {
	res := int64(1)
	for i := 2; i <= n; i++ {
		res *= int64(i)
	}
	return res
}

func countPermutations(m map[int]int) int64 {
	all := 0
	for _, v := range m {
		all += v
	}

	answer := fac(all)

	for _, v := range m {
		answer /= fac(v)
	}

	first := all
	first -= m[0]

	return answer * int64(first) / int64(all)
}

func getPalindrome(rightPart int, partLength int, length int) int {
	rightPartAsString := strconv.Itoa(rightPart)

	missingDigitsCount := partLength - len(rightPartAsString)
	for i := 0; i < missingDigitsCount; i++ {
		rightPartAsString = "0" + rightPartAsString
	}

	var s string
	if length%2 == 0 {
		s = Reverse(rightPartAsString) + rightPartAsString
	} else {
		s = Reverse(rightPartAsString) + rightPartAsString[1:]
	}

	res, _ := strconv.Atoi(s)
	return res
}

func mapDigits(n int) map[int]int {
	m := make(map[int]int)
	for n > 0 {
		m[n%10]++
		n /= 10
	}
	return m
}

func getHash(m map[int]int) int64 {
	data, _ := json.Marshal(m)
	hash := md5.Sum(data)
	hashInt := binary.LittleEndian.Uint64(hash[:8])
	return int64(hashInt)
}

func countGoodIntegers(n int, k int) int64 {
	partLength := (n-1)/2 + 1
	maxPart := int(math.Pow(10, float64(partLength)))

	answer := int64(0)
	usedPalindromes := make(map[int64]bool)

	for i := 0; i < maxPart; i++ {
		if i%10 == 0 {
			continue
		}

		p := getPalindrome(i, partLength, n)
		if p%k != 0 {
			continue
		}
		m := mapDigits(p)

		mapHash := getHash(m)
		if usedPalindromes[mapHash] {
			continue
		}

		usedPalindromes[mapHash] = true
		perms := countPermutations(m)

		answer += perms
	}

	return answer
}
