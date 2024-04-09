package main

func timeRequiredToBuy(tickets []int, k int) int {
	time := 0

	for i := 0; ; i = (i + 1) % len(tickets) {
		if tickets[i] == 0 {
			continue
		}

		time++
		tickets[i]--

		if tickets[i] == 0 && i == k {
			return time
		}
	}
}
