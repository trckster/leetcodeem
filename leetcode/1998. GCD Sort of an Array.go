package main

import (
	"fmt"
	"math"
	"sort"
)

func factor(number int) *Set {
	primes := newSet()

	maxRoot := int(math.Sqrt(float64(number)))

	i := 2
	for i <= maxRoot {
		if number%i == 0 {
			primes.add(i)
			number /= i
		} else {
			i += 1
		}
	}

	if number > maxRoot {
		primes.add(number)
	}

	return primes
}

// Set

type Set struct {
	data map[int]int
}

var itemsCounter *Set

func newSet() *Set {
	return &Set{
		data: make(map[int]int),
	}
}

func (s *Set) add(num int) {
	if _, ok := s.data[num]; !ok {
		s.data[num] = 0
	}

	s.data[num] += 1
}

func (s *Set) toArray(setGlobalAmount bool) []int {
	var result []int

	for value, amount := range s.data {
		if setGlobalAmount {
			amount = itemsCounter.data[value]
		}

		for amount > 0 {
			result = append(result, value)
			amount -= 1
		}
	}

	return result
}

// Node

type Node struct {
	prime   int
	visited bool
	numbers *Set
	links   []int
}

func newNode(prime int) *Node {
	return &Node{
		prime:   prime,
		numbers: newSet(),
		visited: false,
	}
}

func (node *Node) add(number int) {
	node.numbers.add(number)
}

func (node *Node) connect(node2 *Node) {
	node.links = append(node.links, node2.prime)
	node2.links = append(node2.links, node.prime)
}

// Graph

type Graph struct {
	nodes map[int]*Node
}

func (g *Graph) add(number int, prime int) {
	if _, ok := g.nodes[prime]; !ok {
		g.nodes[prime] = newNode(prime)
	}

	node := g.nodes[prime]

	node.add(number)
}

func (g *Graph) connectPrimes(data []int) {
	for i := 0; i < len(data); i++ {
		prime1 := data[i]
		node1 := g.nodes[prime1]

		for j := i + 1; j < len(data); j++ {
			prime2 := data[j]
			node2 := g.nodes[prime2]

			node1.connect(node2)
		}
	}
}

func (g *Graph) dfs(node *Node, cm *ColoredMap) {
	node.visited = true
	cm.paint(node)

	for _, linkedNodeNumber := range node.links {
		linkedNode := g.nodes[linkedNodeNumber]

		if linkedNode.visited {
			continue
		}

		g.dfs(linkedNode, cm)
	}
}

// Colored group

type ColoredGroup struct {
	color                 int
	numbers               *Set
	sortedNumbers         []int
	sortedNumbersIterator int
}

func newColoredGroup(color int) *ColoredGroup {
	return &ColoredGroup{
		color:                 color,
		numbers:               newSet(),
		sortedNumbersIterator: -1,
	}
}

func (cg *ColoredGroup) add(number int) {
	cg.numbers.add(number)
}

func (cg *ColoredGroup) initSortedNumbersIteration() {
	cg.sortedNumbers = cg.numbers.toArray(true)
	sort.Ints(cg.sortedNumbers)
}

// Colored map

type ColoredMap struct {
	currentColor int
	intToColor   map[int]int
	groups       map[int]*ColoredGroup
}

func newColoredMap() *ColoredMap {
	return &ColoredMap{
		currentColor: -1,
		intToColor:   make(map[int]int),
		groups:       make(map[int]*ColoredGroup),
	}
}

func (cm *ColoredMap) nextColor() {
	cm.currentColor++
	cm.groups[cm.currentColor] = newColoredGroup(cm.currentColor)
}

func (cm *ColoredMap) paint(node *Node) {
	for number, _ := range node.numbers.data {
		cm.paintNumber(number)
	}
}

func (cm *ColoredMap) paintNumber(number int) {
	cm.intToColor[number] = cm.currentColor
	cm.groups[cm.currentColor].add(number)
}

func (cm *ColoredMap) constructColoredArray(numbers []int) []ColoredCell {
	var coloredCells []ColoredCell

	for _, number := range numbers {
		cc := ColoredCell{
			color:  cm.intToColor[number],
			number: number,
		}

		coloredCells = append(coloredCells, cc)
	}

	return coloredCells
}

func (cm *ColoredMap) getNextNumber(color int) int {
	group := cm.groups[color]

	if group.sortedNumbersIterator == -1 {
		group.initSortedNumbersIteration()
	}

	group.sortedNumbersIterator += 1

	return group.sortedNumbers[group.sortedNumbersIterator]
}

func (cm *ColoredMap) sortColoredArray(cc []ColoredCell) {
	for i, cell := range cc {
		if cell.number == 1 {
			continue
		}

		cc[i].number = cm.getNextNumber(cell.color)
	}
}

// Colored cell

type ColoredCell struct {
	color  int
	number int
}

// Main code

func constructColoredGroups(graph *Graph) *ColoredMap {
	cm := newColoredMap()

	for _, node := range graph.nodes {
		if node.visited {
			continue
		}

		cm.nextColor()

		graph.dfs(node, cm)
	}

	return cm
}

func constructPrimesGraph(nums []int) *Graph {
	graph := &Graph{nodes: make(map[int]*Node)}
	itemsCounter = newSet()

	for _, num := range nums {
		primes := factor(num)

		for prime, _ := range primes.data {
			graph.add(num, prime)
		}

		graph.connectPrimes(primes.toArray(false))
		itemsCounter.add(num)
	}

	return graph
}

func isSorted(cc []ColoredCell) bool {
	for i := 1; i < len(cc); i++ {
		if cc[i-1].number > cc[i].number {
			return false
		}
	}

	return true
}

func gcdSort(nums []int) bool {
	graph := constructPrimesGraph(nums)
	coloredMap := constructColoredGroups(graph)

	coloredCells := coloredMap.constructColoredArray(nums)
	coloredMap.sortColoredArray(coloredCells)

	return isSorted(coloredCells)
}

func main() {
	test1 := []int{7, 21, 3}
	test2 := []int{5, 2, 6, 2}
	test3 := []int{10, 5, 9, 3, 15}
	test4 := []int{2310}
	test5 := []int{1, 2, 3, 5, 7, 11, 13, 17}
	test6 := []int{4, 2, 1}
	test7 := []int{1, 1, 2, 1}
	test8 := []int{9, 7, 3}

	assert(true, gcdSort(test1))
	assert(false, gcdSort(test2))
	assert(true, gcdSort(test3))
	assert(true, gcdSort(test4))
	assert(true, gcdSort(test5))
	assert(false, gcdSort(test6))
	assert(false, gcdSort(test7))
	assert(true, gcdSort(test8))
}

// Helpers

func printColoredCells(cc []ColoredCell) {
	for _, cell := range cc {
		print(cell.number)
		print(" ")
	}
	println()
}

func (cm *ColoredMap) print() {
	println("----- Count of colors:", cm.currentColor+1, "-----")
	println("Number => color:")

	for number, color := range cm.intToColor {
		println(number, "=>", color)
	}

	println("Color => group:")

	for _, group := range cm.groups {
		group.print()
	}

	println()
}

func (cg *ColoredGroup) print() {
	print(cg.color)
	print(" => [")
	for number, _ := range cg.numbers.data {
		print(number)
		print(", ")
	}
	println("]")
}

func assert(expected bool, found bool) {
	if expected != found {
		println("Test failed")
	} else {
		println("Test passed")
	}
}

func (g *Graph) print() {
	for prime, node := range g.nodes {
		println("-----", prime, "------")
		node.print()
		println()
	}
}

func (node *Node) print() {
	print("Primes:")
	for prime, _ := range node.numbers.data {
		print(" ")
		print(prime)
	}
	println()

	print("Links: ")
	fmt.Printf("%v", node.links)
	println()
}

func (s *Set) print() {
	for value, amount := range s.data {
		println(value, " => ", amount)
	}
}
