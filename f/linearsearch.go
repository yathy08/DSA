package main

import "fmt"

func linearSearch(arr []int, target int) int {
	for i, val := range arr {
		if val == target {
			return i
		}
	}
	return -1
}

func main() {
	arr := []int{1, 2, 3, 4, 5}
	target := 3
	fmt.Println(linearSearch(arr, target)) 
}
