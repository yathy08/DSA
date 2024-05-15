package main

import (
	"fmt"
)

func intersection(arr1, arr2 []int) []int {
	m := make(map[int]bool)
	for _, num := range arr1 {
		m[num] = true
	}

	var result []int
	for _, num := range arr2 {
		if m[num] {
			result = append(result, num)
			m[num] = false
		}
	}

	return result
}

func main() {
	arr1 := []int{1, 2, 2, 1}
	arr2 := []int{2, 2}
	fmt.Println(intersection(arr1, arr2)) // Output: [2]
}
