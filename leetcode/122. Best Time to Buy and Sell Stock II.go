package main

func maxProfit(prices []int) int {
	answer := 0

	for i := 0; i < len(prices)-1; i++ {
		diff := prices[i+1] - prices[i]
		if diff > 0 {
			answer += diff
		}
	}

	return answer
}

func main() {
	prices := []int{7,1,5,3,6,4}
	println(maxProfit(prices))

}
