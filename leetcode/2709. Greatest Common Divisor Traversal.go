package main

func canTraverseAllPairs(nums []int) bool {
	if len(nums) == 1 {
		return true
	}

	isPrime := make([]bool, 317)
	for i := range isPrime {
		isPrime[i] = true
	}
	isPrime[0] = false
	isPrime[1] = false
	for i := range isPrime {
		if !isPrime[i] {
			continue
		}
		for j := i + i; j < len(isPrime); j += i {
			isPrime[j] = false
		}
	}
	primes := make([]int, 0)
	for number, prime := range isPrime {
		if prime {
			primes = append(primes, number)
		}
	}

	factorize := func(n int) []int {
		factors := make(map[int]bool)

		for _, prime := range primes {
			for n%prime == 0 {
				n /= prime
				factors[prime] = true
			}
			if n == 1 {
				break
			}
		}

		if n > 1 {
			factors[n] = true
		}

		factorsAsArray := make([]int, len(factors))
		i := 0
		for factor := range factors {
			factorsAsArray[i] = factor
			i += 1
		}
		return factorsAsArray
	}

	graph := make(map[int]map[int]bool)
	for _, num := range nums {
		if num == 1 {
			return false
		}

		factors := factorize(num)

		for _, factor := range factors {
			if _, ok := graph[factor]; !ok {
				graph[factor] = make(map[int]bool)
			}

			for _, factor2 := range factors {
				if factor != factor2 {
					if _, ok := graph[factor]; ok {
						graph[factor][factor2] = true
					} else {
						graph[factor] = make(map[int]bool)
						graph[factor][factor2] = true
					}
				}
			}
		}
	}

	visited := make(map[int]bool)
	var dfs func(int)
	dfs = func(node int) {
		if visited[node] {
			return
		}
		visited[node] = true

		for child := range graph[node] {
			dfs(child)
		}
	}

	for key := range graph {
		dfs(key)
		break
	}

	return len(visited) == len(graph)
}
