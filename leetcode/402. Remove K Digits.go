package main

import (
	"errors"
)

type MonotonicStack struct {
	limit int
	stack []byte
}

func makeStack(limit int) *MonotonicStack {
	return &MonotonicStack{limit, make([]byte, 0)}
}

func (ms *MonotonicStack) pop() {
	ms.stack = ms.stack[:len(ms.stack)-1]
}

func (ms *MonotonicStack) push(digit byte) error {
	for len(ms.stack) != 0 && digit < ms.stack[len(ms.stack)-1] {
		if ms.limit <= 0 {
			return errors.New("limit of removing elements reached")
		}

		ms.pop()
		ms.limit -= 1
	}

	ms.pushOnTop(digit)
	return nil
}

func (ms *MonotonicStack) pushOnTop(digit byte) {
	ms.stack = append(ms.stack, digit)
}

func (ms *MonotonicStack) toString() string {
	number := "0" + string(ms.stack)

	for len(number) > 1 && number[0] == '0' {
		number = number[1:]
	}

	return number
}

func (ms *MonotonicStack) reachLimit() {
	for ms.limit > 0 {
		ms.pop()
		ms.limit -= 1
	}
}

func removeKdigits(num string, k int) string {
	ms := makeStack(k)

	for i := range num {
		err := ms.push(num[i])

		if err != nil {
			for ; i < len(num); i++ {
				ms.pushOnTop(num[i])
			}
			return ms.toString()
		}
	}

	ms.reachLimit()

	return ms.toString()
}
