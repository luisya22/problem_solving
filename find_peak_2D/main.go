package main

import "fmt"

/*
*  [4]  [6]  [5]  [10]   [6]
*  [5]  [7]  [7]   [9]   [7]
*  [3]  [5]  [6]   [7]   [5]
 */

func main() {

	// Test Case 1
	mat := [][]int{{4, 20, 3}, {6, 7, 5}, {8, 9, 6}}
	peakX, peakY := findPeak2d(mat)

	fmt.Printf("Peak: %v, %v -> %v\n", peakX, peakY, mat[peakX][peakY])

	// Test Case 2

	mat = [][]int{{4, 5, 3}, {6, 7, 5}, {5, 7, 6}, {10, 9, 7}, {6, 7, 5}}
	peakX, peakY = findPeak2d(mat)

	fmt.Printf("Peak: %v, %v -> %v\n", peakX, peakY, mat[peakX][peakY])

	// Test Case 3
	mat = [][]int{{10, 30, 40, 50, 20}, {1, 3, 2, 500, 4}}
	peakX, peakY = findPeak2d(mat)

	fmt.Printf("Peak: %v, %v -> %v\n", peakX, peakY, mat[peakX][peakY])
}

func findPeak2d(mat [][]int) (int, int) {
	leftCol := 0
	rightCol := len(mat) - 1

	for leftCol <= rightCol {
		middleCol := leftCol + (rightCol-leftCol)/2

		// Find maximus
		max := findMaximum(mat[middleCol])

		if middleCol == len(mat)-1 {
			if mat[middleCol-1][max] < mat[middleCol][max] {
				return middleCol, max
			}
		} else if middleCol == 0 {
			if mat[middleCol+1][max] < mat[middleCol][max] {
				return middleCol, max
			}
		} else if mat[middleCol-1][max] < mat[middleCol][max] && mat[middleCol+1][max] < mat[middleCol][max] {
			return middleCol, max
		}

		if middleCol != 0 && mat[middleCol-1][max] > mat[middleCol][max] {
			rightCol = middleCol - 1
		} else {
			leftCol = middleCol + 1
		}

	}

	return 0, 0
}

func findMaximum(nums []int) int {

	max := 0
	for i, _ := range nums {
		if nums[i] > nums[max] {
			max = i
		}
	}

	return max
}

/*
*  [4]  [6]  [8]
*  [20] [7]  [9]
*  [3]  [5]  [6]
 */

/*
*  [4]  [6]  [5]  [10]   [6]
*  [5]  [7]  [7]   [9]   [7]
*  [3]  [5]  [6]   [7]   [5]
 */

// leftCol, rightCol, middleCol
