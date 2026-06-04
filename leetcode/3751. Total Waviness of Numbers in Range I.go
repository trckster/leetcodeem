func waviness(num int) int {
	wav := 0

	for num >= 100 {
		right := num % 10
		left := (num / 100) % 10
		mid := (num / 10) % 10

		if (mid < left && mid < right || mid > left && mid > right) {
			wav++
		}

		num /= 10
	}

	return wav
}

func totalWaviness(num1 int, num2 int) int {
	if num1 < 100 {
		num1 = 100
	}

	from := (num1 - 1) / 10 + 1
	to := num2 / 10

	w := 0
	for i := from; i < to; i++ {
		lastNumber := i % 10
		penultimateNumber := (i / 10) % 10

		if penultimateNumber > lastNumber {
			w += waviness(i) * 10 + 9 - lastNumber
		} else if penultimateNumber == lastNumber {
			w += waviness(i) * 10
		} else {
			w += waviness(i) * 10 + lastNumber
		}
	}

	for i := num1; i <= num2; i++ {
		if i >= from * 10 && i < to * 10 {
			continue
		}

		w += waviness(i)
	}

	return w
}
