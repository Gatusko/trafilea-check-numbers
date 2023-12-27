package numbers

import "testing"

func TestNewWithSlice(t *testing.T) {
	numbers := NewWithSlice()
	for i := 0; i < 100; i++ {
		if i+1%3 == 0 && i+1%5 == 0 && numbers[i] != "Type 3" {
			t.Fatalf("Expecting Type 3 for number %v", i)
		}
		if i+1%3 == 0 && numbers[i] != "Type 1" {
			t.Fatalf("Expecting Type 1 for number %v", i)
		}
		if i+1%5 == 0 && numbers[i] != "Type 2" {
			t.Fatalf("Expecting Type 2 for number %v", i)
		}
	}
}

func TestNewWithSwitch(t *testing.T) {
	numbers := NewWithSwitch()
	for i := 0; i < 100; i++ {
		if i+1%3 == 0 && i+1%5 == 0 && numbers[i] != "Type 3" {
			t.Fatalf("Expecting Type 3 for number %v", i)
		}
		if i+1%3 == 0 && numbers[i] != "Type 1" {
			t.Fatalf("Expecting Type 1 for number %v", i)
		}
		if i+1%5 == 0 && numbers[i] != "Type 2" {
			t.Fatalf("Expecting Type 2 for number %v", i)
		}
	}
}
