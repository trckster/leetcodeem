package main

import (
	"bufio"
	"fmt"
	"os"
)

var reader *bufio.Reader = bufio.NewReader(os.Stdin)
var writer *bufio.Writer = bufio.NewWriter(os.Stdout)

func printAnswer(m map[int]int, bitsCount int) {
	for i := 0; i < bitsCount; i++ {
		if val, ok := m[i]; ok {
			fmt.Fprint(writer, val)
			fmt.Fprint(writer, " ")
		} else {
			fmt.Fprint(writer, "0 ")
		}
	}
	fmt.Fprintln(writer)
}

func flipFirstOne(bits []rune) int {
	i := 0

	for i+1 < len(bits) {
		if bits[i] == '1' {
			break
		}

		i += 1
	}

	if bits[i] == '1' {
		bits[i] = '0'
	} else {
		bits[i] = '1'
	}

	return i
}

func solve(movesCount int, bits []rune) {
	defer writer.Flush()
	answer := make(map[int]int)
	one := '1'

	if movesCount%2 == 1 {
		movesCount -= 1
		flippedPosition := flipFirstOne(bits)
		one = '0'
		answer[flippedPosition] = 1
	}

	i := 0
	for movesCount != 0 && i < len(bits) {
		/** Looking for the first zero */
		for i < len(bits) && bits[i] == one {
			fmt.Fprint(writer, "1")
			i += 1
		}

		if i >= len(bits) {
			break
		} else {
			if _, ok := answer[i]; ok {
				answer[i] += 1
			} else {
				answer[i] = 1
			}

			if i+1 == len(bits) {
				fmt.Fprint(writer, "0")
				if _, ok := answer[i]; ok {
					answer[i] += 1
				} else {
					answer[i] = 1
				}
				i += 1
				movesCount -= 2
				break
			} else {
				fmt.Fprint(writer, "1")
			}
		}

		i += 1

		/** Looking for the second zero */
		for i < len(bits) && bits[i] == one {
			if i+1 == len(bits) {
				break
			}

			fmt.Fprint(writer, "1")
			i += 1
		}

		if i >= len(bits) {
			break
		} else if i+1 == len(bits) {
			if bits[i] == one {
				fmt.Fprint(writer, "0")
			} else {
				fmt.Fprint(writer, "1")
			}
		} else {
			fmt.Fprint(writer, "1")
		}

		if _, ok := answer[i]; ok {
			answer[i] += 1
		} else {
			answer[i] = 1
		}

		i += 1
		movesCount -= 2
	}

	if i >= len(bits) {
		if _, ok := answer[len(bits)-1]; ok {
			answer[len(bits)-1] += movesCount
		} else {
			answer[len(bits)-1] = movesCount
		}
	} else {
		for i < len(bits) {
			if bits[i] == one {
				fmt.Fprint(writer, "1")
			} else {
				fmt.Fprint(writer, "0")
			}

			i += 1
		}
	}

	fmt.Fprintln(writer)
	printAnswer(answer, len(bits))
}

func main() {
	var t, n, k int
	var bitsString string
	var bits []rune
	fmt.Scan(&t)

	for i := 0; i < t; i++ {
		fmt.Fscan(reader, &n, &k)
		fmt.Fscan(reader, &bitsString)

		bits = []rune(bitsString)
		solve(k, bits)
	}
}
