package reverse

import (
	"testing" 
	"fmt"
) 

func ReverseString(s string) string{

	loopSize := len(s)/2
	// Since strings are immutable in go, we have to convert them to a "rune"
	// which is really just an alias for a int32 under the hood
	runeArr := []rune(s)
	fmt.Println("address of s =", &s, "Address of runeArr = ", &runeArr[0])
	last := len(s) - 1
	for  i := 0; i < loopSize; i++ {
		runeArr[i], runeArr[last-i] = runeArr[last-i], runeArr[i] 
	} 

	// Convert the rune slice back to a string
	result := string(runeArr)
	return result
} 

func TestReverseString(t *testing.T) {
	testString := []string{"my string", "abc", "abcd"}
	expected := []string{"gnirts ym", "cba", "dcba"}
	for i, val := range(testString){
		actual := ReverseString(val) 
		if (actual != expected[i]) {
			t.Errorf("Expected %s but got %s instead", expected[i], actual)
		} 
	} 
} 
