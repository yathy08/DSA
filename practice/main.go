package main

import "fmt"

func Heapify(arr []int,n int,i int){
	largest := i
	left := 2*i+1
	right:= 2*i+2

	if left < n && arr[left] > arr[largest] {
		largest = left
	}
	if right < n && arr[right] > arr[largest] {
		largest = right
	}

	if largest != i {
		arr[i], arr[largest] = arr[largest] , arr[i]
		Heapify(arr,n,largest)
	}
}

func heapSort(arr []int){
	n := len(arr)
	for i:=n/2-1;i>=0;i--{
		Heapify(arr,n,i)
	}
	for i:=n-1;i>0;i--{
		arr[0],arr[i] = arr[i],arr[0]
		Heapify(arr,i,0)
	}
}
func Reverse(arr []int){
	start := 0
	end := len(arr)-1
	for start < end {
		arr[start] ,arr[end] = arr[end] , arr[start]
		start++
		end--
	}
}
func main() {
	arr := []int{12,14,56,3,21,4,8,89,54}

	fmt.Println("Array before sorting",arr)

	heapSort(arr)
	// Reverse(arr)

	fmt.Println("Array after sorting",arr)
	largest := Largest(arr)
	fmt.Println("Largest element:", largest)

}


func Largest(arr []int)int{
	if len(arr) == 0{
		fmt.Println("Array is empty")
		return 0
	}
	largest := arr[0]
	for _, value := range arr{
      if value > largest {
		largest = value
	  }
	}
	return largest
}