package main

import (
	"fmt"
)

func moveZeroes(nums []int) {
	lastNonZeroFoundAt := 0
	for i := 0; i < len(nums); i++ {
		if nums[i] != 0 {
			nums[lastNonZeroFoundAt], nums[i] = nums[i], nums[lastNonZeroFoundAt]
			lastNonZeroFoundAt++
		}
	}
}

func main() {
	arr := []int{0, 1, 0, 3, 12}
	moveZeroes(arr)
	fmt.Println(arr) // Output: [1 3 12 0 0]
}
