package main

import "fmt"

func main() {
	nums := []int{2, 3, 4, 5, 6, 7, 5, 4, 3, 2}

	peak := findPeak(nums)

	fmt.Printf("Peak: %v -> %v\n", peak, nums[peak])

	nums = []int{10, 5, 4, 5, 6, 7, 5, 4}

	peak = findPeak(nums)

	fmt.Printf("Peak: %v -> %v\n", peak, nums[peak])

	nums = []int{10, 9, 8, 7, 6, 5, 4, 3, 2, 12}

	peak = findPeak(nums)

	fmt.Printf("Peak: %v -> %v\n", peak, nums[peak])

	nums = []int{2, 3, 4, 5, 6, 7, 5, 4, 3, 2}

	peak = findPeak(nums)

	fmt.Printf("Peak: %v -> %v\n", peak, nums[peak])
}

func findPeak(nums []int) int {

	low := 0
	high := len(nums) - 1

	for low < high {
		middle := low + (high-low)/2

		if nums[middle+1] > nums[middle] {
			low = middle + 1
		} else {
			high = middle
		}
	}

	return low
}
