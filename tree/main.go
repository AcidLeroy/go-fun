package main 

import ( 
	"fmt"
	"golang.org/x/tour/tree"
)

func Walk(t *tree.Tree, ch chan int ) {
	// By walking the left hand side of the binary tree first, it is guaranteed
	// that we will get the smallest values first. 
	// Walk left - Smallest values first (i.e. a sorted list) 
	// Walk right - Largest values first
	// Get value before left or right - Means to print the tree in order

	if  (t.Left != nil) {
		Walk(t.Left, ch) 
	} 
	ch <- t.Value
	if (t.Right != nil){
		Walk(t.Right, ch) 
	} 
} 

func IsSame(t1 *tree.Tree, t2 *tree.Tree) bool {
	ch1 := make(chan int)
	ch2 := make(chan int)
	walkIt := func(t *tree.Tree, ch chan int) {
		Walk(t, ch)
		close(ch)
	}
	go walkIt(t1, ch1)
	go walkIt(t2, ch2)

	for {
		val1, ok1 :=  <-ch1
		val2, ok2 :=  <-ch2
		if (!(ok1 == ok2)) {
			return false
		}
		if (val1 != val2) {
			return false
		}
		if ((ok1 == false) && (ok2 == false) ) {
			return true
		}
	}
}

func main () {
	t1 := tree.New(1)
	t2 := tree.New(2)
	val := IsSame(t1,t2)
	fmt.Println("The value is: ", val)
}

