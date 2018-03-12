package heap 

import (
	"testing" 
	"fmt" 
) 

type compareFunc func(int, int) bool

type Heap struct {
	arr []int
	compare compareFunc // > min heap, < max heap
} 

func NewHeap(i []int, f compareFunc) (*Heap){
	m := new(Heap)
	m.compare = f
	// Insert all values of the heap
	for _, val := range(i) {
		m.Insert(val)
	}
	return m
}

func (h *Heap) Parent(i int) int {
	return (i-1)/2
} 

func (h *Heap) Left(i int) int {
	return 2*i + 1
} 

func (h *Heap) Right(i int) int {
	return 2*i+2 
} 

func (h *Heap) Size() int {
	return len(h.arr)
} 

func (h *Heap) Root() int {
	return h.arr[0]
} 


func (h *Heap) Insert(val int ) {
	h.arr = append(h.arr,val)
	i := len(h.arr)-1
	for (i != 0 &&  h.compare(h.arr[h.Parent(i)], h.arr[i])) {
		fmt.Println("i = ", i) 
		h.arr[i], h.arr[h.Parent(i)] = h.arr[h.Parent(i)], h.arr[i]
		i = h.Parent(i)
	} 
} 


func TestInsert(t *testing.T) {
	testArr := []int{1,3,5,6,10,18} 
	f := func(x, y int) bool { return x > y } 
	heap := NewHeap(testArr, f)
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
	for i, val := range(heap.arr) {

		fmt.Printf("Index = %d, val = %d, Parent: %d, Left: %d, Right: %d\n", i, val, heap.Parent(i), heap.Left(i), heap.Right(i))
	}
} 

type RunningMedian struct {
	l, r *Heap
} 

func  NewRunningMedian() *RunningMedian{
	m := new(RunningMedian)
	min := func(x, y int) bool { return x > y } 
	max := func(x, y int) bool { return x <= y } 
	// Left side is the max heap
	m.l = NewHeap([]int{}, max)
	m.r = NewHeap([]int{}, min)
	return m
}

func Odd(i int) bool {
	return (i % 2) != 0
}

func TestOdd(t *testing.T) { 
	if Odd(2) {
		t.Errorf("2 is not odd!") 
	} 
	if !Odd(1) {
		t.Errorf("1 is ODD!") 
	} 
} 

func (m *RunningMedian) GetMedian() (float64){
	lSize := m.l.Size()
	rSize := m.r.Size()
	fmt.Printf("lSize = %d, rSize = %d\n", lSize, rSize)
	if lSize == 1 && rSize == 0 {
		return float64(m.l.Root()) 
	} else if (Odd(lSize) == Odd(rSize))  {  // both even or both odd
		return float64(m.l.Root() + m.r.Root()) / 2.0
	} else { // Return the Odd one
		if !Odd(lSize) {
			return float64(m.l.Root())
		} else {
			return float64(m.r.Root())
		}
	}
	return 0.0
}

func (m *RunningMedian) AddNewValue(i int) {
	fmt.Println("Inserting i = ", i)
	if m.l.Size() > 0 {
		fmt.Println("Root of left = ", m.l.Root()) 
	} 

	if m.r.Size() > 0 {
		fmt.Println("Root of right = ", m.r.Root())
	} 
	if m.l.Size() == 0 {
		m.l.Insert(i)
	} else if m.r.Size() == 0 {
		m.r.Insert(i)
	} else if i > m.l.Root() {
		fmt.Println("Inserting right: i > m.l.root() = ", m.l.Root())
		m.r.Insert(i)
	} else {
		fmt.Println("Inserting left: i <= m.l.root() = ", m.l.Root())
		m.l.Insert(i)
	} 
} 


func TestRunningMean(t *testing.T) {
	m := NewRunningMedian()

	m.AddNewValue(3)
	expected := 3.0
	if expected != m.GetMedian() {
		t.Errorf("Expected %f as the median, but instead got %f", expected, m.GetMedian())
	}
	m.AddNewValue(4) 
	expected = (3.0 + 4.0)/2.0
	if  expected != m.GetMedian() {
		t.Errorf("Expected %f as the new median, but instead got %f", expected, m.GetMedian())
	}


	m.AddNewValue(7) 
	expected =  4.0
	if expected != m.GetMedian() {
		t.Errorf("Expected %f as the median, but instead got %f", expected, m.GetMedian())
	}

	m.AddNewValue(1.0)
	expected =  (3.0 + 4.0)/ 2.0
	if expected != m.GetMedian() {
		t.Errorf("Expected %f as the median, but instead got %f", expected, m.GetMedian())
	}

	m.AddNewValue(10)
	expected =  4.0
	if expected != m.GetMedian() {
		t.Errorf("Expected %f as the median, but instead got %f", expected, m.GetMedian())
	}
}

