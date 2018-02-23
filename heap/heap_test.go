package heap 

import (
	"testing" 
	"fmt" 
) 

type MinHeap struct {
	arr []int
} 

func NewMinHeap(i []int) (*MinHeap){
	m := new(MinHeap)
	m.arr = i
	return m
}

func (h *MinHeap) Parent(i int) int {
	return i/2
} 

func (h *MinHeap) Left(i int) int {
	return 2*i + 1
} 

func (h *MinHeap) Right(i int) int {
	return 2*i+2 
} 

func (h *MinHeap) Insert(val int ) {
	h.arr = append(h.arr,val)
	i := len(h.arr)-1
	fmt.Println("hparent = ", h.arr[h.Parent(i)])
	fmt.Println("harr = ", h.arr[i])
	for (i != 0 &&  h.arr[h.Parent(i)] > h.arr[i]) {
		fmt.Println("i = ", i) 
		h.arr[i], h.arr[h.Parent(i)] = h.arr[h.Parent(i)], h.arr[i]
		i = h.Parent(i)
	} 
} 


func TestInsert(t *testing.T) {
	testArr := []int{1,3,5,6,10,18} 
	heap := NewMinHeap(testArr)
	for i, _ := range(heap.arr) {
		fmt.Println("i = ", i) 
	} 

	// Print the left, right and parent of node with value of 3
	index := 1
	left := heap.Left(index)
	right := heap.Right(index)
	parent := heap.Parent(index)

	if left != 3 {
		t.Errorf("Expected left to be 3 but instead got %d", left)
	}
	if right != 4 {
		t.Errorf("Expected right to be 4 but instead got %d", right)
	}
	if parent != 0 {
		t.Errorf("Expected parent to be 0 but instead got %d", parent)
	}

	heap.Insert(2)
	fmt.Println("length = ", len(heap.arr) ) 
	expectedArr := []int{1,2, 3,5,6,10,18}
	fmt.Println("length of expected = ", len(expectedArr) ) 
	for index, val := range(expectedArr) {
		fmt.Println("index = ", index)
		if val != heap.arr[index] {
			t.Errorf("Expected %d but got %d instead", val, heap.arr[index])
		} 
	} 
} 
