package main

import "math"

func maximumValueSum(nums []int, k int, edges [][]int) int64 {
	sum := int64(0)
	benefitsCount := 0
	benefitsSum := int64(0)
	smallestBenefit := int64(math.MaxInt64)
	smallestDamage := int64(math.MinInt64)

	for _, n := range nums {
		sum += int64(n)

		benefit := int64(n ^ k - n)
		if benefit > 0 {
			benefitsCount++
			benefitsSum += benefit
			smallestBenefit = min(smallestBenefit, benefit)
		} else {
			smallestDamage = max(smallestDamage, benefit)
		}
	}

	if benefitsCount%2 == 0 {
		return sum + benefitsSum
	}

	return max(
		sum+benefitsSum-smallestBenefit,
		sum+benefitsSum+smallestDamage,
	)
}
