package main

import "testing"

func TestSum(t *testing.T) {

	expected := 30
	total := sum(15, 15)

	if total != expected {
		t.Errorf("Invalid expected result. Expected %d, Got: %d", expected, total)
	}
}

func TestMultiply(t *testing.T) {

	expected := 6
	total := multiply(3, 2)

	if total != expected {
		t.Errorf("Invalid expected result. Expected %d, Got: %d", expected, total)
	}
}
