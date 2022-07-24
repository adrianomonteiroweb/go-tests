package main

import "testing"

func TestSum(t *testing.T) {
	testing := sum(3, 2, 1)

	result := 6

	if testing != result {
		t.Error("Expected: ", result, "Got: ", testing)
	}
}

// Error case

// func TestSum2(t *testing.T) {
// 	testing := sum(3, 2, 1)

// 	result := 7

// 	if testing != result {
// 		t.Error("Expected: ", result, "Got: ", testing)
// 	}
// }

func TestMultiply(t *testing.T) {
	testing := multiply(10, 10)

	result := 100

	if testing != result {
		t.Error("Expected: ", result, "Got: ", testing)
	}
}

// Error case

// func TestMultiply2(t *testing.T) {
// 	testing := multiply(10, 10)

// 	result := 101

// 	if testing != result {
// 		t.Error("Expected: ", result, "Got: ", testing)
// 	}
// }