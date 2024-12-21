package calculator

import "testing"

func TestCalculate(t *testing.T) {
	result, err := Calc("2+2")
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if result != 4 {
		t.Errorf("expected 4, got %v", result)
	}
}
