package calc

import "testing"

func TestAdd(t *testing.T) {
	var result int
	result = Add(15, 10)
	if result != 25 {
		t.Error("Expected 25, got ", result)
	}
}

func TestSub(t *testing.T) {
	var result int
	result = Sub(15, 10)
	if result != 5 {
		t.Error("Expected 5, got ", result)
	}
}
