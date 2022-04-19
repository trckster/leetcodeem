package main

func canPlaceFlowers(flowerbed []int, n int) bool {
	availablePlaces := 0

	flowerbed = append(flowerbed, 0)
	flowerbed = append([]int{0}, flowerbed...)

	for i := 1; i+1 < len(flowerbed); i++ {
		if flowerbed[i-1] == 0 && flowerbed[i] == 0 && flowerbed[i+1] == 0 {
			flowerbed[i] = 1
			availablePlaces += 1
		}
	}

	return availablePlaces >= n
}

func main() {
	flowerbed := []int{1, 0, 0, 0, 1}
	println(canPlaceFlowers(flowerbed, 1))
	println(canPlaceFlowers(flowerbed, 2))

	flowerbed = []int{0, 0, 1, 0, 1}
	println(canPlaceFlowers(flowerbed, 1))

	flowerbed = []int{0, 1, 0, 1, 0, 0}
	println(canPlaceFlowers(flowerbed, 1))

	flowerbed = []int{0, 0, 1}
	println(canPlaceFlowers(flowerbed, 1))
	flowerbed = []int{1, 0, 0}
	println(canPlaceFlowers(flowerbed, 1))
	flowerbed = []int{0}
	println(canPlaceFlowers(flowerbed, 1))
	flowerbed = []int{0, 0, 0, 0, 1}
	println(canPlaceFlowers(flowerbed, 1))
}
