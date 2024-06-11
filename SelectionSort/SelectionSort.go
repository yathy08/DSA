package main
import "fmt"

func findKthLargest(nums []int, k int) int {
    SelectionSort(nums)
    return nums[len(nums)-k]
}
func SelectionSort(nums []int){
    n := len(nums)
    for i := 0 ; i < n-1 ; i++{
        minIndex := i
        for j := i + 1;j < n;j++{
            if nums[j] < nums[minIndex] {
                minIndex = j
             }
        }
        nums[i] , nums[minIndex] =  nums[minIndex] , nums[i]
    }
}
func main() {
	arr := []int{23,35,37,67,12,11,4,6}
	k:= 1
	fmt.Printf("%dth largest element: %d\n", k, findKthLargest(arr, k))

}