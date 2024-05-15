package main

import "fmt"

func searchInsert(nums []int, target int) int {
	left, right := 0, len(nums)-1

	for left <= right {
		mid := left + (right-left)/2

		if nums[mid] == target {
			return mid
		} else if nums[mid] < target {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}
	return left
}

func main() {
	nums := []int{1, 3, 5, 6}
	target := 5
	fmt.Println(searchInsert(nums, target)) // Output: 2

	target = 2
	fmt.Println(searchInsert(nums, target)) // Output: 1

	target = 7
	fmt.Println(searchInsert(nums, target)) // Output: 4

	target = 0
	fmt.Println(searchInsert(nums, target)) // Output: 0
}
