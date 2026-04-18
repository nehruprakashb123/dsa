package main

import "testing"

func TestAdd(t *testing.T) {
	expected := 5
	result := Add(2, 3)

	if result != expected {
		t.Errorf("Add(2, 3) failed. Expected %d, got %d", expected, result)
	}
}