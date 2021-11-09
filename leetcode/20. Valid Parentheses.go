package main

type Stack struct {
	data []rune
}

func newStack() *Stack {
	stack := &Stack{}

	stack.data = make([]rune, 0)

	return stack
}

func (stack *Stack) size() int {
	return len(stack.data)
}

func (stack *Stack) push(value rune) {
	stack.data = append(stack.data, value)
}

func (stack *Stack) pop() rune {
	value := stack.data[stack.size()-1]

	stack.data = stack.data[:stack.size()-1]

	return value
}

func isValid(s string) bool {
	brackets := map[rune]rune{
		'(': ')',
		'[': ']',
		'{': '}',
	}
	stack := newStack()

	for _, bracket := range s {
		if _, ok := brackets[bracket]; ok {
			stack.push(bracket)
		} else {
			if stack.size() == 0 {
				return false
			}

			last := stack.pop()
			if brackets[last] != bracket {
				return false
			}
		}
	}

	return stack.size() == 0
}

func main() {
	println(isValid("({})"))
}
