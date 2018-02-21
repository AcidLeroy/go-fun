package fibonacci

import (
	"testing" 
) 


func TestFibonacci(t *testing.T) {
	testNumbers := []int{0, 1, 2, 3, 4} 
	expected := []int{1, 1, 2, 3, 5}
	for i, val := range(testNumbers) {
		actual := Fibonacci(val)
		if expected[i] != actual {
			t.Errorf("Expected %d but instead got %d", expected[i], actual)
		} 
	}
}
