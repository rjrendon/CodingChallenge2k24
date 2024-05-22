package main

import (
	"fmt"
	"math/rand"
)

// Binary search to find the index of the smallest element >= target
func binarySearch(nums []int, left, right, target int) int {
	for left < right {
		mid := left + (right-left)/2
		if nums[mid] < target {
			left = mid + 1
		} else {
			right = mid
		}
	}
	return left
}

// Patience Sorting algorithm to find the length of the LIS
func lengthOfLIS(nums []int) int {
	// Your code here
	piles := make([]int, 0) // Stores the top card of each pile
	for _, num := range nums {
		
		pile := binarySearch(piles, 0, len(piles), num)
		if pile == len(piles) {
			piles = append(piles, num) // Create a new pile
		} else {
			piles[pile] = num 
		}
	}
	//fmt.Print("Nums found", piles)
	return len(piles)
}

func main() {
	// Example usage
	inputNumbers := []int{10, 9, 2, 5, 3, 7, 101, 18}
	result := lengthOfLIS(inputNumbers)
	fmt.Println(result) // Output: 4

// Test area if it bothers delete me
	fmt.Println("--------------------------------------------------------------")
	fmt.Println("Test Area")
	fmt.Println("--------------------------------------------------------------")

	nums2 := []int{100, 34, 1, 2, 3, 4, 5, 6}
	fmt.Println(nums2)
	fmt.Println("LIS for 1st example test is: ", lengthOfLIS(nums2))
	
	nums3 := []int{38, 22, 56, 5, 2, 7, 1, 22}
	fmt.Println(nums3)
	fmt.Println("LIS for 2nd example test is: ", lengthOfLIS(nums3))

	medArr := rand.Perm(100)
    fmt.Println(medArr)
	fmt.Println("LIS for Medium array size: ", lengthOfLIS(medArr))

	// largeArr := rand.Perm(10000000)
    // fmt.Println(largeArr)
	// fmt.Println("LIS for Very Large array size: ", lengthOfLIS(largeArr))


}
