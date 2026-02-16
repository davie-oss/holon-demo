package main

import "testing"

func TestAdd(t *testing.T) {
	calc := &Calculator{}
	result := calc.Add(2, 3)
	if result != 5 {
		t.Errorf("Add(2, 3) = %d; want 5", result)
	}
}

func TestSubtract(t *testing.T) {
	calc := &Calculator{}
	result := calc.Subtract(5, 2)
	if result != 3 {
		t.Errorf("Subtract(5, 2) = %d; want 3", result)
	}
}

func TestMultiply(t *testing.T) {
	calc := &Calculator{}
	result := calc.Multiply(3, 4)
	if result != 12 {
		t.Errorf("Multiply(3, 4) = %d; want 12", result)
	}
}

func TestDivide(t *testing.T) {
	calc := &Calculator{}
	result := calc.Divide(10, 2)
	if result != 5 {
		t.Errorf("Divide(10, 2) = %d; want 5", result)
	}
}
