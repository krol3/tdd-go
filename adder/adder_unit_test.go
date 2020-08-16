package adder

import "testing"

func TestAdd(t *testing.T) {
	sum, _ := Add(2, 3)
	if sum != 5 {
		t.Errorf("Sum, got: %d, want: %d.", sum, 5)
	}
}
