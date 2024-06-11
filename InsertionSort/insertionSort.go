package main

import (
	"fmt"
)

type Person struct {
	Name string
	Age  int
}

func insertionSortPersons(arr []Person) {
	n := len(arr)
	for i := 1; i < n; i++ {
		key := arr[i]
		j := i - 1
		for j >= 0 && arr[j].Age > key.Age {
			arr[j+1] = arr[j]
			j = j - 1
		}
		arr[j+1] = key
	}
}

func main() {
	arr := []Person{
		{"Alice", 30},
		{"Bob", 25},
		{"Charlie", 35},
		{"Dave", 20},
		{"Eve", 28},
	}
	fmt.Println("Original array:", arr)
	insertionSortPersons(arr)
	fmt.Println("Sorted array by age:", arr)
}
