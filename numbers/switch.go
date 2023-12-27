package numbers

import "fmt"

// As stament said  that you CAN use only 1 if
// But this can be solved only using switch stament
// This take N time and  Space is 1
func PrintWithSwitch() {
	for i := 1; i <= 100; i++ {
		switch {
		case i%3 == 0 && i%5 == 0:
			fmt.Println("Type 3")
		case i%3 == 0:
			fmt.Println("Type 1")
		case i%5 == 0:
			fmt.Println("Type 2")
		default:
			fmt.Println(i)
		}
	}
}

// For testing  Purposes
func NewWithSwitch() []any {
	numbers := []any{}
	for i := 1; i <= 100; i++ {
		switch {
		case i%3 == 0 && i%5 == 0:
			numbers = append(numbers, i)
		case i%3 == 0:
			numbers = append(numbers, i)
		case i%5 == 0:
			numbers = append(numbers, i)
		default:
			numbers = append(numbers, i)
		}
	}
	return numbers
}
