func canPlaceFlowers(flowerbed []int, n int) bool {
	if len(flowerbed) == 1 {
		return (1-flowerbed[0]) >= n
	}
	ans := 0
	for index := range flowerbed {
		if index == len(flowerbed)-1 {
			if flowerbed[index-1] == 0 && flowerbed[index] == 0 {
				flowerbed[index] = 1
				ans++
			}
		} else if index == 0 {
			if flowerbed[index] == 0 && flowerbed[index+1] == 0 {
				flowerbed[index] = 1
				ans++
			}
		} else {
			if flowerbed[index-1] == 0 && flowerbed[index] == 0 && flowerbed[index+1] == 0 {
				flowerbed[index] = 1
				ans++
			}
		}
	}
	return ans >= n
}
