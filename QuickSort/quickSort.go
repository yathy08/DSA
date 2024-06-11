package main

import "fmt"

func QuickSort(arr []int)[]int {
	if len(arr) <= 1 {
		return arr
	}
	pivot := arr[len(arr)/2]
	left := []int{}
	right := []int{}
	middle := []int{}

	for _, num := range arr {
		if num < pivot {
			left = append(left, num)
		} else if num > pivot {
			right = append(right, num)
		} else {
			middle = append(middle, num)
		}
	}
	return append(append(QuickSort(left),middle...),QuickSort(right)...)
}

func main() {
    arr := []int{38, 27, 43, 3, 9, 82, 10}
    fmt.Println("Array before sorting:", arr)
    sortedArr := QuickSort(arr)
    fmt.Println("Array after sorting:", sortedArr)
}