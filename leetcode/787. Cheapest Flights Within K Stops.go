package main

import (
	pq "github.com/emirpasic/gods/queues/priorityqueue"
	"github.com/emirpasic/gods/utils"
	"sort"
)

type RouteTo struct {
	to    int
	price int
}

type Way struct {
	to              int
	step            int
	resultingAmount int
}

type BestWay struct {
	steps int
	price int
}

func wayPriority(a, b interface{}) int {
	return utils.IntComparator(a.(Way).resultingAmount, b.(Way).resultingAmount)
}

func findCheapestPrice(n int, flights [][]int, src int, dst int, k int) int {
	routes := make([][]RouteTo, n)
	for i := range routes {
		routes[i] = make([]RouteTo, 0)
	}

	for _, flight := range flights {
		routes[flight[0]] = append(routes[flight[0]], RouteTo{flight[1], flight[2]})
	}
	for i := range routes {
		sort.Slice(routes[i], func(j, jj int) bool {
			return routes[i][j].price < routes[i][jj].price
		})
	}

	answer := -1
	ways := pq.NewWith(wayPriority)
	ways.Enqueue(Way{src, 0, 0})
	bestWaysToCity := make([]BestWay, n)
	for i := 0; i < n; i++ {
		bestWaysToCity[i] = BestWay{-1, -1}
	}

	for !ways.Empty() {
		item, _ := ways.Dequeue()
		way := item.(Way)

		if bestWaysToCity[way.to].price == -1 {
			bestWaysToCity[way.to].price = way.resultingAmount
			bestWaysToCity[way.to].steps = way.step
		} else if bestWaysToCity[way.to].price <= way.resultingAmount && bestWaysToCity[way.to].steps <= way.step {
			continue
		} else {
			bestWaysToCity[way.to].price = way.resultingAmount
			bestWaysToCity[way.to].steps = way.step
		}

		for _, route := range routes[way.to] {
			newWay := Way{route.to, way.step + 1, way.resultingAmount + route.price}

			if newWay.to == dst {
				if answer == -1 || newWay.resultingAmount < answer {
					answer = newWay.resultingAmount
				}
			}

			if newWay.step <= k {
				ways.Enqueue(newWay)
			}
		}
	}

	return answer
}
