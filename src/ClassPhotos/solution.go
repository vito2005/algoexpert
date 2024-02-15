package main

import (
	"fmt"
	"sort"
)

func ClassPhotos(redShirtHeights []int, blueShirtHeights []int) bool {
	sort.Ints(redShirtHeights)
	sort.Ints(blueShirtHeights)

	frontRaw := redShirtHeights
	backRaw := blueShirtHeights

	if frontRaw[0] > backRaw[0] {
		frontRaw = blueShirtHeights
		backRaw = redShirtHeights
	}

	for i := 0; i < len(backRaw); i++ {
		if backRaw[i] <= frontRaw[i] {
			return false
		}
	}
	return true
}

func main() {
	redShirtHeights := []int{5, 8, 1, 3, 4}
	blueShirtHeights := []int{6, 9, 2, 4, 5}
	expected := true
	actual := ClassPhotos(redShirtHeights, blueShirtHeights)
	fmt.Println(actual == expected)
}
