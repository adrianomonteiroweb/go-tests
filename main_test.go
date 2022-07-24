package main

import "testing"

type test struct {
	data []int
	answer int
}

func TestTableSum(t *testing.T) {
	tests := []test{
		test{data: []int{1, 2, 3}, answer: 6},
		test{[]int{10, 11, 12}, 33},
		test{[]int{-5, 0, 5, 10}, 10},
	}

	for _, v := range tests {
		sumTest := sum(v.data...)

		if sumTest != v.answer {
			t.Error("Expected: ", v.answer, "Got: ", sumTest)
		}
	}
}

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