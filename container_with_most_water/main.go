package main

import (
	"fmt"
)

func main() {
	height := []int{1, 8, 6, 2, 5, 4, 8, 3, 7}
	maxArea := containerWithMostWater(height)

	fmt.Printf("Most Area: %v\n", maxArea)

	height = []int{2, 3, 4, 5, 18, 17, 6}
	maxArea = containerWithMostWater(height)

	fmt.Printf("Most Area: %v", maxArea)
}

func containerWithMostWater(height []int) int {

	left := 0
	right := len(height) - 1
	maxArea := 0

	for left < right {

		width := (right - left)
		area := 0

		if height[left] > height[right] {
			area = width * height[right]
			right--
		} else {
			area = width * height[left]
			left++
		}

		if area > maxArea {
			maxArea = area
		}
	}

	return maxArea
}
