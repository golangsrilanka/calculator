package test

import (
	"calculator/internal"
	"testing"
)

func TestAddTwoNumbers(t *testing.T) {
	if internal.AddTwoNumbers(5, 5) != 10 {
		t.Error("Expected 10, got ", internal.AddTwoNumbers(5, 5))
	}
}

func TestSubtractTwoNumbers(t *testing.T) {
	if internal.SubtractTwoNumbers(5, 2) != 3 {
		t.Error("Expected 0, got ", internal.SubtractTwoNumbers(5, 2))
	}
}

func TestMultiplyTwoNumbers(t *testing.T) {
	if internal.MultiplyTwoNumbers(5, 2) != 10 {
		t.Error("Expected 10, got ", internal.MultiplyTwoNumbers(5, 2))
	}
}

func TestDivideTwoNumbers(t *testing.T) {
	if internal.DivideTwoNumbers(6, 2) != 3 {
		t.Error("Expected 2, got ", internal.DivideTwoNumbers(6, 2))
	}
}
