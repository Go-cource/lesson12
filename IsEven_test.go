package main

import "testing"

func TestIsEven(t *testing.T) {
	result := IsEven(1)
	if result != "no" {
		t.Errorf("Error in test1 - expected: 'no', got: %s", result)
	}
	t.Log("Test1 finished")

	resultTrue := IsEven(2)
	if resultTrue != "yes" {
		t.Errorf("Error in test2 - expected: 'yes', got: %s", resultTrue)
	}
	t.Log("Test2 finished")
}
