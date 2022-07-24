package main

import "testing"

func TestSum(t *testing.T) {
	testing := sum(3, 2, 1)

	result := 6

	if testing != result {
		t.Error("Expected: ", result, "Got: ", testing)
	}
}

func TestSum2(t *testing.T) {
	testing := sum(3, 2, 1)

	result := 7

	if testing != result {
		t.Error("Expected: ", result, "Got: ", testing)
	}
}