package main

type Point struct {
	pacific  bool
	atlantic bool
}

var visited map[int]bool
var width, height int

func (point *Point) canFlowTo(newPoint *Point) {
	point.pacific = point.pacific || newPoint.pacific
	point.atlantic = point.atlantic || newPoint.atlantic
}

func getPoint(island [][]int, i int, j int) *Point {
	position := i*width + j

	visited[position] = true

	point := Point{false, false}

	if i+1 == height || j+1 == width {
		point.atlantic = true
	}

	if i == 0 || j == 0 {
		point.pacific = true
	}

	if i+1 != height && island[i+1][j] <= island[i][j] && !visited[(i+1)*width+j] {
		point.canFlowTo(getPoint(island, i+1, j))
	}

	if j+1 != width && island[i][j+1] <= island[i][j] && !visited[i*width+j+1] {
		point.canFlowTo(getPoint(island, i, j+1))
	}

	if i != 0 && island[i-1][j] <= island[i][j] && !visited[(i-1)*width+j] {
		point.canFlowTo(getPoint(island, i-1, j))
	}

	if j != 0 && island[i][j-1] <= island[i][j] && !visited[i*width+j-1] {
		point.canFlowTo(getPoint(island, i, j-1))
	}

	return &point
}

func pacificAtlantic(heights [][]int) [][]int {
	var answer [][]int

	width = len(heights[0])
	height = len(heights)

	for i := 0; i < height; i++ {
		for j := 0; j < width; j++ {
			visited = make(map[int]bool)

			point := getPoint(heights, i, j)

			if point.atlantic && point.pacific {
				answer = append(answer, []int{i, j})
			}
		}
	}

	return answer
}

func main() {
	arr := [][]int{
		{1, 2, 2, 3, 5},
		{3, 2, 3, 4, 4},
		{2, 4, 5, 3, 1},
		{6, 7, 1, 4, 5},
		{5, 1, 1, 2, 4},
	}

	answer := pacificAtlantic(arr)

	for _, point := range answer {
		println(point[0], point[1])
	}
}
