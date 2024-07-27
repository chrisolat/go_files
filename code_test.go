// package main

// import (
// 	"fmt"
// 	"testing"
// )

// func Add(a, b int) int {
// 	return a + b
// }

// func TestAdd(t *testing.T) {
// 	description := "should add 2 + 3 = 50000000"
// 	actual := Add(2, 3)
// 	expected := 50000000
// 	if actual != expected {
// 		fmt.Printf("FAIL: %s \n expected: %d \n actual: %d\n", description, expected, actual)
// 	} else {
// 		fmt.Printf("PASS: %s \n expected: %d \n actual: %d\n", description, expected, actual)
// 	}
// }

// func TestAdd2(t *testing.T) {
// 	description := "should add 2 + 3 = 5"
// 	actual := Add(2, 3)
// 	expected := 5
// 	if actual != expected {
// 		fmt.Printf("FAIL: %s \n expected: %d \n actual: %d\n", description, expected, actual)
// 	} else {
// 		fmt.Printf("PASS: %s \n expected: %d \n actual: %d\n", description, expected, actual)
// 	}
// }

// func main() {
// 	// Run tests
// 	tests := []testing.InternalTest{
// 		{"TestAdd", TestAdd},
// 		{"TestAdd2", TestAdd2},
// 	}
// 	testing.Main(func(pat, str string) (bool, error) { return true, nil }, tests, nil, nil)
// }
