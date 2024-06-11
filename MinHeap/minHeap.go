package main

import "fmt"

type MaxHeap struct{
	arr []int
}
func (h *MaxHeap)Insert(val int){
	h.arr = append(h.arr, val)
	h.hUp(len(h.arr)-1)
}
func (h *MaxHeap)Remove()int {
  if len(h.arr) == 0 {
	fmt.Println("Heap is empty")
	return -1
  }
  min := h.arr[0]
  h.arr[0] = h.arr[len(h.arr)-1]
  h.arr = h.arr[:len(h.arr)-1]
  h.hDown(0)
  return min
}

func (h *MaxHeap)hUp(index int){
	for h.arr[index] > h.arr[parent(index)] {
		h.swap(index,parent(index))
		index = parent(index)
	}
}

func (h *MaxHeap)hDown(index int){
	lastIndex := len(h.arr)-1
	for {
		left,right := lChild(index) , rChild(index)
		largest := index
		if left < lastIndex && h.arr[left] > h.arr[largest]{
          largest = left
		}
		if right < lastIndex && h.arr[right] > h.arr[largest] {
			largest = right
		}
		if largest == index {
			break
		}
		h.swap(index,largest)
		index = largest
	}
}

func parent(i int)int{
	return (i-1)/2
}

func lChild(i int)int{
	return 2*i+1
}
func rChild(i int)int{
	return 2*i+2
}

func (h *MaxHeap) swap(i,j int){
	h.arr[i],h.arr[j] =  h.arr[j] , h.arr[i]
}

func main() {
	h := &MaxHeap{}
	nums := []int{3, 2, 1, 5, 6, 4}

	for _, num := range nums {
		h.Insert(num)
	}
	fmt.Println("Max-Heap after building:", h.arr)

	h.Insert(0)
	fmt.Println("Max-Heap after inserting 0:", h.arr)

	min := h.Remove()
	fmt.Println("Removed element:", min)
	fmt.Println("Min-Heap after removing minimum element:", h.arr)
}
