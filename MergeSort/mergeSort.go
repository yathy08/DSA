package main

import "fmt"

func MergeSort(arr []int) []int{
	if len(arr) <= 1 {
		return arr
	}
	mid := len(arr) /2 
	left := MergeSort(arr[:mid])
	right := MergeSort(arr[mid:])
	return Merge(left,right)
}

func Merge(left,right []int)[]int{
	result := make([]int,0,len(left)+len(right))
	i,j := 0,0
	for i < len(left) && j< len(right) {
		if left[i] < right[j]{
			result = append(result, left[i])
			i++
		} else {
			result = append(result, right[j])
			j++
		}
	}
	result = append(result, left[i:]...)
	result = append(result, right[j:]...)

	return result
}
func main() {
	arr := []int{38, 27, 43, 3, 9, 82, 10}
	fmt.Println("Array before sorting:", arr)
	sortedArr := MergeSort(arr)
	fmt.Println("Array after sorting:", sortedArr)
}