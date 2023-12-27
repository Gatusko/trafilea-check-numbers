package numbers

import "fmt"

// Using a any Slice to store first the 100 elements
// Then use two loops
// This take N time but Take N Space
func PrintWithSlice() {
	numbers := []any{}
	for i := 0; i < 100; i++ {
		numbers = append(numbers, i+1)
	}
	for i := 2; i <= 100; i = i + 3 {
		numbers[i] = "Type 1"
	}
	for i := 4; i <= 100; i = i + 5 {
		currentVal := numbers[i]
		if currentVal == "Type 1" {
			numbers[i] = "-Type 3-"
			continue
		}
		numbers[i] = "Type 2"
	}
	for i := 0; i < len(numbers); i++ {
		fmt.Println(numbers[i])
	}
}

// For testing  Purposes
func NewWithSlice() []any {
	numbers := []any{}
	for i := 0; i < 100; i++ {
		numbers = append(numbers, i+1)
	}
	for i := 2; i <= 100; i = i + 3 {
		numbers[i] = "Type 1"
	}
	for i := 4; i <= 100; i = i + 5 {
		currentVal := numbers[i]
		if currentVal == "Type 1" {
			numbers[i] = "-Type 3-"
			continue
		}
		numbers[i] = "Type 2"
	}
	return numbers
}
