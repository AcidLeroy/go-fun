package fibonacci

import "fmt"

func Fibonacci( i int) int {
	if i == 0 {
		return 1 
	} else if i == 1  {
		return 1
	}  else if i < 0 {
		fmt.Println("Cannot take Fibonacci of negative number") 
		return 0
	} else {
		return Fibonacci(i-1) + Fibonacci(i-2) 
	} 
} 
