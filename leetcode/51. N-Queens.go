package main

type Answer struct {
	field           [][]rune
	resultingFields [][]string
	currentLevel    int
	size            int
}

func generateInitialField(n int) [][]rune {
	field := make([][]rune, n)

	for i := 0; i < n; i++ {
		field[i] = make([]rune, n)

		for j := 0; j < n; j++ {
			field[i][j] = 'E'
		}
	}

	return field
}

func (answer *Answer) hits(i, j int) {
	if answer.field[i][j] == 'E' {
		answer.field[i][j] = '1'
	} else {
		answer.field[i][j]++
	}
}

func (answer *Answer) deprive(i, j int) {
	if answer.field[i][j] == '1' {
		answer.field[i][j] = 'E'
	} else {
		answer.field[i][j]--
	}
}

func (answer *Answer) setQueen(column int) {
	answer.field[answer.currentLevel][column] = 'Q'

	for i := 1; column-i >= 0 && answer.currentLevel+i < answer.size; i++ {
		answer.hits(answer.currentLevel+i, column-i)
	}
	for i := answer.currentLevel + 1; i < answer.size; i++ {
		answer.hits(i, column)
	}
	for i := 1; column+i < answer.size && answer.currentLevel+i < answer.size; i++ {
		answer.hits(answer.currentLevel+i, column+i)
	}
}

func (answer *Answer) removeQueen(column int) {
	answer.field[answer.currentLevel][column] = 'E'

	for i := 1; column-i >= 0 && answer.currentLevel+i < answer.size; i++ {
		answer.deprive(answer.currentLevel+i, column-i)
	}
	for i := answer.currentLevel + 1; i < answer.size; i++ {
		answer.deprive(i, column)
	}
	for i := 1; column+i < answer.size && answer.currentLevel+i < answer.size; i++ {
		answer.deprive(answer.currentLevel+i, column+i)
	}
}

func (answer *Answer) rememberField() {
	answerField := make([]string, answer.size)

	for i := 0; i < answer.size; i++ {
		line := ""
		for j := 0; j < answer.size; j++ {
			if answer.field[i][j] == 'Q' {
				line += "Q"
			} else {
				line += "."
			}
		}
		answerField[i] = line
	}

	answer.resultingFields = append(answer.resultingFields, answerField)
}

func (answer *Answer) solve() {
	for column := 0; column < answer.size; column++ {
		if answer.field[answer.currentLevel][column] != 'E' {
			continue
		}

		answer.setQueen(column)
		answer.currentLevel++

		if answer.currentLevel == answer.size {
			answer.rememberField()
		} else {
			answer.solve()
		}

		answer.currentLevel--
		answer.removeQueen(column)
	}
}

func solveNQueens(n int) [][]string {
	answer := Answer{
		resultingFields: [][]string{},
		field:           generateInitialField(n),
		currentLevel:    0,
		size:            n,
	}

	answer.solve()

	return answer.resultingFields
}

func main() {
	fields := solveNQueens(8)

	for _, field := range fields {
		for _, line := range field {
			println(line)
		}
		println()
	}
}
