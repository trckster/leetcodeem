package main

type Answer struct {
	field        [][]rune
	result       int
	currentLevel int
	size         int
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

func (answer *Answer) solve() {
	for column := 0; column < answer.size; column++ {
		if answer.field[answer.currentLevel][column] != 'E' {
			continue
		}

		answer.setQueen(column)
		answer.currentLevel++

		if answer.currentLevel == answer.size {
			answer.result++
		} else {
			answer.solve()
		}

		answer.currentLevel--
		answer.removeQueen(column)
	}
}

func totalNQueens(n int) int {
	answer := Answer{
		result:       0,
		field:        generateInitialField(n),
		currentLevel: 0,
		size:         n,
	}

	answer.solve()

	return answer.result
}

func main() {
	for i := 1; i <= 50; i++ {
		println(totalNQueens(i))
	}
}
