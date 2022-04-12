package main

import "math"

type Cache struct {
	data         map[int]int
	immutable    []int
	selectable   []int
	busy         map[int]bool
	currentDepth int
}

func NewCache(immutable []int, selectable []int) *Cache {
	busy := make(map[int]bool)
	for i := 0; i < len(immutable); i++ {
		busy[i] = false
	}

	return &Cache{
		data:         make(map[int]int),
		immutable:    immutable,
		selectable:   selectable,
		busy:         busy,
		currentDepth: 0,
	}
}

func (cache *Cache) getCacheKeyByBusyness() int {
	cacheKey := 0

	for key, value := range cache.busy {
		if value {
			cacheKey |= 1 << key
		}
	}

	return cacheKey
}

func (cache *Cache) getCachedRecursion() int {
	key := cache.getCacheKeyByBusyness()

	if result, ok := cache.data[key]; ok {
		return result
	}

	result := cache.solve()
	cache.data[key] = result

	return result
}

func (cache *Cache) solve() int {
	if cache.currentDepth >= len(cache.immutable) {
		return 0
	}

	bestXor := math.MaxInt32
	var newXor int

	immutableValue := cache.immutable[cache.currentDepth]
	for i, value := range cache.selectable {
		if cache.busy[i] {
			continue
		}

		cache.busy[i] = true
		cache.currentDepth += 1

		newXor = value ^ immutableValue + cache.getCachedRecursion()
		if newXor < bestXor {
			bestXor = newXor
		}

		cache.busy[i] = false
		cache.currentDepth -= 1
	}

	return bestXor
}

func minimumXORSum(nums1 []int, nums2 []int) int {
	return NewCache(nums1, nums2).solve()
}

func main() {
	arr1 := []int{1, 0, 3}
	arr2 := []int{5, 3, 4}
	println(minimumXORSum(arr1, arr2))
}
